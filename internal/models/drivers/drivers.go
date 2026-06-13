package drivers

import (
	"github.com/ayayaakasvin/gomig/internal/models"
	"github.com/ayayaakasvin/gomig/internal/models/core/database"
	"github.com/ayayaakasvin/gomig/internal/models/drivers/mysql"
	"github.com/ayayaakasvin/gomig/internal/models/drivers/postgresql"
	sqlite "github.com/ayayaakasvin/gomig/internal/models/drivers/sqlite"
)

var availableDrivers = map[string]struct{}{
	postgresql.Driver: {},
	mysql.Driver:      {},
	sqlite.Driver:    {},
}

func ValidDriver(driver string) bool {
	_, ok := availableDrivers[driver]
	return ok
}

var DriversMap map[string]func(*database.DatabaseConfig) models.DatabaseConnection = map[string]func(*database.DatabaseConfig) models.DatabaseConnection{
	postgresql.Driver: postgresql.NewPostgreSQL,
	mysql.Driver:      mysql.NewMySQL,
	sqlite.Driver:    sqlite.NewSQLite,
}

func AvailableDrivers() []string {
	res := []string{}
	for key := range availableDrivers {
		res = append(res, key)
	}

	return res
}
