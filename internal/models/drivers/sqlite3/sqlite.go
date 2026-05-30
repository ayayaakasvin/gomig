package sqlite3

import (
	"github.com/ayayaakasvin/gomig/internal/models"
	"github.com/ayayaakasvin/gomig/internal/models/core/database"

	_ "modernc.org/sqlite"
)

const (
	Driver           = "sqlite3"
	inMemoryLocation = ":memory:"
)

type SQLite struct {
	database.BaseDatabase
}

func NewSQLite(cfg *database.DatabaseConfig) models.DatabaseConnection {
	return &SQLite{
		BaseDatabase: database.BaseDatabase{
			Config: cfg,
		},
	}
}

func (s3 *SQLite) ConnectionString() string {
	if s3.Config.URL != "" {
		return inMemoryLocation
	}

	return s3.Config.DBName
}

func (s3 *SQLite) Open() error {
	return s3.BaseDatabase.Open(Driver, s3.ConnectionString())
}
