package api

import (
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

// Create player
// @Accept      json
// @Produce     json
// @Router	/player
func (a *Api) CreatePlayer(w http.ResponseWriter, r *http.Request) error {
	var payload gateway.CreatePlayerRequest
	if err := jsonutil.Parse(r, &payload); err != nil {
		return errs.InvalidJson(err)
	}

	rsp, err := a.player.CreatePlayer(r.Context(), &payload)
	if err != nil {
		return errs.InternalServerError(err)
	}

	return jsonutil.Write(w, http.StatusOK, gateway.ApiResponse{
		Data: map[string]string{
			"id": rsp.PlayerID,
		},
	})
}
