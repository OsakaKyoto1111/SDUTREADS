package repository

import (
	"myapp/internal/model"

	"gorm.io/gorm"
)

type LikeRepository interface {
	Add(userID, postID int64) error
	Remove(userID, postID int64) error
	Exists(userID, postID int64) (bool, error)
	Count(postID int64) (int64, error)
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeRepository{db: db}
}

func (r *likeRepository) Add(userID, postID int64) error {
	like := model.Like{UserID: userID, PostID: postID}
	return r.db.Create(&like).Error
}

func (r *likeRepository) Remove(userID, postID int64) error {
	return r.db.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&model.Like{}).Error
}

func (r *likeRepository) Exists(userID, postID int64) (bool, error) {
	var count int64
	if err := r.db.Model(&model.Like{}).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *likeRepository) Count(postID int64) (int64, error) {
	var count int64
	if err := r.db.Model(&model.Like{}).
		Where("post_id = ?", postID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
