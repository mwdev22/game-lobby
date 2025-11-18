package gateway

import "context"

// ==================== Matchmaking Client ====================

type MatchmakingClient interface {
	JoinQueue(ctx context.Context, req *JoinQueueRequest) (*JoinQueueResponse, error)
	LeaveQueue(ctx context.Context, req *LeaveQueueRequest) (*LeaveQueueResponse, error)
	GetQueueStatus(ctx context.Context, playerID string) (*QueueStatusResponse, error)
	CancelMatch(ctx context.Context, playerID, matchID string) (*CancelMatchResponse, error)
	GetMatchHistory(ctx context.Context, playerID string, limit, offset int) (*MatchHistoryResponse, error)
}

type JoinQueueRequest struct {
	PlayerID string `json:"player_id"`
	GameMode string `json:"game_mode"`
	Region   string `json:"region"`
}

type JoinQueueResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	QueuePosition int32  `json:"queue_position"`
}

type LeaveQueueRequest struct {
	PlayerID string `json:"player_id"`
}

type LeaveQueueResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type QueueStatusResponse struct {
	PlayerID             string `json:"player_id"`
	Position             int32  `json:"position"`
	EstimatedWaitSeconds int32  `json:"estimated_wait_seconds"`
	Status               string `json:"status"`
	InQueue              bool   `json:"in_queue"`
}

type CancelMatchResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type MatchHistoryResponse struct {
	Matches    []MatchInfo `json:"matches"`
	TotalCount int32       `json:"total_count"`
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
	CreatePlayer(ctx context.Context, req *CreatePlayerRequest) (*CreatePlayerResponse, error)
	GetPlayer(ctx context.Context, playerID string) (*PlayerProfileResponse, error)
	UpdatePlayer(ctx context.Context, req *UpdatePlayerRequest) (*UpdatePlayerResponse, error)
	GetPlayerStats(ctx context.Context, playerID string) (*PlayerStatsResponse, error)
	SearchPlayers(ctx context.Context, query string, limit, offset int) (*SearchPlayersResponse, error)
}

type CreatePlayerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreatePlayerResponse struct {
	PlayerID  string `json:"player_id"`
	Name      string `json:"name"`
	SkillRank int32  `json:"skill_rank"`
	CreatedAt int64  `json:"created_at"`
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

type PlayerProfileResponse struct {
	Player PlayerProfile `json:"player"`
}

type UpdatePlayerRequest struct {
	PlayerID  string  `json:"player_id"`
	Name      *string `json:"name,omitempty"`
	SkillRank *int32  `json:"skill_rank,omitempty"`
}

type UpdatePlayerResponse struct {
	Success bool          `json:"success"`
	Player  PlayerProfile `json:"player"`
}

type PlayerStats struct {
	TotalMatches        int32   `json:"total_matches"`
	Wins                int32   `json:"wins"`
	Losses              int32   `json:"losses"`
	WinRate             float64 `json:"win_rate"`
	SkillRank           int32   `json:"skill_rank"`
	RankChangeLastMatch int32   `json:"rank_change_last_match"`
}

type PlayerStatsResponse struct {
	Stats PlayerStats `json:"stats"`
}

type SearchPlayersResponse struct {
	Players    []PlayerProfile `json:"players"`
	TotalCount int32           `json:"total_count"`
}

// ==================== Session Client ====================

type SessionClient interface {
	CreateSession(ctx context.Context, req *CreateSessionRequest) (*CreateSessionResponse, error)
	GetSession(ctx context.Context, sessionID string) (*SessionInfoResponse, error)
	EndSession(ctx context.Context, req *EndSessionRequest) (*EndSessionResponse, error)
	ValidateSessionToken(ctx context.Context, token, playerID string) (*ValidateTokenResponse, error)
	GetActiveSessions(ctx context.Context, playerID string) (*ActiveSessionsResponse, error)
}

type CreateSessionRequest struct {
	MatchID   string   `json:"match_id"`
	PlayerIDs []string `json:"player_ids"`
	GameMode  string   `json:"game_mode"`
	Region    string   `json:"region"`
}

type CreateSessionResponse struct {
	SessionID     string `json:"session_id"`
	SessionToken  string `json:"session_token"`
	ServerAddress string `json:"server_address"`
	ServerPort    int32  `json:"server_port"`
	ExpiresAt     int64  `json:"expires_at"`
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

type SessionInfoResponse struct {
	Session SessionInfo `json:"session"`
}

type EndSessionRequest struct {
	SessionID string                `json:"session_id"`
	Results   []PlayerSessionResult `json:"results"`
}

type PlayerSessionResult struct {
	PlayerID string `json:"player_id"`
	Score    int32  `json:"score"`
	Winner   bool   `json:"winner"`
}

type EndSessionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ValidateTokenResponse struct {
	Valid     bool         `json:"valid"`
	SessionID string       `json:"session_id"`
	Session   *SessionInfo `json:"session,omitempty"`
}

type ActiveSessionsResponse struct {
	Sessions []SessionInfo `json:"sessions"`
}

