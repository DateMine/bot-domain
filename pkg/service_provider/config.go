package service_provider

import (
	"github.com/joho/godotenv"
)

func LoadConfig(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
