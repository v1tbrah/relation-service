package storage

import (
	"context"

	"github.com/pkg/errors"
)

const (
	removeFriendQuery = `
MATCH (u1:User {id:$u1ID})-[f:FRIEND]-(u2:User {id: $u2ID})
DELETE f
`
)

func (s *Storage) RemoveFriend(ctx context.Context, userID, friendID int64) error {
	_, err := s.dbSes.Run(ctx, removeFriendQuery, map[string]any{"u1ID": userID, "u2ID": friendID})
	if err != nil {
		return errors.Wrapf(err, "dbSes.Run removeFriendQuery, userID %d, friendID %d", userID, friendID)
	}

	return nil
}
