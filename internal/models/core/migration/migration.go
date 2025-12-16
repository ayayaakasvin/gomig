package migration

import "fmt"

type MigrationType string

const (
	Up      MigrationType = "up"
	Down    MigrationType = "down"
	Unknown MigrationType = "unknown"
)

type MigrationConfig struct {
	MigrationType
	SourcePath string
	ToLogAfter bool
}

// NewMigrationConfig creates a new MigrationConfig with the specified migration type and source path.
// It returns a pointer to the MigrationConfig instance.
//
// Parameters:
//   - migrationType: The type of migration to be performed
//   - path: The source path for the migration
//
// Returns:
//   - *MigrationConfig: A pointer to the newly created MigrationConfig
func NewMigrationConfig(migrationType MigrationType, path string, toLogAfter bool) *MigrationConfig {
	return &MigrationConfig{
		SourcePath:    path,
		MigrationType: migrationType,
		ToLogAfter: toLogAfter,
	}
}

func (mc *MigrationConfig) String() string {
	return fmt.Sprintf("MigrationConfig{Type: %s, SourcePath: %s, ToLogAfter: %t}", mc.MigrationType, mc.SourcePath, mc.ToLogAfter)
}