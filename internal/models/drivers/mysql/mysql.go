package mysql

import (
	"fmt"

	"github.com/ayayaakasvin/gomig/internal/models"
	"github.com/ayayaakasvin/gomig/internal/models/core/database"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Driver = "mysql"
)

type MySQL struct {
	database.BaseDatabase
}

func NewMySQL(cfg *database.DatabaseConfig) models.DatabaseConnection {
	return &MySQL{
		BaseDatabase: database.BaseDatabase{
			Config: cfg,
		},
	}
}

func (m *MySQL) ConnectionString() string {
	if m.Config.URL != "" {
		return m.Config.URL
	}

	// user:password@tcp(host:port)/dbname?param=value
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?tls=preferred",
		m.Config.User,
		m.Config.Password,
		m.Config.Host,
		m.Config.Port,
		m.Config.DBName,
	)
}

func (m *MySQL) Open() error {
	return m.BaseDatabase.Open(Driver, m.ConnectionString())
}
