package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		return nil
	}
	key := os.Getenv("KEY")

	return &Config{
		Key: key,
	}

}
