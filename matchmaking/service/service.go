package service

import (
	matchmakingpb "gen/matchmakingpb"

	"matchmaking"
)

type Service struct {
	matchmakingpb.UnimplementedMatchmakingServiceServer
	queue matchmaking.Queue
	match matchmaking.MatchMaker
}

func New() *Service {
	return &Service{}
}
