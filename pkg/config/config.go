package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Enviroment       string `mapstructure:"ENVIRONMENT"`
	DB_SOURCE        string `mapstructure:"DB_SOURCE"`
	DBDriver         string `mapstructure:"DB_DRIVER"`
	PORT             string `mapstructure:"PORT"`
	CookiesKey       string `mapstructure:"COOKIES_KEY"`
	MinioAccessKey   string `mapstructure:"MINIO_ACCESS_KEY"`
	MinioSecretKey   string `mapstructure:"MINIO_SECRET_KEY"`
	MinioBucketName  string `mapstructure:"MINIO_BUCKET_NAME"`
	SmptPort         string `mapstructure:"SMTP_PORT"`
	SmtpHost         string `mapstructure:"SMTP_HOST"`
	SmtpFrom         string `mapstructure:"SMTP_FROM_EMAIL"`
	SmtpFromPassword string `mapstructure:"SMTP_FROM_EMAIL_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
