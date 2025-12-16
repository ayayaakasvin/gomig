package database

import (
	"database/sql"
	"errors"
)

const NilPointerError string = "nil pointer database connection"

// BaseDatabase represents the configuration required to connect to a database.
type BaseDatabase struct {
	Config *DatabaseConfig // Config holds the database configuration.
	DB     *sql.DB                // DB is the sql.DB object for database connection.
}

// Open opens a new database connection and assigns it to the DB field.
func (ps *BaseDatabase) Open(driver, connStr string) error {
	db, err := sql.Open(ps.Config.Driver, connStr)
	if err != nil {
		return err
	}
	ps.DB = db
	return nil
}

// Close closes the database connection.
func (ps *BaseDatabase) Close() error {
	if ps.DB != nil {
		return ps.DB.Close()
	}
	return nil
}

// Ping verifies a connection to the database is still alive
func (ps *BaseDatabase) Ping() error {
	if ps.DB == nil {
		return errors.New(NilPointerError)
	}
	return ps.DB.Ping()
}

// Query executes a query that returns rows, typically a SELECT.
func (ps *BaseDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return ps.DB.Query(query, args...)
}

// Exec executes a query without returning any rows, typically an INSERT, UPDATE, or DELETE.
func (ps *BaseDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	return ps.DB.Exec(query, args...)
}

// ExecTx executes a series of queries within a transaction.
func (ps *BaseDatabase) ExecTx(queries []string, args [][]interface{}) error {
	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}

	for i, query := range queries {
		if _, err := tx.Exec(query, args[i]...); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
