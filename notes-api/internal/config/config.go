package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	MongoDB string
	ServerPort string
}

func Load() (Config, error) {
	// `godotenv.Load()` reads .env and sets them into the process env.
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("Failed to load .env")
	}

	mongoURI, err := extractEnv("MONGO_URI")
	if err != nil {
		return Config{}, err
	}

	mongoDB, err := extractEnv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}

	serverPort, err := extractEnv("PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoURI: mongoURI,
		MongoDB: mongoDB,
		ServerPort: serverPort,
	}, nil
}

func extractEnv(key string) (string, error) {
	// `os.Getenv()` gets the val of the key in the .env file.
	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("Missing req env")
	}

	return val, nil
}