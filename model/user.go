package model

import "time"

// User 用户模型
type User struct {
	ID        int    `gorm:"primaryKey"`
	Phone     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
