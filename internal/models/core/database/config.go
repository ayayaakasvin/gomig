package database

import (
	"fmt"
	"net/url"
)

type DatabaseConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DBName   string `json:"dbname" yaml:"dbname"`
	SSLMode  string `json:"ssl" yaml:"ssl"`

	Driver string `json:"driver" yaml:"driver"`

	// URL is prebuilt connection URL for remote connection
	// As example: Render Database Instance: External database URL
	URL string
}

// New creates and returns a new DatabaseConfig with the specified connection parameters.
// Parameters:
//   - host: The database server host address
//   - port: The database server port number
//   - user: The database user name
//   - password: The database user password
//   - dbname: The name of the database to connect to
//   - sslmode: The SSL mode for the connection (e.g., "disable", "require", "verify-full")
//
// Returns:
//   - *DatabaseConfig: A pointer to the newly created DatabaseConfig struct
func NewDatabaseConfig(host string, port int, user, password, dbname, sslmode string) *DatabaseConfig {
	return &DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbname,
		SSLMode:  sslmode,
	}
}

func NewDatabaseConfigWithURL(URL string) *DatabaseConfig {
	return &DatabaseConfig{
		URL: URL,
	}
}

const (
	DefaultHost     = "localhost"
	DefaultPort     = 5432
	DefaultUser     = "user"
	DefaultPassword = ""
	DefaultDBName   = "default"
	DefaultSSLMode  = "disable"
)

func (dc DatabaseConfig) String() string {
	redactedURL := "<nil>"
	if dc.URL != "" {
		parsed, err := url.Parse(dc.URL)
		if err != nil {
			// fallback: hide completely if parsing fails
			redactedURL = "<invalid-url>"
		} else {
			if parsed.User != nil {
				// remove password, keep username if exists
				username := parsed.User.Username()
				parsed.User = url.User(username)
			}
			redactedURL = parsed.String()
		}
	}

	return fmt.Sprintf(
		"DatabaseConfig{Host: %q, Port: %d, User: %q, DBName: %q, SSLMode: %q, Driver: %q, URL: %q, Password: <hidden>}",
		dc.Host,
		dc.Port,
		dc.User,
		dc.DBName,
		dc.SSLMode,
		dc.Driver,
		redactedURL,
	)
}
