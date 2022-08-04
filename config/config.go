package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type Config struct {
	LogLevel         string
	PgURL            string
	PgMigrationsPath string
	HTTPAddr         string
	SigningKey       string
}

var (
	config *Config
	once   sync.Once
)

func New() *Config {
	return &Config{
		LogLevel:         getEnv("LOG_LEVEL", "debug"),
		PgURL:            getEnv("PG_URL", "host=localhost dbname=todo_api sslmode=disable user=admin password=1111"),
		PgMigrationsPath: getEnv("PG_MIGRATIONS_PATH", "../store/migrations"),
		HTTPAddr:         getEnv("HTTP_ADDR", "localhost:8080"),
		SigningKey:       getEnv("SIGNING_KEY", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Get reads config from environment. Once.
func Get() *Config {

	once.Do(func() {
		config = New()
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return config
}
