package gateway

// ==================== HTTP Request DTOs ====================
type CreatePlayerRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=50"`
	Email string `json:"email" validate:"required,email"`
}

type UpdatePlayerRequest struct {
	PlayerID  string  `json:"player_id" validate:"required"`
	Name      *string `json:"name,omitempty" validate:"omitempty,min=3,max=50"`
	SkillRank *int32  `json:"skill_rank,omitempty" validate:"omitempty,min=0"`
}

type JoinQueueRequest struct {
	PlayerID string `json:"player_id" validate:"required"`
	GameMode string `json:"game_mode" validate:"required"`
	Region   string `json:"region" validate:"required"`
}

type CancelMatchRequest struct {
	PlayerID string `json:"player_id" validate:"required"`
	MatchID  string `json:"match_id" validate:"required"`
}

type CreateSessionRequest struct {
	MatchID   string   `json:"match_id" validate:"required"`
	PlayerIDs []string `json:"player_ids" validate:"required,min=1"`
	GameMode  string   `json:"game_mode" validate:"required"`
	Region    string   `json:"region" validate:"required"`
}

type EndSessionRequest struct {
	SessionID string                `json:"session_id" validate:"required"`
	Results   []PlayerSessionResult `json:"results" validate:"required"`
}

type ValidateTokenRequest struct {
	SessionToken string `json:"session_token" validate:"required"`
	PlayerID     string `json:"player_id" validate:"required"`
}

// ==================== Common HTTP Response ====================

type ApiResponse struct {
	Data interface{} `json:"data,omitempty"`
}
