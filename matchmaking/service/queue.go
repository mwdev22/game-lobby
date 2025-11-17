package service

import (
	"context"
	matchmakingpb "gen/matchmaking"
)

func (s *Service) JoinQueue(ctx context.Context, req *matchmakingpb.JoinQueueRequest) (*matchmakingpb.JoinQueueResponse, error) {
	return &matchmakingpb.JoinQueueResponse{}, nil
}

func (s *Service) LeaveQueue(ctx context.Context, req *matchmakingpb.LeaveQueueRequest) (*matchmakingpb.LeaveQueueResponse, error) {
	return &matchmakingpb.LeaveQueueResponse{}, nil
}

func (s *Service) GetQueueStatus(ctx context.Context, req *matchmakingpb.GetQueueStatusRequest) (*matchmakingpb.GetQueueStatusResponse, error) {
	return &matchmakingpb.GetQueueStatusResponse{}, nil
}
