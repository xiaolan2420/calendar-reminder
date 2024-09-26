package test

import (
	"calendar-reminder/model"
	"calendar-reminder/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

// 新建提醒测试
func TestCreateReminder(t *testing.T) {
	// 建立真实的数据库连接
	dsn := "root:1@tcp(localhost:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	type args struct {
		userID         int
		content        string
		reminderAt     string
		ReminderMethod string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "新建提醒成功",
			args: args{
				userID:         1,
				content:        "睡觉",
				reminderAt:     "2024-09-25T10:00:00+08:00",
				ReminderMethod: "sms",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 启动事务
			tx := db.Begin()
			if tx.Error != nil {
				t.Fatalf("启动事务失败: %v", tx.Error)
			}

			// 新增提醒
			err := service.CreateReminder(tx, tt.args.userID, tt.args.content, tt.args.reminderAt, tt.args.ReminderMethod)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

// 删除提醒测试
func TestDeleteReminder(t *testing.T) {
	// 建立真实的数据库连接
	dsn := "root:1@tcp(localhost:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	type args struct {
		userID     int
		reminderID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "删除提醒成功",
			args: args{
				userID:     1,
				reminderID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := db.Begin()
			if tx.Error != nil {
				t.Fatalf("启动事务失败: %v", tx.Error)
			}

			if err := service.DeleteReminder(tx, tt.args.userID, tt.args.reminderID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteReminder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 展示提醒测试
func TestGetRemindersByUserID(t *testing.T) {
	// 建立真实的数据库连接
	dsn := "root:1@tcp(localhost:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Reminder
		wantErr bool
	}{
		{
			name: "展示提醒成功",
			args: args{
				userID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := db.Begin()
			if tx.Error != nil {
				t.Fatalf("启动事务失败: %v", tx.Error)
			}

			_, err := service.GetRemindersByUserID(tx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRemindersByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

// 更新提醒测试
func TestUpdateReminder(t *testing.T) {
	// 建立真实的数据库连接
	dsn := "root:1@tcp(localhost:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	type args struct {
		userID     int
		reminderID int
		content    string
		reminderAt string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "更新提醒成功",
			args: args{
				userID:     1,
				reminderID: 1,
				content:    "睡觉",
				reminderAt: "2024-09-25T10:00:00+08:00",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 启动事务
			tx := db.Begin()
			if tx.Error != nil {
				t.Fatalf("启动事务失败: %v", tx.Error)
			}

			// 新增提醒
			err := service.UpdateReminder(tx, tt.args.userID, tt.args.reminderID, tt.args.content, tt.args.reminderAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestSendSmsReminder(t *testing.T) {
	// 建立真实的数据库连接
	dsn := "root:1@tcp(localhost:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	type args struct {
		reminder *model.Reminder
	}
	reminders := &model.Reminder{
		ID:             1,
		CreatorID:      1,
		Content:        "睡觉睡觉",
		ReminderAt:     time.Now(),
		ReminderMethod: "sms",
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid reminder",
			args: args{
				reminder: reminders,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 启动事务
			tx := db.Begin()
			if tx.Error != nil {
				t.Fatalf("启动事务失败: %v", tx.Error)
			}

			if err := service.SendSmsReminder(tx, tt.args.reminder); (err != nil) != tt.wantErr {
				t.Errorf("SendSmsReminder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
