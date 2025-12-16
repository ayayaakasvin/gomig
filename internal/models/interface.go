package models

import "database/sql"

// DatabaseConnection represents an interface for database operations.
// It provides methods for managing database connections and executing queries.
// This interface abstracts away the underlying database implementation,
// allowing for different database engines to be used interchangeably.
// Implementations must provide connection management, query execution,
// and transaction support capabilities.
type DatabaseConnection interface {
	ConnectionString() string

	Open() error
	Ping() error
	Close() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecTx(queries []string, args [][]interface{}) error
}
