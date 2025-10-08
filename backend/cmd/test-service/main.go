package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"sso/internal/utils"
	"sso/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Codes struct {
	challenge string
	verifier  string
}

var (
	codes  map[string]Codes
	client http.Client = http.Client{}
)

func main() {
	godotenv.Load()
	config.MustLoad()

	codes = make(map[string]Codes)

	r := gin.New()
	r.LoadHTMLGlob("cmd/test-service/*")
	r.GET("/login", login)
	r.GET("/callback", callback)
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "cmd/test-service-provider/index.tmpl", gin.H{
			"title": "Frontend",
		})
	})

	r.Run(":9102")
}

func login(c *gin.Context) {
	state := utils.RandomString(8)
	verifier := utils.RandomString(8)
	challenge := utils.GenerateS256Challenge(verifier)
	codes[state] = Codes{challenge, verifier}
	url := fmt.Sprintf(
		"http://localhost:9100/oauth/authorize?response_type=code&state=%s&client_id=%s&redirect_uri=http://localhost:9102/callback&code_challenge_method=S256&code_challenge=%s",
		state,
		"6ffa873a-01d5-4879-9880-60d4025af9f3",
		challenge)
	c.Redirect(302, url)
}

func callback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")
	body := fmt.Sprintf("grant_type=authorization_code&code=%s&code_verifier=%s", code, codes[state].verifier)

	url := "http://localhost:9100/oauth/token"
	resp, err := client.Post(url, "application/x-www-form-urlencoded", bytes.NewReader([]byte(body)))
	if err != nil {
		c.JSON(500, resp)
	}
	b, _ := io.ReadAll(resp.Body)
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}
	json.Unmarshal(b, &tokenResponse)

	c.SetCookie("TOPLIVO_ACCESS_TOKEN", tokenResponse.AccessToken, tokenResponse.ExpiresIn, "/", "localhost", false, false)
	c.Redirect(302, "/")
}
