package main

import (
	"os"

	"github.com/ayayaakasvin/gomig/internal/models/drivers"
	"github.com/ayayaakasvin/gomig/internal/options"
	"github.com/ayayaakasvin/gomig/internal/scripts"
	slogger "github.com/ayayaakasvin/gomig/pkg/logger"
)

func main() {
	dbconf, migrconf, err := options.ParseFlags()
	if err != nil {
		if err.Error() == options.HelpCommand {
			os.Exit(0)
		}
		logger := slogger.NewLogger(false)
		logger.Error("parse flag error", "error", err)
		os.Exit(1)
	}

	logger := slogger.NewLogger(migrconf.ToLogAfter)

	logger.Info("Database configuration", "db_cfg", dbconf.String())
	logger.Info("Migration configuration", "migr_cfg", migrconf.String())

	db := drivers.DriversMap[dbconf.Driver](dbconf)
	defer db.Close()

	if err := db.Open(); err != nil {
		logger.Error("failed to open database connection", "err", err)
		os.Exit(1)
	}

	sqlScripts, err := scripts.ParseMigrationFiles(migrconf)
	if err != nil {
		logger.Error("failed to parse scripts", "error", err)
		os.Exit(1)
	}

	logger.Info("Number of SQL scripts", "sql_scripts_count", len(sqlScripts))

	err = scripts.ExecuteScripts(db, sqlScripts)
	if err != nil {
		logger.Error("failed to execute scripts", "error", err)
		os.Exit(1)
	}

	logger.Info("Migration completed successfully")
	os.Exit(0)
}
