package service

import (
	"net/url"

	"sso/internal/config"
	"sso/internal/errors"
	"sso/internal/models"
	"sso/internal/repository"
	"sso/internal/utils"

	"github.com/google/uuid"
)

type OAuthService interface {
	ValidateAuthorizeInput(models.AuthorizeInput) error
	NewAuthReq(code string, input models.AuthorizeInput, userID uuid.UUID) (*models.AuthCodes, error)
	AuthCodeByState(state string) (*models.AuthCodes, error)
	AuthCodeByCode(id string) (*models.AuthCodes, error)
	UpdateAuthReq(*models.AuthCodes, *models.Session, models.AuthorizeOutput) error

	CallbackData(models.AuthorizeInput) models.AuthorizeOutput
}

type oauthService struct {
	clientRepo repository.ClientRepository
	authRepo   repository.AuthRepository
	config     config.Config
}

func NewOAuthService(clientRepo repository.ClientRepository, authRepo repository.AuthRepository, cfg config.Config) OAuthService {
	return &oauthService{
		clientRepo,
		authRepo,
		cfg,
	}
}

// ValidateAuthorizeReq validates auth request and returns a client
func (o oauthService) ValidateAuthorizeInput(req models.AuthorizeInput) error {
	switch req.ResponseType {
	case models.Code, models.Token:
	default:
		return errors.AppErr(400, "invalid response_type")
	}
	if req.CodeChallenge == "" && req.CodeChallengeMethod != "" {
		return errors.AppErr(400, "redundnant code challenge method without code challenge")
	}
	switch req.CodeChallengeMethod {
	case "", models.Plain, models.S256:
	default:
		return errors.AppErr(400, "invalid code_challenge_method")
	}

	if _, err := url.ParseRequestURI(req.RedirectURI); err != nil {
		return errors.AppErr(400, "invalid redirect_uri")
	}

	client, err := o.clientRepo.ClientByID(req.ClientID)
	if err != nil {
		return errors.AppErr(400, "invalid client_id")
	}

	if req.RedirectURI != client.CallbackURL {
		return errors.AppErr(400, "redirect_uri doesn't match client's redirect uri")
	}

	return nil
}

func (o oauthService) NewAuthReq(code string, input models.AuthorizeInput, userID uuid.UUID) (*models.AuthCodes, error) {
	clientID,_ := uuid.Parse(input.ClientID)
	newCode := models.AuthCodes {
		Code: code,
		ResponseType: input.RedirectURI,
		ClientID: clientID,
		RedirectURI: input.RedirectURI,
		Scopes: input.Scope,
		CodeChallenge: input.CodeChallenge,
		CodeChallengeMethod: string(input.CodeChallengeMethod),
		State: input.State,
		UserID: userID,
	}
	err := o.authRepo.Create(&newCode)
	if err != nil {
		return nil, err
	}
	return o.authRepo.AuthReqByCode(newCode.Code)
}

func (o oauthService) AuthCodeByState(state string) (*models.AuthCodes, error) {
	return o.authRepo.AuthReqByState(state)
}

func (o oauthService) AuthCodeByCode(code string) (*models.AuthCodes, error) {
	return o.authRepo.AuthReqByCode(code)
}

func (o oauthService) CallbackData(input models.AuthorizeInput) models.AuthorizeOutput {
	output := models.AuthorizeOutput{
		State: input.State,
		Iss:   o.config.App.BaseURL,
		Code:  utils.RandomString(32),
	}
	return output
}

func (o oauthService) UpdateAuthReq(req *models.AuthCodes, session *models.Session, output models.AuthorizeOutput) error {
	req.Code = output.Code
	req.UserID = session.UserID
	return o.authRepo.Update(req)
}
