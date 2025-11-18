package service

import (
	"context"

	matchmakingpb "gen/matchmakingpb"
)

func (s *Service) CancelMatch(ctx context.Context, req *matchmakingpb.CancelMatchRequest) (*matchmakingpb.CancelMatchResponse, error) {
	return &matchmakingpb.CancelMatchResponse{}, nil
}
