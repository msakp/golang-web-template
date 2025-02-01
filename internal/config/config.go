package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresUrl string `mapstructure:"POSTGRES_URL"`
	SecretKey   string `mapstructure:"SECRET_KEY"`
}

func New() *Config {
	var config Config = Config{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Config error: %s", err.Error())
		return nil
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Printf("Config error: %s", err.Error())
		return nil
	}
	return &config
}
