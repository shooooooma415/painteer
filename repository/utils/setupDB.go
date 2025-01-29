package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func findEnvFile() (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %v", err)
	}
	for {
		envPath := filepath.Join(rootDir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			return envPath, nil
		}

		parentDir := filepath.Dir(rootDir)
		if parentDir == rootDir {
			break
		}
		rootDir = parentDir
	}

	return "", errors.New(".env file not found")
}

func ConnectDB() (*sql.DB, error) {
	envPath, err := findEnvFile()
	if err != nil {
		return nil, fmt.Errorf("failed to locate .env file: %v", err)
	}

	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("missing required database connection information in environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	log.Println("Database connection successfully established.")
	return db, nil
}
