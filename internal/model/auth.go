package model

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm "uniqueIndex"`
	Password  string    `json "-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
