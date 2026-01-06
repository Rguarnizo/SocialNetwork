package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rguarnizo/SocialMedia/pkg/auth"
	"github.com/rguarnizo/SocialMedia/profile-service/internal/handler"
)

func SetupRouter(jwtSecrect string, profileHandler *handler.ProfileHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Group(func(r chi.Router) {
		r.Use(auth.JWTMiddleware(jwtSecrect))

		r.Get("/profile/me", profileHandler.Me)
	})
	return r
}
