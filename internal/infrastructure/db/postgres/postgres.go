package postgres

import (
	"database/sql"
	"fmt"

	"github.com/LidorAlmkays/MineServerForge/pkg/logger"
)

func InitializeDB(l logger.Logger, dbUser, dbPassword, dbHost, dbName string, dbPort int) (*sql.DB, error) {
	// Step 1: Connect to the default database (postgres)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/postgres?sslmode=disable", dbUser, dbPassword, dbHost, dbPort)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to PostgreSQL server: %w", err)
	}

	// Step 2: Create the database if it doesn't exist
	err = createDatabaseIfNotExist(l, db, dbName)
	if err != nil {
		return nil, fmt.Errorf("Failed to create database: %w", err.Error())
	}
	l.Info("Post database checked or created.")

	// Step 3: Reconnect to the actual database
	dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the specified database: %w", err)
	}

	l.Message(fmt.Sprintf("Successfully initialized postgres db"))

	return db, nil
}
func createDatabaseIfNotExist(l logger.Logger, db *sql.DB, dbName string) error {
	// Check if the database exists
	checkDB := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", dbName)
	var exists bool
	err := db.QueryRow(checkDB).Scan(&exists)

	// Check for no rows error (database does not exist)
	if err == sql.ErrNoRows {
		// If the database doesn't exist, create it
		_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
		if err != nil {
			return fmt.Errorf("error creating database: %w", err)
		}
		l.Info(fmt.Sprintf("Database %s created successfully", dbName))
		return nil
	}

	// Handle any other unexpected errors
	if err != nil {
		return fmt.Errorf("error checking database existence: %w", err)
	}

	// If database exists, nothing needs to be done
	l.Info(fmt.Sprintf("Database %s already exists", dbName))
	return nil
}
