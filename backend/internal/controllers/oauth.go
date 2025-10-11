package controllers

import (
	"fmt"

	"sso/internal/service"
	"sso/internal/storage/models"
	"sso/internal/utils"
	"sso/pkg/config"
	"sso/pkg/errors"

	"github.com/gin-gonic/gin"
)

type OAuthController interface {
	Authorize(c *gin.Context)
	Token(c *gin.Context)
	UserInfo(c *gin.Context)
}

type oauth struct {
	oauthService  service.OAuthService
	userService   service.AuthService
	clientService service.ClientService
	config        *config.Config
}

func NewOAuth(os service.OAuthService, us service.AuthService, cs service.ClientService, cfg *config.Config) OAuthController {
	return &oauth{os, us, cs, cfg}
}

// Authorize godoc
//
// @Summary authorize with openid
// @Description authorize
// @Param response_type query string true "Response type" Enums(code)
// @Param state query string false "Random state to be preserved by frontend"
// @Param scope query string false "Scopes separated by plus sign (openid+profile+email)"
// @Param client_id query string true "Client ID"
// @Param redirect_uri query string true "URI of client callback redirect"
// @Param code_challenge query string false "Challenge code"
// @Param code_challenge_method query string false "Challenge code verification algorithm. Defaults to plain if not present" Enums(plain, S256)
// @Param auth_request query string false "Auth request ID. Is passed ONLY after login page"
// @Tags OpenID
// @Success 302
// @Failure 400
// @Router /oauth/authorize [get]
func (o *oauth) Authorize(c *gin.Context) {
	var err error
	input := models.AuthorizeInput{}

	if err = c.ShouldBindQuery(&input); err != nil {
		// FIXME: proper oauth errors
		c.Error(errors.AppErr(400, err.Error()))
		return
	}

	if err = o.oauthService.ValidateAuthorizeInput(input); err != nil {
		c.Error(err)
		return
	}

	// Initialize or attach to auth request based on query
	var authReq *models.AuthRequest
	if input.AuthRequest == "" {
		authReq, err = o.oauthService.NewAuthReq(input)
	} else {
		authReq, err = o.oauthService.FindAuthReq(input.AuthRequest)
	}
	if err != nil {
		c.Error(err)
		return
	}

	// Check if user has valid session
	s, exists := c.Get("session")
	if exists {
		session := s.(*models.Session)
		output := models.AuthorizeOutput{
			State: input.State,
			Iss:   o.config.App.BaseURL,
			Code:  utils.RandomString(32),
		}
		authReq.Code.Scan(output.Code)
		authReq.UserID = session.UserID

		err := o.oauthService.UpdateAuthReq(authReq)
		if err != nil {
			c.Error(err)
			return
		}

		c.Redirect(302, input.RedirectURI+"?"+output.Query())
	}

	// User has no session, go to login page
	c.Redirect(302, fmt.Sprintf("/login?%s&auth_request=%s", input.Query(), authReq.ID))
}

// Token godoc
//
// @Summary get token
// @Description authorize
// @Accept application/x-www-form-urlencoded
// @Param grant_type formData string true "Grant type" Enums(authorization_code)
// @Param code formData string true "Authorization code (for authorization_code grant)"
// @Param code_verifier formData string false "PKCE code verifier"
// @Tags OpenID
// @Success 302
// @Failure 400
// @Router /oauth/token [post]
func (o *oauth) Token(c *gin.Context) {
	var input models.TokenInput
	if err := c.ShouldBind(&input); err != nil {
		c.Error(errors.AppErr(400, err.Error()))
		return
	}

	req, err := o.oauthService.FindAuthReqByCode(input.Code)
	if err != nil {
		c.Error(errors.AppErr(400, err.Error()))
		return
	}

	codeValid := utils.ValidateCodeChallenge(req.CodeChallenge, input.CodeVerifier, models.CodeChallengeMethod(req.CodeChallengeMethod))
	if !codeValid {
		c.Error(errors.AppErr(400, "invalid code verifier"))
		return
	}

	// perms := o.clientService.Permissions(req.ClientID.String(), req.UserID.String())
	// scopes := make([]string, len(perms))
	// for i, perm := range perms {
	// 	scopes[i] = perm.ScopeString()
	// }

	client, _ := o.clientService.FindClientByID(req.ClientID.String())
	user, _ := o.userService.FindUserByID(req.UserID.String())
	user.Scopes = o.userService.FindUserPermissions(user.ID.String(), client.ID.String())

	token, err := utils.NewAccessToken(client, user)
	if err != nil {
		c.Error(errors.AppErr(500, err.Error()))
		return
	}

	var output models.TokenOutput
	output.AccessToken = token
	output.TokenType = "bearer"
	output.ExpiresIn = 600

	// FIXME: if output is passed to c.JSON, it is being sent as base64 encoded
	// thats why right now i used gin.H
	c.JSON(200, output)
}

// POST /token HTTP/1.1
// Host: server.example.com
// Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
// Content-Type: application/x-www-form-urlencoded
// grant_type=authorization_code
// &code=SplxlOBeZQQYbYS6WxSbIA
// &code_verifier=3641a2d12d66101249cdf7a79c000c1f8c05d2aafcf14bf146497bed

// UserInfo godoc
func (o *oauth) UserInfo(c *gin.Context) {
	panic("unimplemented")
}
