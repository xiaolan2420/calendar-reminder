package service

import (
	"calendar-reminder/config"
	"calendar-reminder/dao"
	"calendar-reminder/model"
	"calendar-reminder/myWebsocket"
	"calendar-reminder/utils"
	"fmt"
	"github.com/robfig/cron/v3"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"time"
)

// 定义定时器
var reminderCron *cron.Cron

// Init 定时发送器
func Init() {
	reminderCron = cron.New(cron.WithSeconds())
	reminderCron.Start()

	// 添加新的定时任务，每 5 分钟检查一次数据库中距离提醒时间小于个5分钟的数据并设置提醒
	_, err := reminderCron.AddFunc("@every 5m", func() {
		reminders, err := dao.GetRemindersAfterTime(config.DB)
		if err != nil {
			fmt.Println("Error getting reminders:", err)
			return
		}
		for _, reminder := range reminders {
			fmt.Println("reminder:", reminder)

			scheduleTime := reminder.ReminderAt.Sub(time.Now())

			time.AfterFunc(scheduleTime, func() {
				err := SendReminder(config.DB, &reminder)
				if err != nil {
					fmt.Println("Error sending reminder in scheduled job:", err)
				}
			})

			if err != nil {
				fmt.Println("Error scheduling reminder:", err)
			}

		}
	})
	if err != nil {
		fmt.Println("Error adding new cron job:", err)
	}
}

// CreateReminder 增加提醒信息
func CreateReminder(db *gorm.DB, userID int, content string, reminderAt string, reminderMethod string) error {
	// 将字符串转换为时间格式
	reminderTime, err := time.Parse(time.RFC3339, reminderAt)
	if err != nil {
		fmt.Println("init time err")
		return err
	}

	reminder := model.Reminder{
		CreatorID:      userID,
		Content:        content,
		ReminderAt:     reminderTime,
		ReminderMethod: reminderMethod,
	}

	// 保存提醒消息到数据库
	err = dao.CreateReminder(db, &reminder)
	if err != nil {
		return err
	}

	return nil
}

// GetRemindersByUserID 获取用户的提醒信息列表
func GetRemindersByUserID(db *gorm.DB, userID int) ([]model.Reminder, error) {
	return dao.GetRemindersByUserID(db, userID)
}

// UpdateReminder 更新提醒信息
func UpdateReminder(db *gorm.DB, userID int, reminderID int, content string, reminderAt string) error {
	reminderTime, err := time.Parse(time.RFC3339, reminderAt)
	if err != nil {
		return err
	}

	// 根据用户id和消息id查找要更新的消息
	reminder, err := dao.GetReminderByIDAndUserID(db, reminderID, userID)
	if err != nil {
		return err
	}

	reminder.Content = content
	reminder.ReminderAt = reminderTime
	return dao.UpdateReminder(db, reminder)
}

// DeleteReminder 删除提醒信息
func DeleteReminder(db *gorm.DB, userID int, reminderID int) error {
	return dao.DeleteReminderByIDAndUserID(db, reminderID, userID)
}

// SendReminder 发送消息提醒
func SendReminder(db *gorm.DB, reminder *model.Reminder) error {
	// websocket提醒
	handler := myWebsocket.Handler{}
	err := handler.SendReminderToClient(reminder.CreatorID, reminder.Content)
	if err != nil {
		return err
	}
	// 根据提醒方式进行不同的处理
	if reminder.ReminderMethod == "sms" {
		// 调用短信发送函数
		return SendSmsReminder(db, reminder)
	} else if reminder.ReminderMethod == "email" {
		// 调用邮件发送函数
		return SendEmailReminder(db, reminder)
	}
	return nil
}

// SendSmsReminder 发送短信
func SendSmsReminder(db *gorm.DB, reminder *model.Reminder) error {
	// 获取用户手机号，可以从数据库根据用户ID查询或者在提醒结构体中添加手机号字段
	user, err := dao.GetUserByID(db, reminder.CreatorID)
	if err != nil {
		return err
	}
	phoneNumber := user.Phone
	// 发送短信
	return utils.SendSmsReminder(phoneNumber, reminder.Content)
}

func SendEmailReminder(db *gorm.DB, reminder *model.Reminder) error {
	// 获取用户的邮箱地址，可以从数据库根据用户ID查询或者在提醒结构体中添加邮箱字段
	user, err := dao.GetUserByID(db, reminder.CreatorID)
	if err != nil {
		return err
	}
	email := user.Email

	m := gomail.NewMessage()
	m.SetHeader("From", "2472559141@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "日历提醒")
	m.SetBody("text/plain", reminder.Content)

	d := gomail.NewDialer("smtp.qq.com", 587, "2472559141@qq.com", "rjaxtlmtalsydiga")
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
