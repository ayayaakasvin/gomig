package integration_test

import (
	"path"
)

var (
	pathToMigration = ""
)

const (
	PsqlUpMigrationFile   = "0001_psql.up.sql"
	PsqlDownMigrationFile = "0001_psql.down.sql"
)

func init() {
	pathToMigration = path.Join(".", "testdata")
}
