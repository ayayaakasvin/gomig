package core

import (
	"github.com/ayayaakasvin/gomig/internal/models/core/database"
	"github.com/ayayaakasvin/gomig/internal/models/core/migration"
)

type Options struct {
	Database  *database.DatabaseConfig   `json:"db_cfg" yaml:"db_cfg"`
	Migration *migration.MigrationConfig `json:"-" yaml:"-"`
}
