package service

import (
	"context"
	matchmakingpb "gen/matchmaking"
)

func (s *Service) CancelMatch(ctx context.Context, req *matchmakingpb.CancelMatchRequest) (*matchmakingpb.CancelMatchResponse, error) {
	return &matchmakingpb.CancelMatchResponse{}, nil
}

func (s *Service) StreamMatchFound(ctx context.Context, req *matchmakingpb.MatchFoundRequest) (*matchmakingpb.MatchInfo, error) {
	return &matchmakingpb.MatchInfo{}, nil
}
