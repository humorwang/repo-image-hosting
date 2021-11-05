package config

import "os"

type app struct {
	Port     string
	Platform string
	Token    string
	Owner    string
	Repo     string
	Path     string
	Branch   string
	Cnd      string
}

func EnvSettingApp() {
	envPort := os.Getenv("APP_PORT")
	envPlatform := os.Getenv("APP_PLATFORM")
	envToken := os.Getenv("APP_TOKEN")
	envOwner := os.Getenv("APP_OWNER")
	envRepo := os.Getenv("APP_REPO")
	envPath := os.Getenv("APP_PATH")
	envBranch := os.Getenv("APP_BRANCH")
	envCdn := os.Getenv("APP_CDN")
	if envPort != "" {
		Conf.App.Port = envPort
	}
	if envPlatform != "" {
		Conf.App.Platform = envPlatform
	}
	if envToken != "" {
		Conf.App.Token = envToken
	}
	if envOwner != "" {
		Conf.App.Owner = envOwner
	}
	if envRepo != "" {
		Conf.App.Repo = envRepo
	}
	if envPath != "" {
		Conf.App.Path = envPath
	}
	if envBranch != "" {
		Conf.App.Branch = envBranch
	}
	if envCdn != "" {
		Conf.App.Cnd = envCdn
	}
}
