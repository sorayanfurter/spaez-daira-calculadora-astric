package environment

import (
	"github.com/joho/godotenv"
)

// LoadEnv lee archivo .env
func LoadEnv() error {

	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
