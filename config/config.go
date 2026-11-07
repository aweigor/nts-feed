package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Nts    NtsAPIProperties
	Server ServerProperties
	Auth   AuthConfig
	Feeds  FeedsConfig
	Shows  ShowsConfig
}

type AuthConfig struct {
	Secret string `mapstructure:"secret"`
}

type FeedsConfig struct {
	Nts       NtsAPIProperties
	Tracklist bool `mapstructure:"tracklist"`
}

type ShowsConfig struct {
	Nts NtsAPIProperties
}

type NtsAPIProperties struct {
	APIV2Url string `mapstructure:"api_v2_url"`
}

type ServerProperties struct {
	PublicURL string `mapstructure:"public_url"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, falling back to environment variables")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// read .env
	viper.Set("auth.secret", os.Getenv("API_SECRET"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
