package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/rand/v2"
	"sso/internal/models"
)

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Uint32N(uint32(len(letterRunes)))]
	}
	return string(b)
}

func ValidateCodeChallenge(challenge, verifier string, method models.CodeChallengeMethod) bool {
	// plain
	//   code_challenge = code_verifier
	switch method {
	case models.Plain:
		return challenge == verifier
	case models.S256:
		return challenge == GenerateS256Challenge(verifier)
	}
	return false
}

func GenerateS256Challenge(verifier string) string {
	// S256
	//   code_challenge = BASE64URL-ENCODE(SHA256(ASCII(code_verifier)))
	h := sha256.New()
	h.Write([]byte(verifier))
	// NOTE: im not sure that bs should be 64 long
	// also no ascii encoding [strconv.QuoteToASCII(verifier)]
	bs := make([]byte, 64)
	hex.Encode(bs, h.Sum(nil))
	return base64.URLEncoding.EncodeToString(bs)
}
