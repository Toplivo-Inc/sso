package utils

type (
	ResponseType        string
	CodeChallengeMethod string
	GrantType           string
)

const (
	Code  ResponseType = "code"
	Token ResponseType = "token"

	Plain CodeChallengeMethod = "plain"
	S256  CodeChallengeMethod = "S256"

	AuthorizationCode GrantType = "authorization_code"
)
