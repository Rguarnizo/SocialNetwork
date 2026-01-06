package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfile struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	FirstName string         `json:"first_name" gorm:"column:first_name;not null"`
	LastName  *string        `json:"last_name,omitempty" gorm:"column:last_name"`
	BornDate  *time.Time     `json:"born_date,omitempty" gorm:"column:born_date"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index"`
}

func (UserProfile) TableName() string {
	return "auth.users"
}
