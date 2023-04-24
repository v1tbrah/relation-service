package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/pet-pr-social-network/relation-service/pbapi"
)

func (a *API) AddFriend(ctx context.Context, req *pbapi.AddFriendRequest) (*pbapi.Empty, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, ErrEmptyRequest.Error())
	}

	err := a.storage.AddFriend(ctx, req.GetUserID(), req.GetFriendID())
	if err != nil {
		log.Error().Err(err).Int64("userID", req.GetUserID()).Int64("friendID", req.GetFriendID()).Msg("storage.AddFriend")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pbapi.Empty{}, nil
}
