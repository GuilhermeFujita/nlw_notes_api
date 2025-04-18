package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string
	DBDriver   string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort int
	Timeout    time.Duration
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(".env")
	v.AutomaticEnv()

	v.SetDefault("DB_PORT", 5432)
	v.SetDefault("SERVER_PORT", 9090)
	v.SetDefault("TIMEOUT", "5s")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	timeout, err := time.ParseDuration(v.GetString("TIMEOUT"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Env:        v.GetString("ENV"),
		DBDriver:   v.GetString("DB_DRIVER"),
		DBHost:     v.GetString("DB_HOST"),
		DBPort:     v.GetInt("DB_PORT"),
		DBUser:     v.GetString("DB_USER"),
		DBPassword: v.GetString("DB_PASSWORD"),
		DBName:     v.GetString("DB_NAME"),
		ServerPort: v.GetInt("SERVER_PORT"),
		Timeout:    timeout,
	}

	return cfg, nil
}
