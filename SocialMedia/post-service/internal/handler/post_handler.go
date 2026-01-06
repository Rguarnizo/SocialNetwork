package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	ctx "github.com/rguarnizo/SocialMedia/pkg/auth"
	"github.com/rguarnizo/SocialMedia/post-service/internal/repository"
	"github.com/rguarnizo/SocialMedia/post-service/internal/service"
)

type PostsHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *PostsHandler {
	return &PostsHandler{service: service}
}

type AddPostRequest struct {
	Content string `json:"content"`
}

type AddPostResponse struct {
	Post repository.Post `json:"token"`
}

type PostResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"created_at"`
}

func (h *PostsHandler) GetPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := h.service.GetPosts()
	if err != nil {
		http.Error(w, "Error al obtener los posts"+err.Error(), http.StatusUnauthorized)
		return
	}

	resp := make([]PostResponse, 0, len(posts))
	for _, p := range posts {
		resp = append(resp, PostResponse{
			ID:        p.ID,
			UserID:    p.UserID,
			Content:   p.Content,
			CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *PostsHandler) AddPost(w http.ResponseWriter, r *http.Request) {

	user := ctx.UserFromRequest(r)

	var req AddPostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	post, err := h.service.CreatePost(user, req.Content)
	if err != nil {
		http.Error(w, "Usuario o clave incorrectos"+err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AddPostResponse{Post: post})
}

func (h *PostsHandler) LikePost(w http.ResponseWriter, r *http.Request) {

	user := ctx.UserFromRequest(r)
	postIDStr := chi.URLParam(r, "id")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		http.Error(w, "invalid post id", http.StatusBadRequest)
		return
	}

	err = h.service.LikePost(user, postID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Se ha dado like a la publicacion")
}
