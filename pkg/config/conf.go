package config

import (
	"fmt"
	// "net/url"
	// "os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/github"
)

type Config struct {
	App App
	DB  DB
	// Github oauth2.Config
}

type App struct {
	Production bool   `env:"PROD" envDefault:"false"`
	BaseURL    string `env:"BASE_URL,required"`
}

type DB struct {
	User     string `env:"POSTGRES_USER,notEmpty"`
	Password string `env:"POSTGRES_PASS,notEmpty"`
	Host     string `env:"POSTGRES_HOST,notEmpty"`
	Port     string `env:"POSTGRES_PORT,notEmpty"`
}

func (db *DB) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=sso port=%s sslmode=disable", db.Host, db.User, db.Password, db.Port)
}

func MustLoad() *Config {
	var cfg Config
	godotenv.Load()
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	// config.Github.ClientID = os.Getenv("GITHUB_CLIENT_ID")
	// config.Github.ClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	// config.Github.RedirectURL = fmt.Sprintf("%s/api/v1/oauth/github/callback", Misc.APIBase.String())
	// config.Github.Scopes = []string{"read:user"}
	// config.Github.Endpoint = github.Endpoint

	return &cfg
}
