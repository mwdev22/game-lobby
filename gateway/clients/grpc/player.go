package grpc

import (
	"context"
	"gateway"

	pb "gen/playerpb"

	"google.golang.org/grpc"
)

type playerClient struct {
	client pb.PlayerServiceClient
}

func NewPlayerClient(conn *grpc.ClientConn) gateway.PlayerClient {
	return &playerClient{
		client: pb.NewPlayerServiceClient(conn),
	}
}

func (c *playerClient) CreatePlayer(ctx context.Context, req *gateway.CreatePlayerRequest) (*gateway.PlayerProfile, error) {
	pbReq := &pb.CreatePlayerRequest{
		Name:  req.Name,
		Email: req.Email,
	}

	pbResp, err := c.client.CreatePlayer(ctx, pbReq)
	if err != nil {
		return nil, err
	}

	return &gateway.PlayerProfile{
		PlayerID:  pbResp.PlayerId,
		Name:      pbResp.Name,
		SkillRank: pbResp.SkillRank,
		CreatedAt: pbResp.CreatedAt,
	}, nil
}

func (c *playerClient) GetPlayer(ctx context.Context, playerID string) (*gateway.PlayerProfile, error) {
	pbReq := &pb.GetPlayerRequest{
		PlayerId: playerID,
	}

	pbResp, err := c.client.GetPlayer(ctx, pbReq)
	if err != nil {
		return nil, err
	}

	if pbResp.Player == nil {
		return nil, nil
	}

	return &gateway.PlayerProfile{
		PlayerID:     pbResp.Player.PlayerId,
		Name:         pbResp.Player.Name,
		Email:        pbResp.Player.Email,
		SkillRank:    pbResp.Player.SkillRank,
		TotalMatches: pbResp.Player.TotalMatches,
		Wins:         pbResp.Player.Wins,
		Losses:       pbResp.Player.Losses,
		CreatedAt:    pbResp.Player.CreatedAt,
		LastOnline:   pbResp.Player.LastOnline,
	}, nil
}

func (c *playerClient) UpdatePlayer(ctx context.Context, req *gateway.UpdatePlayerRequest) (*gateway.PlayerProfile, error) {
	pbReq := &pb.UpdatePlayerRequest{
		PlayerId: req.PlayerID,
	}

	if req.Name != nil {
		pbReq.Name = req.Name
	}
	if req.SkillRank != nil {
		pbReq.SkillRank = req.SkillRank
	}

	pbResp, err := c.client.UpdatePlayer(ctx, pbReq)
	if err != nil {
		return nil, err
	}

	if pbResp.Player == nil {
		return nil, nil
	}

	return &gateway.PlayerProfile{
		PlayerID:     pbResp.Player.PlayerId,
		Name:         pbResp.Player.Name,
		Email:        pbResp.Player.Email,
		SkillRank:    pbResp.Player.SkillRank,
		TotalMatches: pbResp.Player.TotalMatches,
		Wins:         pbResp.Player.Wins,
		Losses:       pbResp.Player.Losses,
		CreatedAt:    pbResp.Player.CreatedAt,
		LastOnline:   pbResp.Player.LastOnline,
	}, nil
}

func (c *playerClient) GetPlayerStats(ctx context.Context, playerID string) (*gateway.PlayerStats, error) {
	pbReq := &pb.GetPlayerStatsRequest{
		PlayerId: playerID,
	}

	pbResp, err := c.client.GetPlayerStats(ctx, pbReq)
	if err != nil {
		return nil, err
	}

	if pbResp.Stats == nil {
		return nil, nil
	}

	return &gateway.PlayerStats{
		TotalMatches:        pbResp.Stats.TotalMatches,
		Wins:                pbResp.Stats.Wins,
		Losses:              pbResp.Stats.Losses,
		WinRate:             pbResp.Stats.WinRate,
		SkillRank:           pbResp.Stats.SkillRank,
		RankChangeLastMatch: pbResp.Stats.RankChangeLastMatch,
	}, nil
}

func (c *playerClient) SearchPlayers(ctx context.Context, query string, limit, offset int) ([]gateway.PlayerProfile, error) {
	pbReq := &pb.SearchPlayersRequest{
		Query:  query,
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	pbResp, err := c.client.SearchPlayers(ctx, pbReq)
	if err != nil {
		return nil, err
	}

	players := make([]gateway.PlayerProfile, 0, len(pbResp.Players))
	for _, player := range pbResp.Players {
		players = append(players, gateway.PlayerProfile{
			PlayerID:     player.PlayerId,
			Name:         player.Name,
			Email:        player.Email,
			SkillRank:    player.SkillRank,
			TotalMatches: player.TotalMatches,
			Wins:         player.Wins,
			Losses:       player.Losses,
			CreatedAt:    player.CreatedAt,
			LastOnline:   player.LastOnline,
		})
	}

	return players, nil
}
