package dao

import (
	"calendar-reminder/model"
	"gorm.io/gorm"
)

// CreateUser 创建用户
func CreateUser(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}

// GetUserByPhone 根据手机号查找用户
func GetUserByPhone(db *gorm.DB, phone string) (*model.User, error) {
	var user model.User
	result := db.Where("phone = ?", phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByID 根据用户 ID 查找用户
func GetUserByID(db *gorm.DB, userID int) (*model.User, error) {
	var user model.User
	result := db.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
