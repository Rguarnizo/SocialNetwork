package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindByID(id uuid.UUID) (*UserProfile, error) {
	var user UserProfile
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}
