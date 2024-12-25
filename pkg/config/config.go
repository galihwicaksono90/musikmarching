package config

import (
	"os"
	"path/filepath"

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

func LoadConfig() (config Config, err error) {
	ex, err := os.Executable()
	if err != nil {
		return 
	}
	execPath := filepath.Dir(ex)

	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Add multiple config paths in order of priority
	viper.AddConfigPath(".")              // First look in current directory
	viper.AddConfigPath(execPath)         // Then in the binary's directory
	viper.AddConfigPath("/home/galih/Documents/musikmarching-be/")  // Then in /etc/yourapp

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
