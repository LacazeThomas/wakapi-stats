package config

import (
	"github.com/caarlos0/env"
)

const (
	WakapiSummary = "/api/summary"
)

var appConfig AppConfig

//AppConfig struct contains WakaAPI needed env
type AppConfig struct {
	APIKey string `env:"PLUGIN_API_KEY"`
	APIURL string `env:"PLUGIN_API_URL"`
	Env    string `default:"dev" env:"PLUGIN_ENVIRONMENT"`
}

//IsDev return true if application is on dev stack
func (c *AppConfig) IsDev() bool {
	return IsDev(c.Env)
}

//IsDev return true if application is on dev stack
func IsDev(env string) bool {
	return env == "dev" || env == "development"
}

//InitAppConfig struct with env variables
func initAppConfig() {
	env.Parse(&appConfig)
	return
}

//GetAppConfig return initialize config structure with variable env
func GetAppConfig() AppConfig {
	if (appConfig == AppConfig{}) {
		initAppConfig()
	}

	return appConfig
}
