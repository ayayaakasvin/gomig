package integration_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	epg *embeddedpostgres.EmbeddedPostgres
)
var binaryPath string

func TestMain(m *testing.M) {
	epg = embeddedpostgres.NewDatabase(
		embeddedpostgres.DefaultConfig().
			Port(5432).
			Database("testdb").
			Username("postgres").
			Password("postgres"),
	)

	if err := epg.Start(); err != nil {
		panic(err)
	}

	var err error
	db, err = sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres password=postgres dbname=testdb sslmode=disable",
	)
	if err != nil {
		panic(err)
	}
	binaryPath = filepath.Join("testbin", "gomig")

	if runtime.GOOS == "windows" {
		binaryPath += ".exe"
	}

	cmd := exec.Command(
		"go", "build",
		"-o", binaryPath,
		"../cmd/gomig",
	)

	if out, err := cmd.CombinedOutput(); err != nil {
		panic(fmt.Sprintf(
			"build failed: %v\n%s",
			err,
			out,
		))
	}

	code := m.Run()
	os.Exit(code)
}

func TestConnectivity(t *testing.T) {
	if err := db.Ping(); err != nil {
		t.Fatalf("db ping failed: %v", err)
	}
}

func TestMigrationUp(t *testing.T) {
	cmd := exec.Command(
		binaryPath,
		"up",
		"-driver=postgres",
		"-host=localhost",
		"-port=5432",
		"-user=postgres",
		"-password=postgres",
		"-dbname=testdb",
		"-sslmode=disable",
		fmt.Sprintf("-path=%s", filepath.Join(pathToMigration, PsqlUpMigrationFile)),
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf(
			"migration failed: %v\n%s",
			err,
			out,
		)
	}
}

func TestMigrationDown(t *testing.T) {
	cmd := exec.Command(
		binaryPath,
		"down",
		"-driver=postgres",
		"-host=localhost",
		"-port=5432",
		"-user=postgres",
		"-password=postgres",
		"-dbname=testdb",
		"-sslmode=disable",
		fmt.Sprintf("-path=%s", filepath.Join(pathToMigration, PsqlDownMigrationFile)),
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf(
			"migration failed: %v\n%s",
			err,
			out,
		)
	}
}
