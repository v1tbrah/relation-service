package storage

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/pkg/errors"

	"github.com/v1tbrah/relation-service/internal/model"
)

const getFriendsQueryTmpl = `MATCH (u1:User {id: $u1ID})-[:FRIEND]->(friends)`

func (s *Storage) GetFriends(ctx context.Context, userID int64, direction model.Direction, userOffsetID, limit int64) (friends []int64, err error) {
	getFriendsQuery := getFriendsQueryTmpl

	switch direction {
	case model.First:
	case model.Next:
		getFriendsQuery += fmt.Sprintf(" WHERE friends.id > %d", userOffsetID)
	case model.Prev:
		getFriendsQuery += fmt.Sprintf(" WHERE friends.id < %d", userOffsetID)
	default:
		return nil, fmt.Errorf("get friends: %w", model.ErrInvalidDirection)
	}

	getFriendsQuery += " RETURN friends"

	switch direction {
	case model.First, model.Next:
		getFriendsQuery += ` ORDER BY friends.id ASC`
	case model.Prev:
		getFriendsQuery += " ORDER BY friends.id DESC"
	default:
		return nil, fmt.Errorf("get friends: %w", model.ErrInvalidDirection)
	}

	if limit > 0 {
		getFriendsQuery += fmt.Sprintf(" LIMIT %d", limit)
	}

	result, err := s.dbSes.Run(ctx, getFriendsQuery, map[string]any{"u1ID": userID})
	if err != nil {
		return nil, err
	}

	for result.Next(ctx) {
		record := result.Record()

		rAny, ok := record.Get("friends")
		if !ok {
			continue
		}

		rNode, ok := rAny.(dbtype.Node)
		if !ok {
			continue
		}

		props := rNode.GetProperties()
		if props == nil {
			continue
		}

		idAny, ok := props["id"]
		if !ok {
			continue
		}

		id, ok := idAny.(int64)
		if !ok {
			continue
		}

		friends = append(friends, id)
	}

	if err = result.Err(); err != nil {
		return nil, errors.Wrap(err, "result.Err")
	}

	return friends, nil
}
