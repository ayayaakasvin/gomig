package postgresql

import (
	"fmt"

	"github.com/ayayaakasvin/gomig/internal/models"
	"github.com/ayayaakasvin/gomig/internal/models/core/database"
	_ "github.com/lib/pq"
)

const Driver = "postgres"

type PostgresSql struct {
	database.BaseDatabase
}

func NewPostgreSQL(cfg *database.DatabaseConfig) models.DatabaseConnection {
	return &PostgresSql{
		BaseDatabase: database.BaseDatabase{
			Config: cfg,
		},
	}
}

func (ps *PostgresSql) ConnectionString() string {
	if ps.Config.URL != "" {
		return ps.Config.URL
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", ps.Config.Host, ps.Config.Port, ps.Config.User, ps.Config.Password, ps.Config.DBName, ps.Config.SSLMode)
}

func (p *PostgresSql) Open() error {
	return p.BaseDatabase.Open(Driver, p.ConnectionString())
}
