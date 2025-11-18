package gateway

// ==================== REST API Request/Response DTOs ====================

// Matchmaking DTOs
type JoinQueueHTTPRequest struct {
	PlayerID string `json:"player_id" validate:"required"`
	GameMode string `json:"game_mode" validate:"required"`
	Region   string `json:"region" validate:"required"`
}

type LeaveQueueHTTPRequest struct {
	PlayerID string `json:"player_id" validate:"required"`
}

type CancelMatchHTTPRequest struct {
	PlayerID string `json:"player_id" validate:"required"`
	MatchID  string `json:"match_id" validate:"required"`
}

// Player DTOs
type CreatePlayerHTTPRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=50"`
	Email string `json:"email" validate:"required,email"`
}

type UpdatePlayerHTTPRequest struct {
	Name      *string `json:"name,omitempty" validate:"omitempty,min=3,max=50"`
	SkillRank *int32  `json:"skill_rank,omitempty" validate:"omitempty,min=0"`
}

// Session DTOs
type CreateSessionHTTPRequest struct {
	MatchID   string   `json:"match_id" validate:"required"`
	PlayerIDs []string `json:"player_ids" validate:"required,min=1"`
	GameMode  string   `json:"game_mode" validate:"required"`
	Region    string   `json:"region" validate:"required"`
}

type EndSessionHTTPRequest struct {
	SessionID string                `json:"session_id" validate:"required"`
	Results   []PlayerSessionResult `json:"results" validate:"required"`
}

type ValidateTokenHTTPRequest struct {
	SessionToken string `json:"session_token" validate:"required"`
	PlayerID     string `json:"player_id" validate:"required"`
}

// Common response wrapper
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

