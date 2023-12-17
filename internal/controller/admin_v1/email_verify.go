package admin_v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/smtp"
	"time"
	"youdangzhe/internal/controller"
	log "youdangzhe/internal/pkg/logger"
	"youdangzhe/internal/validator"
	"youdangzhe/internal/validator/form"
)

type EmailVerificationController struct {
	controller.Api
}

func NewEmailVerificationController() *EmailVerificationController {
	return &EmailVerificationController{}
}

func (api *EmailVerificationController) SendCode(c *gin.Context) {
	emailForm := form.NewEmailCodeForm()
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckQueryParams(c, &emailForm); err != nil {
		api.Fail(c, 404, "Email format is not correct.")
		return
	}
	result := sendVerifyCode(emailForm.Email)
	if !result {
		api.Fail(c, 404, "send verify code to email fail")
		return
	}

	api.Success(c, "send verify code success")
	return
}

func sendVerifyCode(recipientEmail string) bool {
	// 配置邮件服务器信息
	smtpServer := "smtp.163.com"
	smtpPort := 25
	senderEmail := "cyzabc66@163.com"
	password := "GULYZNDKQQKMFUZR"

	verifyCode := generateInviteCode()
	log.Logger.Info("the verify code is " + verifyCode)
	subject := "验证码"
	body := "您的验证码是：" + verifyCode
	// 邮件内容
	message := fmt.Sprintf("To: %s\r\n", recipientEmail) +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body

	// 配置SMTP客户端
	auth := smtp.PlainAuth("", senderEmail, password, smtpServer)

	// 连接到SMTP服务器
	err := smtp.SendMail(smtpServer+":"+fmt.Sprintf("%d", smtpPort), auth, senderEmail, []string{recipientEmail}, []byte(message))
	if err != nil {
		log.Logger.Error("邮件发送失败: " + err.Error())
		return false
	}

	log.Logger.Info("邀请码已发送到邮箱: " + recipientEmail)
	return true

}

func generateInviteCode() string {
	// 种子值
	rand.NewSource(time.Now().UnixNano())

	// 生成六位数的随机邀请码
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ" // 可用字符集合
	codeLength := 6                                        // 邀请码长度
	code := make([]byte, codeLength)
	for i := 0; i < codeLength; i++ {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
