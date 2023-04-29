package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"gitlab.com/pet-pr-social-network/relation-service/rpbapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/pet-pr-social-network/relation-service/internal/model"
)

func (a *API) GetFriends(ctx context.Context, req *rpbapi.GetFriendsRequest) (*rpbapi.GetFriendsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, ErrEmptyRequest.Error())
	}

	friends, err := a.storage.GetFriends(ctx, req.GetUserID(), model.Direction(req.GetDirection()), req.GetUserOffsetID(), req.GetLimit())
	if err != nil {
		log.Error().Err(err).
			Int64("userID", req.GetUserID()).
			Int32("direction", int32(req.GetDirection())).
			Int64("userOffsetID", req.GetUserOffsetID()).
			Int64("limit", req.GetLimit()).
			Msg("storage.GetFriends")
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &rpbapi.GetFriendsResponse{Friends: friends}, nil
}
