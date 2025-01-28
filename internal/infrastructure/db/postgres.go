package db

import (
	"fmt"
	"strings"

	"github.com/LidorAlmkays/MineServerForge/queries"
	"github.com/jmoiron/sqlx"
)

const PostgresQueryDirectory = "postgres_basic_queries/"

func (f *dbFactory) initializePostgresDB(dbUser, dbPassword, dbHost, dbName string, dbPort int) (*sqlx.DB, error) {
	// Step 1: Connect to the default database (postgres)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/postgres?sslmode=disable", dbUser, dbPassword, dbHost, dbPort)
	var err error
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to PostgreSQL server: %w", err)
	}
	// Step 2: Check if the database exists
	var exists bool
	err = db.Get(&exists, "SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)", dbName)
	if err != nil {
		return nil, fmt.Errorf("Unable to check if database exists: %w", err)
	}

	// Step 3: Create the database if it doesn't exist
	if !exists {
		f.l.Info("Database was not found, creating database")
		err := f.createPostgresDatabase(db, dbName)
		if err != nil {
			return nil, err
		}
	}

	// Step 4: Reconnect to the actual database
	dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the specified database: %w", err)
	}

	f.l.Message(fmt.Sprintf("Successfully initialized postgres db"))

	return db, nil
}

func (f *dbFactory) createPostgresDatabase(db *sqlx.DB, dbName string) error {
	queryTemplate, err := queries.GetQuery("postgres_basic_queries/create_database.sql")
	if err != nil {
		return err
	}

	query := strings.ReplaceAll(queryTemplate, ":dbName", dbName)

	_, err = db.Exec(query)
	if err != nil {
		err = fmt.Errorf("Failed to execute query:", err)
		return err
	}
	f.l.Message(fmt.Sprintf("Successfully created the database, named: %s.", dbName))

	return nil
}
