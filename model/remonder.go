package model

import "time"

// Reminder 提醒模型
type Reminder struct {
	ID             int       `gorm:"primaryKey"`
	CreatorID      int       `gorm:"not null"`
	Content        string    `gorm:"type:text;not null"`
	ReminderAt     time.Time `gorm:"not null"`
	ReminderMethod string    `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
