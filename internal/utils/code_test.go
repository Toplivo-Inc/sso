package utils

import (
	"sso/internal/storage/models"
	"testing"
)


func TestValidateCodeChallenge(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		challenge string
		verifier  string
		method    models.CodeChallengeMethod
		want      bool
	}{
		{
			"plain",
			"bebra",
			"bebra",
			models.Plain,
			true,
		},
		{
			"basic",
			"NDhiOTcyNzBlMTAxYzE2ZTJkOGNiNWJiMzA3YjhlMzllZjNlZTQwM2I2NGFiOTg3NDA1ZGI4YjExZDBkNTE2ZQ==",
			"bebra",
			models.S256,
			true,
		},
		{
			"basic",
			"ODNjMmFmZWM0Yzg4Y2ZjMmU1ZWQyNjg2NmY2NTFkNjVhNjA3Y2MyZTkzNTZkOWJmNDQ1N2M1NjlkNjFkNDBiNQ==",
			"bebra2",
			models.S256,
			true,
		},
		{
			"invalid method",
			"bebra",
			"bebra",
			"invalid",
			false,
		},
		{
			"non ascii",
			"NDk4OTYzNjYyYzA2NmM1MWRhNjU0ZGNjODQyYmE4ZTZkNmYzZDc1Y2FlZTJmYTE5OTBhOWE4NmZiZDk4MTIxMg==",
			"бебра",
			"invalid",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateCodeChallenge(tt.challenge, tt.verifier, tt.method)
			if got != tt.want {
				t.Errorf("ValidateCodeChallenge() = %v, want %v", got, tt.want)
			}
		})
	}
}
