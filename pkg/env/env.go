package env

import (
	"github.com/joho/godotenv"
)

var Env map[string]string

type Environment string

const (
	Production  Environment = "production"
	Development Environment = "development"
)

func GetEnv(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func SetupEnvFile(envScope Environment) {
	var envFile string

	switch envScope {
	case Production:
		envFile = ".env.production"
	case Development:
		envFile = ".env.development"
	default:
		envFile = ".env"
	}

	var err error

	Env, err = godotenv.Read(envFile)
	if err != nil {
		panic(err)
	}
}
