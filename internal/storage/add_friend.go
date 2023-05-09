package storage

import (
	"context"

	"github.com/pkg/errors"
)

const (
	addFriendQuery = `
MERGE (u1:User {id: $u1ID})
MERGE (u2:User {id: $u2ID})
MERGE (u1)-[:FRIEND]->(u2)
`
)

func (s *Storage) AddFriend(ctx context.Context, userID, friendID int64) error {
	_, err := s.dbSes.Run(ctx, addFriendQuery, map[string]any{"u1ID": userID, "u2ID": friendID})
	if err != nil {
		return errors.Wrapf(err, "dbSes.Run addFriendQuery, userID %d, friendID %d", userID, friendID)
	}

	return nil
}
