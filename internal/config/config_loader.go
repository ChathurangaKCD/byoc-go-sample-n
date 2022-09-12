package config

import (
	"log"
	"os"
	"strconv"
)

const (
	DefaultPort = 8000
)

var (
	EnvName        = "ENV"
	MaxPetsAllowed = "MAX_PETS_ALLOWED"
)

var config Config

func GetConfig() *Config {
	return &config
}

func LoadConfig() (*Config, error) {
	config = Config{
		Port:           DefaultPort,
		Env:            os.Getenv(EnvName),
		MaxPetsAllowed: getEnvInt(MaxPetsAllowed, 5),
	}

	return &config, nil
}

func getEnvInt(key string, defaultVal int) int {
	s := os.Getenv(key)
	if s == "" {
		return defaultVal
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return v
}
