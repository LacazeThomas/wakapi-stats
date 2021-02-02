package config

import (
	"github.com/caarlos0/env"
)

const (
	WakapiSummary = "/api/summary"
)

type GitConfig struct {
	AccessToken string `env:"AccessToken"`
	Branch      string `env:"Branch"`
	Message     string `env:"Message"`
	CommitName  string `env:"CommitName"`
	CommitEmail string `env:"CommitEmail"`
	UserName    string `env:"UserName"`
	RepoName    string `env:"RepoName"`
}

type AppConfig struct {
	APIKey string `env:"API_KEY"`
	APIURL string `env:"API_URL"`
	Env    string `default:"dev" env:"ENVIRONMENT"`
}

func (c *AppConfig) IsDev() bool {
	return IsDev(c.Env)
}

func IsDev(env string) bool {
	return env == "dev" || env == "development"
}

func Load() (cfg AppConfig, git GitConfig) {

	env.Parse(&cfg)
	env.Parse(&git)
	return
}
