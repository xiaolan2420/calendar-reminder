package service

import (
	"calendar-reminder/dao"
	"calendar-reminder/model"
	"calendar-reminder/utils"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RegisterUser 注册用户
func RegisterUser(db *gorm.DB, phone string, password string, email string, code string) error {
	// 校验验证码
	if !utils.VerifySmsCode(phone, code) {
		return errors.New("code error")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 创建用户记录
	user := model.User{
		Phone:    phone,
		Password: string(hashedPassword),
		Email:    email,
	}
	return dao.CreateUser(db, &user)
}

// LoginUser 使用密码登录
func LoginUser(db *gorm.DB, phone string, password string) (string, error) {
	user, err := dao.GetUserByPhone(db, phone)
	if err != nil {
		return "not user", err
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {

		return "password error", err
	}

	// 生成 JWT Token
	jwtService := &utils.JWT{}
	token, err := jwtService.GenerateToken(user.ID)
	if err != nil {
		return "GenerateToken error", err
	}

	return token, nil
}

// LoginUserByCode 使用验证码登录
func LoginUserByCode(db *gorm.DB, phone string, code string) (string, error) {
	user, err := dao.GetUserByPhone(db, phone)
	if err != nil {
		return "not user", err
	}

	// 校验验证码
	if !utils.VerifySmsCode(phone, code) {
		return "SmsCode error", errors.New("code error")
	}

	// 生成 JWT Token
	jwtService := &utils.JWT{}
	token, err := jwtService.GenerateToken(user.ID)
	if err != nil {
		return "GenerateToken error", err
	}

	return token, nil
}

// GetCode 生成发送验证码
func GetCode(phone string) (bool, error) {

	// 生成验证码
	code := utils.GenerateVerificationCode()
	fmt.Println("code", code)

	// 保存验证码
	if err := utils.SaveVerificationCode(phone, code); err != nil {
		return false, err
	}

	// 发送验证码
	if err := utils.SendSmsCode(phone, code); err != nil {
		return false, err
	}
	return true, nil
}
