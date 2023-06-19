package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

const (
	defaultLogLvl = zerolog.InfoLevel
	envNameLogLvl = "LOG_LVL"
)

type Config struct {
	GRPC    GRPC
	HTTP    HTTP
	Storage Storage
	Kafka   Kafka
	LogLvl  zerolog.Level
}

func NewDefaultConfig() Config {
	return Config{
		GRPC:    newDefaultGRPCConfig(),
		HTTP:    newDefaultHTTPConfig(),
		Storage: newDefaultStorageConfig(),
		Kafka:   newDefaultKafkaConfig(),
		LogLvl:  defaultLogLvl,
	}
}

func (c *Config) ParseEnv() error {
	c.GRPC.parseEnv()

	c.HTTP.parseEnv()

	c.Storage.parseEnv()

	c.Kafka.parseEnv()

	if err := c.parseEnvLogLvl(); err != nil {
		return err
	}

	return nil
}

func (c *Config) parseEnvLogLvl() error {
	envLogLvl := os.Getenv(envNameLogLvl)
	if envLogLvl != "" {
		logLevel, err := zerolog.ParseLevel(envLogLvl)
		if err != nil {
			return fmt.Errorf("parse log lvl: %s", envLogLvl)
		}
		c.LogLvl = logLevel
	}
	return nil
}
