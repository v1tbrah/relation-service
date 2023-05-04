package storage

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/pkg/errors"
	"gitlab.com/pet-pr-social-network/relation-service/internal/config"
)

type Storage struct {
	dbSes neo4j.SessionWithContext

	cfg config.StorageConfig
}

func Init(cfg config.StorageConfig) (*Storage, error) {
	newStorage := &Storage{cfg: cfg}

	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		"neo4j://"+cfg.Host+":"+cfg.Port,
		neo4j.BasicAuth(cfg.User, cfg.Password, ""),
	)
	if err != nil {
		return nil, errors.Wrap(err, "neo4j.NewDriverWithContext")
	}

	if err = driver.VerifyConnectivity(ctx); err != nil {
		return nil, errors.Wrap(err, "driver.VerifyConnectivity")
	}

	newStorage.dbSes = driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: cfg.DBName,
	})

	return newStorage, nil
}

func (s *Storage) Close(ctx context.Context) (err error) {
	if err = s.dbSes.Close(ctx); err != nil {
		return errors.Wrap(err, "close db session")
	}

	return nil
}