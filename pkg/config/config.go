package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Enviroment string `mapstructure:"ENVIRONMENT"`
	DB_SOURCE  string `mapstructure:"DB_SOURCE"`
	DBDriver   string `mapstructure:"DB_DRIVER"`
	PORT       string `mapstructure:"PORT"`
	CookiesKey string `mapstructure:"COOKIES_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath("../../")
	viper.SetConfigName(".")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
