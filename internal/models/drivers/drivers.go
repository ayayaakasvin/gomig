package drivers

import (
	"github.com/ayayaakasvin/gomig/internal/models"
	"github.com/ayayaakasvin/gomig/internal/models/core/database"
	"github.com/ayayaakasvin/gomig/internal/models/drivers/mysql"
	"github.com/ayayaakasvin/gomig/internal/models/drivers/postgresql"
	"github.com/ayayaakasvin/gomig/internal/models/drivers/sqlite"
)

var availableDrivers = map[string]struct{}{
	postgresql.Driver: {},
	mysql.Driver:      {},
	sqlite3.Driver:    {},
}

func ValidDriver(driver string) bool {
	_, ok := availableDrivers[driver]
	return ok
}

var DriversMap map[string]func(*database.DatabaseConfig) models.DatabaseConnection = map[string]func(*database.DatabaseConfig) models.DatabaseConnection{
	postgresql.Driver: postgresql.NewPostgreSQL,
	mysql.Driver:      mysql.NewMySQL,
	sqlite3.Driver:    sqlite3.NewSQLite,
}

func AvailableDrivers() []string {
	res := []string{}
	for key, _ := range availableDrivers {
		res = append(res, key)
	}

	return res
}
