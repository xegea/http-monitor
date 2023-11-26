package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env    string
	DbConn string
	ApiKey string
	ApiUrl string
}

func LoadConfig(env *string) (*Config, error) {

	err := godotenv.Load(*env)
	if err != nil {
		log.Printf("Error loading %s file", *env)
	}

	environment := os.Getenv("ENV")

	dbConn := os.Getenv("DB_CONN")
	if dbConn == "" {
		return nil, fmt.Errorf("DB_CONN cannot be empty")
	}

	return &Config{
		Env:    environment,
		DbConn: dbConn,
	}, nil
}
