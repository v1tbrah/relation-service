package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"gitlab.com/pet-pr-social-network/relation-service/rpbapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *API) AddFriend(ctx context.Context, req *rpbapi.AddFriendRequest) (*rpbapi.Empty, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, ErrEmptyRequest.Error())
	}

	err := a.storage.AddFriend(ctx, req.GetUserID(), req.GetFriendID())
	if err != nil {
		log.Error().Err(err).Int64("userID", req.GetUserID()).Int64("friendID", req.GetFriendID()).Msg("storage.AddFriend")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &rpbapi.Empty{}, nil
}
