package test

import (
	"calendar-reminder/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

// 验证码发送,不保存验证码，测试成功
func TestGetCode(t *testing.T) {
	type args struct {
		phone string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "用户验证码发送测试",
			args: args{
				phone: "18476563988",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetCode(tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// 注册测试，不用redis通过
func TestRegisterUser(t *testing.T) {
	// 建立数据库连接
	dsn := "root:1@tcp(localhost:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	type args struct {
		phone    string
		password string
		email    string
		code     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "成功注册用户",
			args: args{
				phone:    "12345678901",
				password: "123456",
				email:    "123@qq.com",
				code:     "1234",
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

			// 测试注册用户
			err := service.RegisterUser(tx, tt.args.phone, tt.args.password, tt.args.email, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

// 密码登录测试
func TestLoginUser(t *testing.T) {
	// 建立数据库连接
	dsn := "root:1@tcp(localhost:3306)/reminder?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	type args struct {
		phone    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "用密码成功登录",
			args: args{
				phone:    "18476563988",
				password: "123456",
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

			// 测试注册用户
			_, err := service.LoginUser(tx, tt.args.phone, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}

		})

	}
}
