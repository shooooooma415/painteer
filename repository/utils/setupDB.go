package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// ConnectDB establishes a connection to the database.
func ConnectDB() (*sql.DB, error) {
	// Load .env file only in local development
	if os.Getenv("RENDER") == "" { // Assuming "RENDER" is set in the deployment environment
		if err := loadEnvFile(); err != nil {
			log.Printf("Warning: %v\n", err)
		}
	}

	// Fetch database connection information from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Validate the required environment variables
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("missing required database connection information in environment variables")
	}

	// Create the Data Source Name (DSN)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname,
	)

	// Open the database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Ping the database to verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	log.Println("Database connection successfully established.")
	return db, nil
}

// loadEnvFile loads the .env file from the root directory or its parent directories.
func loadEnvFile() error {
	rootDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %v", err)
	}

	for {
		envPath := fmt.Sprintf("%s/.env", rootDir)
		if _, err := os.Stat(envPath); err == nil {
			if err := godotenv.Load(envPath); err != nil {
				return fmt.Errorf("error loading .env file: %v", err)
			}
			return nil
		}

		parentDir := filepath.Dir(rootDir)
		if parentDir == rootDir {
			break
		}
		rootDir = parentDir
	}

	return fmt.Errorf(".env file not found")
}
