package controllers

import (
	"sso/internal/config"
	"sso/internal/errors"
	"sso/internal/models"
	"sso/internal/service"
	"sso/internal/utils"

	"github.com/gin-gonic/gin"
)

type OAuthController interface {
	Authorize(c *gin.Context)
	Token(c *gin.Context)
	UserInfo(c *gin.Context)
	Logout(c *gin.Context)
}

type oauth struct {
	oauthService  service.OAuthService
	authService   service.AuthService
	clientService service.ClientService
	config        config.Config
}

func NewOAuth(os service.OAuthService, us service.AuthService, cs service.ClientService, cfg config.Config) OAuthController {
	return &oauth{os, us, cs, cfg}
}

// Authorize godoc
//
// @Summary authorize with openid
// @Description authorize
// @Param response_type query string true "Response type" Enums(code)
// @Param state query string false "Random state to be preserved by frontend (against XSRF attacks)"
// @Param client_id query string true "Client ID"
// @Param redirect_uri query string true "URI of client callback redirect"
// @Param scope query string false "Scopes separated by plus sign (openid+profile+email)"
// @Param code_challenge query string false "Challenge code"
// @Param code_challenge_method query string false "Challenge code verification algorithm. Defaults to plain if not present" Enums(plain, S256)
// @Tags OpenID
// @Success 302
// @Failure 400
// @Router /oauth/authorize [get]
func (o *oauth) Authorize(c *gin.Context) {
	var err error
	input := models.AuthorizeInput{}

	if err = c.ShouldBindQuery(&input); err != nil {
		// FIXME: different oauth error codes
		// https://openid.net/specs/openid-connect-core-1_0.html#AuthError
		c.Redirect(302, errors.OauthErrRedirect(input.RedirectURI, "invalid_request", err.Error(), input.State))
		return
	}

	if err = o.oauthService.ValidateAuthorizeInput(input); err != nil {
		c.Redirect(302, errors.OauthErrRedirect(input.RedirectURI, "invalid_request", err.Error(), input.State))
		return
	}

	var redirectURL string
	// Check if user has valid session
	s, exists := c.Get("session")
	if exists {
		session := s.(*models.Session)
		output := o.oauthService.CallbackData(input)
		_, err := o.oauthService.NewAuthReq(output.Code, input, session.UserID)
		if err != nil {
			c.Redirect(302, errors.OauthErrRedirect(input.RedirectURI, "internal_server_error", err.Error(), input.State))
			return
		}

		// User has session, go to callback
		redirectURL = input.RedirectURI + "?" + output.Query()
	} else {
		// User has no session, go to login page
		redirectURL = "/login?" + input.Query()
	}

	c.Redirect(302, redirectURL)
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

	req, err := o.oauthService.AuthCodeByCode(input.Code)
	if err != nil {
		c.Error(errors.AppErr(400, err.Error()))
		return
	}

	codeValid := utils.ValidateCodeChallenge(req.CodeChallenge, input.CodeVerifier, models.CodeChallengeMethod(req.CodeChallengeMethod))
	if !codeValid {
		c.Error(errors.AppErr(400, "invalid code verifier"))
		return
	}

	client, _ := o.clientService.FindClientByID(req.ClientID.String())
	user, _ := o.authService.FindUserByID(req.UserID.String())
	user.Scopes = o.authService.FindUserPermissions(user.ID.String(), client.ID.String())

	access, err := utils.NewAccessToken(client, user, o.config)
	if err != nil {
		c.Error(errors.AppErr(500, err.Error()))
		return
	}
	identity, err := utils.NewIDToken(client, user, o.config)
	if err != nil {
		c.Error(errors.AppErr(500, err.Error()))
		return
	}

	var output models.TokenOutput
	output.AccessToken = access
	output.TokenType = "bearer"
	output.ExpiresIn = 300
	output.IDToken = identity

	c.Header("cache-control", "no-store")
	c.JSON(200, output)
}

// FIXME:
// userinfo should take bearer auth header, but swaggo/swag dun' hav nuthin' as ya se
// and pull request from february that adds the definition is open and has 0 comments

// UserInfo godoc
//
// @Summary get user info
// @Description returns claims
// @Tags OpenID
// @Success 200
// @Failure 401
// @Failure 500
// @Router /oauth/logout [get]
func (o *oauth) UserInfo(c *gin.Context) {
	panic("unimplemented")
	// TODO: i should use sso's secret for signing id tokens
	// i need bootstrapping
	// idToken := strings.Split(c.GetHeader("Authorization"), " ")[1]
	// err := utils.VerifyToken(nil, idToken)
	// if err != nil {
	// }
}

// Logout godoc
//
// @Summary logout from session
// @Description deletes session from DB and revokes session token cookie
// @Param redirect_uri query string false "URI to redirect after"
// @Tags OpenID
// @Success 200
// @Success 204
// @Success 302
// @Failure 500
// @Router /oauth/logout [get]
func (o *oauth) Logout(c *gin.Context) {
	redir := c.Query("redirect_uri")
	s, exists := c.Get("session")
	if !exists {
		c.Status(200)
		return
	}

	session := s.(*models.Session)
	err := o.authService.DeleteSession(session)
	if err != nil {
		c.Error(errors.AppErr(500, err.Error()))
		return
	}

	c.SetCookie("TOPLIVO_SESSION_TOKEN", "", -1, "/", "localhost", false, true)
	c.SetCookie("TOPLIVO_ACCESS_TOKEN", "", -1, "/", "localhost", false, true)
	c.SetCookie("TOPLIVO_IDENTITY_TOKEN", "", -1, "/", "localhost", false, true)
	if redir != "" {
		c.Redirect(302, redir)
	} else {
		c.Status(204)
	}
}
