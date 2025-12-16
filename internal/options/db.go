package options

import (
	"errors"
	"flag"

	"github.com/ayayaakasvin/gomig/internal/models/core/database"
	"github.com/ayayaakasvin/gomig/internal/models/drivers"
)

type dbFlags struct {
	driver   string
	host     string
	port     int
	user     string
	password string
	dbname   string
	sslmode  string
	url      string
}

func (d *dbFlags) Register(fs *flag.FlagSet) {
	fs.StringVar(&d.driver, "driver", "", "Database driver")
	fs.StringVar(&d.host, "host", database.DefaultHost, "Database host")
	fs.IntVar(&d.port, "port", database.DefaultPort, "Database port")
	fs.StringVar(&d.user, "user", database.DefaultUser, "Database user")
	fs.StringVar(&d.password, "password", database.DefaultPassword, "Database password")
	fs.StringVar(&d.dbname, "dbname", database.DefaultDBName, "Database name")
	fs.StringVar(&d.sslmode, "sslmode", database.DefaultSSLMode, "SSL mode")
	fs.StringVar(&d.url, "url", "", "Database URL (overrides other flags)")
}

func (d *dbFlags) Validate() error {
	if d.driver == "" {
		return errors.New(DriverIsNotSpecified)
	}
	if !drivers.ValidDriver(d.driver) {
		return errors.New(InvalidDriver)
	}
	return nil
}

func (d *dbFlags) Build() *database.DatabaseConfig {
	if d.url != "" {
		cfg := database.NewDatabaseConfigWithURL(d.url)
		cfg.Driver = d.driver
		return cfg
	}

	cfg := database.NewDatabaseConfig(
		d.host,
		d.port,
		d.user,
		d.password,
		d.dbname,
		d.sslmode,
	)
	cfg.Driver = d.driver
	return cfg
}
