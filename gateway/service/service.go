package service

import "gateway"

type Service struct {
	player      gateway.PlayerClient
	matchmaking gateway.MatchmakingClient
	session     gateway.SessionClient
}

func NewService(player gateway.PlayerClient, matchmaking gateway.MatchmakingClient, session gateway.SessionClient) *Service {
	return &Service{
		player:      player,
		matchmaking: matchmaking,
		session:     session,
	}
}
