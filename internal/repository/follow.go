package repository

import (
	"myapp/internal/model"

	"gorm.io/gorm"
)

type FollowRepository interface {
	Follow(followerID, followingID int64) error
	Unfollow(followerID, followingID int64) error
	GetFollowingIDs(followerID int64) ([]int64, error)
	IsFollowing(followerID, followingID int64) (bool, error)
}

type followRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) FollowRepository {
	return &followRepository{db: db}
}

func (r *followRepository) Follow(followerID, followingID int64) error {
	f := model.Follow{FollowerID: followerID, FollowingID: followerID}
	return r.db.Create(&f).Error
}
func (r *followRepository) Unfollow(followerID, followingID int64) error {
	return r.db.Where("follower_id = ? AND following_id = ?", followerID, followingID).
		Delete(&model.Follow{}).Error
}
func (r *followRepository) GetFollowingIDs(followerID int64) ([]int64, error) {
	var ids []int64
	if err := r.db.Model(&model.Follow{}).
		Where("follower_id = ?", followerID).
		Pluck("following_id", &ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}
func (r *followRepository) IsFollowing(followerID, followingID int64) (bool, error) {
	var count int64
	if err := r.db.Model(&model.Follow{}).
		Where("follower_id = ? AND following_id = ?", followerID, followingID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
