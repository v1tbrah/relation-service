package api

import (
	"context"

	"gitlab.com/pet-pr-social-network/relation-service/internal/model"
)

//go:generate mockery --name Storage
type Storage interface {
	AddFriend(ctx context.Context, userID, friendID int64) error
	RemoveFriend(ctx context.Context, userID, friendID int64) error
	GetFriends(ctx context.Context, userID int64, direction model.Direction, userOffsetID, limit int64) (friends []int64, err error)
}

//go:generate mockery --name FriendMsgSender
type FriendMsgSender interface {
	SendMsgFriendAdded(userID, friendID int64)
	SendMsgFriendRemoved(userID, friendID int64)
}
