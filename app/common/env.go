package common

import (
	"errors"
	"os"
)

var availableEnvs map[string]bool

func init() {
	availableEnvs = map[string]bool{
		"dev":   true,
		"stage": true,
		"prod":  false,
	}
}

func GetEnv() (*string, error) {
	env := os.Getenv("APP_MODE")
	if env == "" {
		env = "dev"
	}

	if _, ok := availableEnvs[env]; !ok {
		return nil, errors.New("env does not exist or not activated")
	}

	return &env, nil
}
