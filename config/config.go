package config

import (
	"database/sql"
	"errors"
	"log"
	"os"
	_ "github.com/lib/pq"
)

type Config struct {
	ServerPort 	string
	Mode 		string
	Version 	string
}

type JWTConfig struct {
	Secret 		string
	Iss			string
}

func LoadConfig() *Config {
	cfg := &Config{
		ServerPort: getEnv("SERVER_PORT", ":8000"),
		Mode: 		getEnv("GIN_MODE", "debug"),
		Version: 	getEnv("API_VERSION", "v1"),	
	}

	return cfg
}

func getEnv (key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func GetDBConnection() (*sql.DB, error) {
	connStr := getEnv("DATABASE_CONN", "")
	if (connStr == "") {
		return nil, errors.New("connection string not found")
	}

	db, err := sql.Open("postgres", connStr)
	if (err != nil) {
		log.Fatal(err)
	}

	return db, nil
}

func GetJWTConfig() *JWTConfig {
	cfg := &JWTConfig{
		Secret: getEnv("JWT_SECRET", ""),
		Iss: 	getEnv("JWT_ISS", "localhost.com"),
	}
	
	return cfg
}