package models

import (
	"fmt"
	 "sso/internal/utils"
)

// AuthorizeQuery is parsed from query parameters in /oauth/authorize endpoint.
type AuthorizeQuery struct {
	ResponseType        utils.ResponseType        `form:"response_type" binding:"required"`
	ClientID            string              `form:"client_id" binding:"required"`
	RedirectURI         string              `form:"redirect_uri" binding:"required"`
	State               string              `form:"state"`
	Scope               string              `form:"scope"`
	CodeChallenge       string              `form:"code_challenge"`
	CodeChallengeMethod utils.CodeChallengeMethod `form:"code_challenge_method"`
}

// String builds a string in form "response_type=<..>&client_id=<..>", etc.
func (i AuthorizeQuery) String() string {
	query := fmt.Sprintf("response_type=%s&client_id=%s&redirect_uri=%s",
		i.ResponseType,
		i.ClientID,
		i.RedirectURI,
	)
	if i.Scope != "" {
		query += "&scope=" + i.Scope
	}
	if i.State != "" {
		query += "&state=" + i.State
	}
	if i.CodeChallenge != "" {
		query += "&code_challenge=" + i.CodeChallenge
	}
	if i.CodeChallengeMethod != "" && i.CodeChallengeMethod != "plain" {
		query += "&code_challenge_method=" + string(i.CodeChallengeMethod)
	}
	return query
}

// CallbackQuery is built into a query string as a result in /oauth/authorize endpoint.
type CallbackQuery struct {
	Code  string
	Iss   string
	State string
}

// String builds a string in form "code=<code>&iss=<iss>" plus "&state=<state>" if present
func (ar *CallbackQuery) String() string {
	res := fmt.Sprint("code=", ar.Code, "&iss=", ar.Iss)
	if ar.State != "" {
		res += "&state=" + ar.State
	}
	return res
}


// TokenRequest is parsed from body in /oauth/token endpoint.
type TokenRequest struct {
	GrantType    utils.GrantType `form:"grant_type" binding:"required"`
	Code         string    `form:"code" binding:"required"`
	CodeVerifier string    `form:"code_verifier"`
}

// TokenResponse is sent in a json format by /oauth/token endpoint.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"identity_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
}
