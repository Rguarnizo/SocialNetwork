package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/rguarnizo/SocialMedia/pkg/auth"
	"github.com/rguarnizo/SocialMedia/post-service/internal/repository"
)

type PostService interface {
	CreatePost(user auth.User, content string) (repository.Post, error)
	GetPosts() ([]repository.Post, error)
	LikePost(user auth.User, post uuid.UUID) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(user auth.User, content string) (repository.Post, error) {
	content = strings.TrimSpace(content)

	if content == "" {
		return repository.Post{}, errors.New("content cannot be empty")
	}

	post := repository.Post{
		ID:      uuid.New(),
		UserID:  user.ID,
		Content: content,
	}

	return s.repo.Create(post)
}

func (s *postService) GetPosts() ([]repository.Post, error) {
	posts, err := s.repo.List()
	if err != nil {
		return nil, errors.New("content cannot be empty")
	}

	return posts, nil
}

func (s *postService) LikePost(user auth.User, post uuid.UUID) error {
	err := s.repo.Like(user.ID, post)
	if err != nil {
		return errors.New("content cannot be empty")
	}

	return nil
}
