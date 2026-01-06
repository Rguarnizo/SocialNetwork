package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post Post) (Post, error)
	List() ([]Post, error)
	Like(userID, postID uuid.UUID) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post Post) (Post, error) {

	if err := r.db.Create(&post).Error; err != nil {
		return Post{}, err
	}
	return post, nil
}

func (r *postRepository) List() ([]Post, error) {
	var posts []Post

	if err := r.db.
		Order("created_at DESC").
		Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *postRepository) Like(userID, postID uuid.UUID) error {
	like := PostLike{
		PostID: postID,
		UserID: userID,
	}

	return r.db.Create(&like).Error
}
