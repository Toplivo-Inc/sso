package service

import (
	"net/url"

	"sso/internal/storage"
	"sso/internal/storage/models"
	"sso/pkg/errors"
)

type oauthService struct {
	clientRepo storage.ClientRepository
	authRepo   storage.AuthRepository
}

func NewOAuthService(clientRepo storage.ClientRepository, authRepo storage.AuthRepository) OAuthService {
	return &oauthService{
		clientRepo,
		authRepo,
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
		return  errors.AppErr(400, "redundnant code challenge method without code challenge")
	}
	switch req.CodeChallengeMethod {
	case "", models.Plain, models.S256:
	default:
		return  errors.AppErr(400, "invalid code_challenge_method")
	}

	if _, err := url.ParseRequestURI(req.RedirectURI); err != nil {
		return  errors.AppErr(400, "invalid redirect_uri")
	}

	_, err := o.clientRepo.GetByID(req.ClientID)
	if err != nil {
		return errors.AppErr(400, "invalid client_id")
	}

	return nil
}

func (o oauthService) NewAuthReq(input models.AuthorizeInput) (*models.AuthRequest, error) {
	req, err := o.authRepo.CreateFromInput(&input)
	if err != nil {
		return nil, err
	}
	return o.authRepo.GetByID(req.ID.String())
}

func (o oauthService) FindAuthReq(id string) (*models.AuthRequest, error) {
	return o.authRepo.GetByID(id)
}

func (o oauthService) FindAuthReqByCode(code string) (*models.AuthRequest, error) {
	return o.authRepo.GetByCode(code)
}

func (o oauthService) UpdateAuthReq(req *models.AuthRequest) error {
	return o.authRepo.Update(req)
}
