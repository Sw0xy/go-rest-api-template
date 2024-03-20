package bootstrap

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	env "github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Env struct {
	MongoUri               string `env:"MONGO_URI"`
	AccessTokenSecret      string `env:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `env:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpiryHour  int    `env:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `env:"REFRESH_TOKEN_EXPIRY_HOUR"`
}

func InitEnv() *Env {
	env, err := LoadEnv()
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return env
}

func LoadEnv() (*Env, error) {
	cfg := Env{}
	currentDir, err := os.Getwd()
	if err != nil {
		return &Env{}, errors.New(fmt.Sprintf("Error getting current working directory: %s", err))
	}

	envFiles, err := filepath.Glob(filepath.Join(currentDir, "*.env"))
	if err != nil {
		return &Env{}, errors.New(fmt.Sprintf("Error listing .env files: %s", err))
	}

	// Load .env file
	err = godotenv.Load(envFiles...)
	if err != nil {
		return &Env{}, errors.New(fmt.Sprintf("Error loading .env file: %s", err))
	}

	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg, nil
}
