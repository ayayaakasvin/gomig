package options

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ayayaakasvin/gomig/internal/models/core/database"
	"github.com/ayayaakasvin/gomig/internal/models/core/migration"
)

func ParseFlags() (*database.DatabaseConfig, *migration.MigrationConfig, error) {
	if len(os.Args) < 2 {
		return nil, nil, errors.New(InvalidCommand)
	}

	mtype := os.Args[1]
	if err := subcommand(mtype); err != nil {
		return nil, nil, err
	}

	fs := flag.NewFlagSet("gomig", flag.ExitOnError)

	dbf := &dbFlags{}
	mf := &migrFlags{}

	dbf.Register(fs)
	mf.Register(fs)

	if err := fs.Parse(os.Args[2:]); err != nil {
		return nil, nil, err
	}

	if err := dbf.Validate(); err != nil {
		return nil, nil, err
	}

	mf.Normalize()

	return dbf.Build(), mf.Build(mtype), nil
}

func printMigrateHelp() {
	fmt.Println(`Usage:
  gomig migrate [flags]

Description:
  Runs database migrations

Flags:
  --driver      Database driver (postgres, mysql, sqlite)
  --host        Database host
  --port        Database port
  --user        Database user
  --password    Database password
  --dbname      Database name, or in case of SQLite, the name of file or empty for cache database
  --sslmode     SSL mode
  --url         Full database URL (overrides flags)
  --path        Path to migration files
  --log         Enable JSON log file
	
  Example: gomig up -driver="postgres" -host"url-of-your-database-domain.example.com" -dbname="your-database-name" -port=5432 -user="your-database-username" -password="your-database-password" -sslmode=ssl_mode -path="./path_to_migration_files"`)
}

func subcommand(command string) error {
	switch command {
	case "help":
		printMigrateHelp()
		return errors.New(HelpCommand)
	case "up", "down":
		return nil
	default:
		fmt.Println("invalid/unknown command, try `gomig help`")
		return errors.New(InvalidCommand)
	}
}
