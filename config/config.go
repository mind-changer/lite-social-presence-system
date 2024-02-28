package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Postgres *PostgresConfig
}

type PostgresConfig struct {
	Password string
	Username string
	Port     string
	Host     string
	Database string
}

func LoadConfigFromEnv(cfg *Config, envFilePath string) error {
	logrus.Info("loading env")
	if err := godotenv.Load(envFilePath); err != nil {
		log.Println(".env file not found. Defaulting to environment")
	}
	err := envconfig.Process("", cfg)
	if err != nil {
		logrus.Fatal(err.Error())
		return err
	}
	fmt.Println(*cfg)
	return nil
}
