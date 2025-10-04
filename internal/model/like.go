package model

type Like struct {
	ID     int64 `json:"id" gorm:"primaryKey"`
	UserID int64 `json:"user_id" gorm:"index;not null"`
	PostID int64 `json:"post_id" gorm:"index;not null"`
}
