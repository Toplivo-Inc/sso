package handlers

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// FIXME: no unicode when serving html like this

type Frontend struct {
	mainPage  []byte
	loginPage []byte
	registerPage []byte
}

func MustLoadFrontend() *Frontend {
	var front Frontend
	if err := loadFile("static/frontend/main.html", &front.mainPage); err != nil {
		panic(err)
	}
	if err := loadFile("static/frontend/login.html", &front.loginPage); err != nil {
		panic(err)
	}
	if err := loadFile("static/frontend/register.html", &front.registerPage); err != nil {
		panic(err)
	}
	return &front
}

func loadFile(path string, v *[]byte) error {
	f, err := os.OpenFile(path, os.O_RDONLY, 0x444)
	if err != nil {
		return err
	}
	*v, err = io.ReadAll(f)
	return err
}

func (f Frontend) Main(c *gin.Context) {
	c.Header("content-type", "text/html")
	c.Writer.Write(f.mainPage)
}

func (f Frontend) Login(c *gin.Context) {
	c.Header("content-type", "text/html")
	c.Writer.Write(f.loginPage)
}

func (f Frontend) Register(c *gin.Context) {
	c.Header("content-type", "text/html")
	c.Writer.Write(f.registerPage)
}

