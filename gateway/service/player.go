package service

import (
	"context"
	"gateway"
)

func (s *Service) CreatePlayer(ctx context.Context, req *gateway.CreatePlayerRequest) (string, error) {
	id, err := s.player.CreatePlayer(ctx, req)
	if err != nil {
		return "", gateway.ErrInternal(err)
	}
	return id, nil
}

func (s *Service) GetPlayer(ctx context.Context, playerID string) (*gateway.PlayerProfile, error) {
	player, err := s.player.GetPlayer(ctx, playerID)
	if err != nil {
		return nil, gateway.ErrInternal(err)
	}
	if player == nil {
		return nil, gateway.ErrNotFound("player")
	}
	return player, nil
}

func (s *Service) UpdatePlayer(ctx context.Context, req *gateway.UpdatePlayerRequest) error {
	err := s.player.UpdatePlayer(ctx, req)
	if err != nil {
		return gateway.ErrInternal(err)
	}
	return nil
}

func (s *Service) GetPlayerStats(ctx context.Context, playerID string) (*gateway.PlayerStats, error) {
	stats, err := s.player.GetPlayerStats(ctx, playerID)
	if err != nil {
		return nil, gateway.ErrInternal(err)
	}
	return stats, nil
}

func (s *Service) SearchPlayers(ctx context.Context, query string, limit, offset int) ([]gateway.PlayerProfile, error) {
	players, err := s.player.SearchPlayers(ctx, query, limit, offset)
	if err != nil {
		return nil, gateway.ErrInternal(err)
	}
	return players, nil
}
