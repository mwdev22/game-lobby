package service

import (
	matchmakingpb "gen/matchmaking"

	"github.com/mwdev22/matchmaking"
)

type Service struct {
	matchmakingpb.UnimplementedMatchmakingServiceServer
	queue matchmaking.Queue
	match matchmaking.MatchMaker
}

func New() *Service {
	return &Service{}
}
