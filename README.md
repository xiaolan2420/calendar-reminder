# 日历提醒服务

## 项目介绍
实现了注册、登录、日程的增删改查、短信通知、邮箱通知和websocket通知
###登录和注册 
使用了阿里云短信服务实现验证码发放，注册需要接收验证码，有密码登录和验证码登录两种方式。
验证码存放在redis中，设置了10分钟有效期
使用了JWT鉴权，只有该用户才可以管理本人的日程
###日程的提醒
实现了通过邮箱、短信和websocket(还有些缺陷)在提醒时间发送日程提醒
## 项目不足
未能完整完成单元测试部分，接口测试用多了，这部分有些欠缺，还得继续学习


