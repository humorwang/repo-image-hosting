package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	Conf *config
)

type config struct {
	App app
}

func Setup() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	Conf = &config{}
	var err error
	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("load app config error: %v", err)
	}
	EnvSettingApp()
}
