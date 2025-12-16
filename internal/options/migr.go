package options

import (
	"flag"

	"github.com/ayayaakasvin/gomig/internal/models/core/migration"
)

type migrFlags struct {
	path       string
	toLogAfter bool
}

func (m *migrFlags) Register(fs *flag.FlagSet) {
	fs.StringVar(&m.path, "path", "", "Path to migrations")
	fs.BoolVar(&m.toLogAfter, "log", false, "Enable JSON log file")
}

func (m *migrFlags) Normalize() {
	if m.path == "" {
		m.path = "./"
	}
}

func (m *migrFlags) Build(mtype string) *migration.MigrationConfig {
	return migration.NewMigrationConfig(
		migration.MigrationType(mtype),
		m.path,
		m.toLogAfter,
	)
}
