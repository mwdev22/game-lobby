package service

import (
	"context"
	"gateway"
)

func (s *Service) JoinQueue(ctx context.Context, req *gateway.JoinQueueRequest) (*gateway.QueueJoinResult, error) {
	return s.matchmaking.JoinQueue(ctx, req)
}
