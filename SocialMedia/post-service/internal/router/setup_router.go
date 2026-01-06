package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rguarnizo/SocialMedia/pkg/auth"
	posthandler "github.com/rguarnizo/SocialMedia/post-service/internal/handler"
)

func SetupRouter(jwtSecrect string, postsHandler *posthandler.PostsHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Group(func(r chi.Router) {
		r.Use(auth.JWTMiddleware(jwtSecrect))

		r.Get("/posts", postsHandler.GetPosts)
		r.Post("/posts", postsHandler.AddPost)
		r.Post("/post/{id}/like", postsHandler.LikePost)
	})
	return r
}
