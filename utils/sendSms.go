package utils

import (
	"calendar-reminder/config"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

// RedisContext
var ctx = context.Background()

// CreateClient 创建阿里云短信客户端
func CreateClient() (*dysmsapi20170525.Client, error) {
	accessKeyId := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")

	if accessKeyId == "" || accessKeySecret == "" {
		return nil, fmt.Errorf("AccessKeyId or AccessKeySecret is not set or is empty")
	}

	aLiYunConfig := &client.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}

	aLiYunClient, err := dysmsapi20170525.NewClient(aLiYunConfig)
	if err != nil {
		return nil, err
	}

	return aLiYunClient, nil
}

// GenerateVerificationCode 生成六位数字验证码
func GenerateVerificationCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := r.Intn(900000) + 100000
	return strconv.Itoa(code)
}

// SendSmsCode 发送短信验证码
func SendSmsCode(phoneNumber string, code string) error {
	aLiYunClient, err := CreateClient()
	if err != nil {
		return err
	}

	request := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phoneNumber),
		SignName:      tea.String("皛懒日程提醒"),
		TemplateCode:  tea.String("SMS_304946879"),
		TemplateParam: tea.String(fmt.Sprintf("{\"code\":\"%s\"}", code)),
	}

	_, err = aLiYunClient.SendSms(request)
	if err != nil {
		return err
	}

	return nil
}

// SaveVerificationCode 保存验证码到 Redis
func SaveVerificationCode(phoneNumber string, code string) error {
	err := config.RDB.Set(ctx, phoneNumber, code, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

// VerifySmsCode 验证验证码
func VerifySmsCode(phoneNumber string, code string) bool {
	storedCode, err := config.RDB.Get(ctx, phoneNumber).Result()
	if err != nil {
		return false
	}
	return storedCode == code
}

// SendSmsReminder 发送短信提醒消息
func SendSmsReminder(phoneNumber string, content string) error {
	aLiYunClient, err := CreateClient()
	if err != nil {
		return err
	}

	request := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phoneNumber),
		SignName:      tea.String("皛懒日程提醒"),
		TemplateCode:  tea.String("SMS_473855188"),
		TemplateParam: tea.String(fmt.Sprintf("{\"text\":\"%s\"}", content)),
	}

	_, err = aLiYunClient.SendSms(request)
	if err != nil {
		return err
	}

	return nil
}
