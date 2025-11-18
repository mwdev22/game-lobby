package api

import (
	"errors"
	"gateway"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mwdev22/rest/jsonutil"
	"github.com/mwdev22/rest/middleware"
	"github.com/mwdev22/rest/utils/errs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *Api) Mount(r chi.Router) {
	// Middleware must be defined before any routes
	r.Use(middleware.Logger, middleware.RealIP)

	r.Get("/ping", a.Ping)

	r.Handle("/swagger/", httpSwagger.WrapHandler)

	fs := http.StripPrefix("/media/", http.FileServer(http.Dir("./media")))
	r.Handle("/media/*", fs)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/player", func(r chi.Router) {
			r.Post("/", middleware.Wrap(a.CreatePlayer))
		})
	})
}

// ping example
// @Summary     Ping the server
// @Description Simple endpoint to check if the server is running
// @Accept      json
// @Produce     json
// @Router      /ping [get]
func (a *Api) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"pong"}`))
}

// ----------- Player Routes -----------

func (a *Api) GetPlayer(w http.ResponseWriter, r *http.Request) error {
	playerID := chi.URLParam(r, "id")
	profile, err := a.service.GetPlayer(r.Context(), playerID)
	if err != nil {
		return errs.InternalServerError(err)
	}

	return jsonutil.Write(w, http.StatusOK, gateway.ApiResponse{
		Data: profile,
	})
}

// Create player
// @Accept      json
// @Produce     json
// @Router	/player
func (a *Api) CreatePlayer(w http.ResponseWriter, r *http.Request) error {
	var payload gateway.CreatePlayerRequest
	if err := jsonutil.Parse(r, &payload); err != nil {
		return errs.InvalidJson(err)
	}

	if payload.Name == "" || payload.Email == "" {
		return errs.InvalidJson(errors.New("empty name or email"))
	}

	id, err := a.service.CreatePlayer(r.Context(), &payload)
	if err != nil {
		return errs.InternalServerError(err)
	}

	return jsonutil.Write(w, http.StatusOK, gateway.ApiResponse{
		Data: map[string]string{
			"id": id,
		},
	})
}

func (a *Api) UpdatePlayer(w http.ResponseWriter, r *http.Request) error {
	var payload gateway.UpdatePlayerRequest
	if err := jsonutil.Parse(r, &payload); err != nil {
		return errs.InvalidJson(err)
	}

	err := a.service.UpdatePlayer(r.Context(), &payload)
	if err != nil {
		return errs.InternalServerError(err)
	}

	return jsonutil.Write(w, http.StatusOK, gateway.ApiResponse{
		Data: map[string]string{
			"status": "updated",
		},
	})
}

func (a *Api) GetPlayerStats(w http.ResponseWriter, r *http.Request) error {
	playerID := chi.URLParam(r, "id")
	stats, err := a.service.GetPlayerStats(r.Context(), playerID)
	if err != nil {
		return errs.InternalServerError(err)
	}

	return jsonutil.Write(w, http.StatusOK, gateway.ApiResponse{
		Data: stats,
	})
}

func (a *Api) SearchPlayers(w http.ResponseWriter, r *http.Request) error {
	// query := r.URL.Query().Get("query")
	// limit, offset := middleware.GetLimitOffset(r)

	// players, err := a.service.SearchPlayers(r.Context(), query, limit, offset)
	// if err != nil {
	// 	return errs.InternalServerError(err)
	// }

	// return jsonutil.Write(w, http.StatusOK, gateway.ApiResponse{
	// 	Data: players,
	// })
	return nil
}

// ----------- Matchmaking Routes -----------

func (a *Api) JoinQueue(w http.ResponseWriter, r *http.Request) error {
	var payload gateway.JoinQueueRequest
	if err := jsonutil.Parse(r, &payload); err != nil {
		return errs.InvalidJson(err)
	}

	return nil
}

func (a *Api) LeaveQueue(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *Api) GetQueueStatus(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *Api) GetMatch(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *Api) CancelMatch(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *Api) GetMatchHistory(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *Api) GetPlayersByMatchID(w http.ResponseWriter, r *http.Request) error {
	return nil
}
