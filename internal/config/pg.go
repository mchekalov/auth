package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	pgHostEnvName  = "PG_HOST"
	pgPortEnvName  = "PG_PORT"
	pgDataBaseName = "PG_DATABASE_NAME"
	pgUser         = "PG_USER"
	pgPassword     = "PG_PASSWORD"
	pgSslMode      = "PG_SSLMODE"
)

// PGConfig defines the interface for PostgreSQL configuration.
type PGConfig interface {
	DsnString() string
}

type pgConfig struct {
	host     string
	port     string
	basename string
	user     string
	password string
	sslmode  string
}

// NewPGConfig creates a new instance of PGConfig using environment variables.
// It returns an error if any required environmental variable is not set.
func NewPGConfig() (PGConfig, error) {
	host := os.Getenv(pgHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("postgres host not found")
	}
	port := os.Getenv(pgPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("postgres port not found")
	}
	dbname := os.Getenv(pgDataBaseName)
	if len(dbname) == 0 {
		return nil, errors.New("postgres basename not found")
	}
	user := os.Getenv(pgUser)
	if len(user) == 0 {
		return nil, errors.New("postgres user not found")
	}
	pass := os.Getenv(pgPassword)
	if len(pass) == 0 {
		return nil, errors.New("postgres password not found")
	}
	mode := os.Getenv(pgSslMode)
	if len(mode) == 0 {
		return nil, errors.New("postgres sslmode parametr not found")
	}

	return &pgConfig{
		host:     host,
		port:     port,
		basename: dbname,
		user:     user,
		password: pass,
		sslmode:  mode,
	}, nil
}

// DsnString constructs the connection string for PostgreSQL using configuration parameters.
func (cfg *pgConfig) DsnString() string {
	return fmt.Sprintf("host=%v port= %v dbname=%v user=%v password=%v sslmode=%v",
		cfg.host, cfg.port, cfg.basename, cfg.user, cfg.password, cfg.sslmode)
}
