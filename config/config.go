package config

import (
	"github.com/caarlos0/env"
)

const (
	WakapiSummary = "/api/summary"
)

type GitConfig struct {
	AccessToken string `env:"PLUGIN_ACCESSTOKEN"`
	Branch      string `env:"PLUGIN_BRANCH"`
	Message     string `env:"PLUGIN_MESSAGE"`
	CommitName  string `env:"PLUGIN_COMMITNAME"`
	CommitEmail string `env:"PLUGIN_COMMITEMAIL"`
	UserName    string `env:"PLUGIN_USERNAME"`
	RepoName    string `env:"PLUGIN_REPO"`
}

type AppConfig struct {
	APIKey string `env:"PLUGIN_API_KEY"`
	APIURL string `env:"PLUGIN_API_URL"`
	Env    string `default:"dev" env:"PLUGIN_ENVIRONMENT"`
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
