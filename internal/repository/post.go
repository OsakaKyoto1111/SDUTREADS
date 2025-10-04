package repository

import (
	"myapp/internal/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
	GetByID(id int64) (*model.Post, error)
	GetAll() ([]model.Post, error)
	Delete(id int64) error
	GetByUserIDs(ids []int64) ([]model.Post, error) // ✅ добавили
}

type postRepository struct{ db *gorm.DB }

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

// ➕ Создать пост
func (r *postRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

// 🔍 Получить пост по ID
func (r *postRepository) GetByID(id int64) (*model.Post, error) {
	var p model.Post
	if err := r.db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

// 🗑 Удалить пост
func (r *postRepository) Delete(id int64) error {
	return r.db.Delete(&model.Post{}, id).Error
}

// 📋 Получить все посты (новые сверху)
func (r *postRepository) GetAll() ([]model.Post, error) {
	var posts []model.Post
	if err := r.db.Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// 📄 Получить посты от конкретных пользователей
func (r *postRepository) GetByUserIDs(ids []int64) ([]model.Post, error) {
	var posts []model.Post
	if err := r.db.Where("user_id IN ?", ids).
		Order("created_at DESC").
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
