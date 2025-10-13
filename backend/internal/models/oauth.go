package models

import (
	"fmt"
)

type (
	ResponseType        string
	CodeChallengeMethod string
)

const (
	Code  ResponseType = "code"
	Token ResponseType = "token"

	Plain CodeChallengeMethod = "plain"
	S256  CodeChallengeMethod = "S256"
)

// AuthorizeInput is parsed from query parameters in /oauth/authorize endpoint.
type AuthorizeInput struct {
	ResponseType        ResponseType        `form:"response_type" binding:"required"`
	ClientID            string              `form:"client_id" binding:"required"`
	RedirectURI         string              `form:"redirect_uri" binding:"required"`
	State               string              `form:"state"`
	Scope               string              `form:"scope"`
	CodeChallenge       string              `form:"code_challenge"`
	CodeChallengeMethod CodeChallengeMethod `form:"code_challenge_method"`
}

// Query builds a string in form "response_type=<..>&client_id=<..>", etc.
func (i AuthorizeInput) Query() string {
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

// AuthorizeOutput is built into a query string as a result in /oauth/authorize endpoint.
type AuthorizeOutput struct {
	Code  string
	Iss   string
	State string
}

// Query builds a string in form "code=<code>&iss=<iss>" plus "&state=<state>" if present
func (ar *AuthorizeOutput) Query() string {
	res := fmt.Sprint("code=", ar.Code, "&iss=", ar.Iss)
	if ar.State != "" {
		res += "&state=" + ar.State
	}
	return res
}

type GrantType string

const (
	AuthorizationCode GrantType = "authorization_code"
)

// TokenInput is parsed from body in /oauth/token endpoint.
type TokenInput struct {
	GrantType    GrantType `form:"grant_type" binding:"required"`
	Code         string    `form:"code" binding:"required"`
	CodeVerifier string    `form:"code_verifier"`
}

// TokenOutput is sent in a json format by /oauth/token endpoint.
type TokenOutput struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
}
