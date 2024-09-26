package dao

import (
	"calendar-reminder/model"
	"gorm.io/gorm"
	"time"
)

// CreateReminder 创建提醒信息
func CreateReminder(db *gorm.DB, reminder *model.Reminder) error {
	return db.Create(reminder).Error
}

// GetRemindersByUserID 根据用户 ID 获取提醒信息列表
func GetRemindersByUserID(db *gorm.DB, userID int) ([]model.Reminder, error) {
	var reminders []model.Reminder
	result := db.Where("creator_id = ?", userID).Find(&reminders)
	return reminders, result.Error
}

// GetReminderByIDAndUserID 根据提醒 ID 和用户 ID 获取提醒信息
func GetReminderByIDAndUserID(db *gorm.DB, reminderID int, userID int) (*model.Reminder, error) {
	var reminder model.Reminder
	result := db.Where("id = ? AND creator_id = ?", reminderID, userID).First(&reminder)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reminder, nil
}

// UpdateReminder 更新提醒信息
func UpdateReminder(db *gorm.DB, reminder *model.Reminder) error {
	return db.Save(reminder).Error
}

// DeleteReminderByIDAndUserID 根据提醒 ID 和用户 ID 删除提醒信息
func DeleteReminderByIDAndUserID(db *gorm.DB, reminderID int, userID int) error {
	return db.Where("id = ? AND creator_id = ?", reminderID, userID).Delete(&model.Reminder{}).Error
}

// GetRemindersAfterTime 找出还剩5分钟的
func GetRemindersAfterTime(db *gorm.DB) ([]model.Reminder, error) {
	now := time.Now()
	HourLater := now.Add(5 * time.Minute)
	var reminders []model.Reminder
	err := db.Where("reminder_at >? AND reminder_at <=?", now, HourLater).Find(&reminders).Error
	return reminders, err
}
