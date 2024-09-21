package configs

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBName                 string
	DBPort                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

var Envs = initConfig()

func initConfig() Config {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	envPath := filepath.Join(workingDir, ".env")
	godotenv.Load(envPath)

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "fallback"),
		DBPassword:             getEnv("DB_PASSWORD", "fallback"),
		DBName:                 getEnv("DB_NAME", "fallback"),
		DBPort:                 getEnv("DB_PORT", "fallback"),
		JWTSecret:              getEnv("JWT_SECRET", "fallback"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
