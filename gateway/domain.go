package gateway

import (
	"context"
	"errors"
	"fmt"
)

// ==================== Matchmaking Client ====================

type MatchmakingClient interface {
	JoinQueue(ctx context.Context, req *JoinQueueRequest) (*QueueJoinResult, error)
	LeaveQueue(ctx context.Context, playerID string) (bool, error)
	GetQueueStatus(ctx context.Context, playerID string) (*QueueStatus, error)
	CancelMatch(ctx context.Context, playerID, matchID string) (bool, error)
	GetMatchHistory(ctx context.Context, playerID string, limit, offset int) ([]MatchInfo, error)
	GetPlayersByMatchID(ctx context.Context, matchIDs []string) ([]PlayerProfile, error)
	GetMatch(ctx context.Context, id string) (*MatchInfo, error)
}

type QueueJoinResult struct {
	Success       bool
	Message       string
	QueuePosition int32
}

type QueueStatus struct {
	PlayerID             string
	Position             int32
	EstimatedWaitSeconds int32
	Status               string
	InQueue              bool
}

type MatchInfo struct {
	MatchID      string       `json:"match_id"`
	Players      []PlayerInfo `json:"players"`
	GameMode     string       `json:"game_mode"`
	Region       string       `json:"region"`
	SessionToken string       `json:"session_token"`
	CreatedAt    int64        `json:"created_at"`
}

type PlayerInfo struct {
	PlayerID  string `json:"player_id"`
	Name      string `json:"name"`
	SkillRank int32  `json:"skill_rank"`
}

// ==================== Player Client ====================

type PlayerClient interface {
	CreatePlayer(ctx context.Context, req *CreatePlayerRequest) (string, error)

	GetPlayer(ctx context.Context, playerID string) (*PlayerProfile, error)
	UpdatePlayer(ctx context.Context, req *UpdatePlayerRequest) error
	GetPlayerStats(ctx context.Context, playerID string) (*PlayerStats, error)
	SearchPlayers(ctx context.Context, query string, limit, offset int) ([]PlayerProfile, error)
}

type PlayerProfile struct {
	PlayerID     string `json:"player_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	SkillRank    int32  `json:"skill_rank"`
	TotalMatches int32  `json:"total_matches"`
	Wins         int32  `json:"wins"`
	Losses       int32  `json:"losses"`
	CreatedAt    int64  `json:"created_at"`
	LastOnline   int64  `json:"last_online"`
}

type PlayerStats struct {
	TotalMatches        int32   `json:"total_matches"`
	Wins                int32   `json:"wins"`
	Losses              int32   `json:"losses"`
	WinRate             float64 `json:"win_rate"`
	SkillRank           int32   `json:"skill_rank"`
	RankChangeLastMatch int32   `json:"rank_change_last_match"`
}

// ==================== Session Client ====================

type SessionClient interface {
	CreateSession(ctx context.Context, req *CreateSessionRequest) (string, error)
	GetSession(ctx context.Context, sessionID string) (*SessionInfo, error)
	EndSession(ctx context.Context, sessionID string, results []PlayerSessionResult) error
	ValidateSessionToken(ctx context.Context, token, playerID string) (*SessionInfo, bool, error)
	GetActiveSessions(ctx context.Context, playerID string) ([]SessionInfo, error)
}

type SessionInfo struct {
	SessionID     string   `json:"session_id"`
	MatchID       string   `json:"match_id"`
	PlayerIDs     []string `json:"player_ids"`
	GameMode      string   `json:"game_mode"`
	Region        string   `json:"region"`
	Status        string   `json:"status"`
	ServerAddress string   `json:"server_address"`
	ServerPort    int32    `json:"server_port"`
	CreatedAt     int64    `json:"created_at"`
	ExpiresAt     int64    `json:"expires_at"`
}

type PlayerSessionResult struct {
	PlayerID string `json:"player_id"`
	Score    int32  `json:"score"`
	Winner   bool   `json:"winner"`
}

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrInternal     = func(err error) error {
		return fmt.Errorf("internal server error: %w", err)
	}
	ErrNotFound = func(name string) error {
		return fmt.Errorf("%s not found", name)
	}
)
