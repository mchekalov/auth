package config

import "github.com/joho/godotenv"

// Load text from .env file
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
