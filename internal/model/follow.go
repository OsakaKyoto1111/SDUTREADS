package model

type Follow struct {
	ID          int64 `json:"id" gorm:"primaryKey"`
	FollowerID  int64 `json:"follower_id" gorm:"index;not null"`
	FollowingID int64 `json:"following_id" gorm:"index;not null"`
}
