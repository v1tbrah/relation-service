package config

import "os"

const (
	defaultStorageHost     = "127.0.0.1"
	defaultStoragePort     = "7687"
	defaultStorageUser     = "neo4j"
	defaultStoragePassword = "password"
	defaultStorageDBName   = "neo4j"
)
const (
	envNameStorageHost     = "STORAGE_HOST"
	envNameStoragePort     = "STORAGE_PORT"
	envNameStorageUser     = "STORAGE_USER"
	envNameStoragePassword = "STORAGE_PASSWORD"
	envNameStorageDBName   = "STORAGE_DB_NAME"
)

type Storage struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func newDefaultStorageConfig() Storage {
	return Storage{
		Host:     defaultStorageHost,
		Port:     defaultStoragePort,
		User:     defaultStorageUser,
		Password: defaultStoragePassword,
		DBName:   defaultStorageDBName,
	}
}

func (c *Storage) parseEnv() {
	envHost := os.Getenv(envNameStorageHost)
	if envHost != "" {
		c.Host = envHost
	}

	envPort := os.Getenv(envNameStoragePort)
	if envPort != "" {
		c.Port = envPort
	}

	envUser := os.Getenv(envNameStorageUser)
	if envUser != "" {
		c.User = envUser
	}

	envPassword := os.Getenv(envNameStoragePassword)
	if envPassword != "" {
		c.Password = envPassword
	}

	envDBName := os.Getenv(envNameStorageDBName)
	if envDBName != "" {
		c.DBName = envDBName
	}
}
