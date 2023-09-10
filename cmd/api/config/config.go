package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoURI string
}

// LoadEnv loads the environment variables from a .env file
func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

// Returns a MongoDB client connection
func getMongoDBConnection() (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGOURI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}
	return client, nil
}

// GetConfig loads the aplication config from .env
func GetConfig() (*Config, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	mongoURI := os.Getenv("MONGOURI")

	config := &Config{
		MongoURI: mongoURI,
	}

	return config, nil
}
