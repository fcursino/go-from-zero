package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvKey string

const (
	Env                     EnvKey = "ENV"
	JwtSecret               EnvKey = "JWT_SECRET"
	PostgresUser            EnvKey = "POSTGRES_USER"
	PostgresPassword        EnvKey = "POSTGRES_PASSWORD"
	PostgresDefaultDatabase EnvKey = "POSTGRES_DB"
	PostgresDatabaseAddress EnvKey = "POSTGRES_DB_ADDRESS"
	PostgresDatabasePort    EnvKey = "POSTGRES_DB_PORT"
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

func IsDevelopmentMode() bool {
	return Env.GetValue() == EnvDevelopment
}

func (key EnvKey) GetValue() string {
	return os.Getenv(string(key))
}

func Load() error {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return err
}
