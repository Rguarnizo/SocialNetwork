package service

import (
	"github.com/google/uuid"
	"github.com/rguarnizo/SocialMedia/profile-service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetProfile(userID uuid.UUID) (*repository.UserProfile, error) {
	return s.repo.FindByID(userID)
}
