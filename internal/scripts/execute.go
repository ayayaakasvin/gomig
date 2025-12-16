package scripts

import (
	"fmt"

	"github.com/ayayaakasvin/gomig/internal/models"
)

// ExecuteScripts executes a series of SQL scripts against the provided database.
// It takes a Database interface and a slice of SQL script strings as input.
// The function executes each script sequentially using the database connection.
// If any script fails to execute, it logs the error and returns immediately with an error.
// Returns an error if the database connection is nil, if no scripts are provided,
// or if any script execution fails. Returns nil on successful execution of all scripts.
func ExecuteScripts(db models.DatabaseConnection, scripts []string) error {
	if db == nil || len(scripts) == 0 {
		return fmt.Errorf("Empty args: %v, %v", db, scripts)
	}

	for _, script := range scripts {
		_, err := db.Exec(script)
		if err != nil {
			return fmt.Errorf("failed to execute script: %v", err)
		}
	}

	return nil
}
