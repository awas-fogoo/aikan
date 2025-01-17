package server

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"net/smtp"
)

// SendVerificationCode server发送验证码
func SendVerificationCode(em string, code string) bool {
	username := viper.GetString("email.username")
	password := viper.GetString("email.password")
	e := email.NewEmail()
	e.From = "lee <yassus@foxmail.com>" // 发件人 这里的收件人邮件要和下面的一样才可以
	//e.To = []string{"211126958@qq.com", "2259768992@qq.com", "1991469616@qq.com", "94633380@qq.com"}
	{
		/*
		   CC全称是Carbon Copy，意为抄送，BCC全称Blind Carbon Copy，意为暗抄送，收件人看不到被暗抄送给了谁。
		*/
		e.Bcc = []string{em} // 密送 (Blind carbon copy)
		//e.Cc = []string{"211126958@qq.com"}  // 副本
	}

	e.Subject = "布比卡因《 BUP.PUB 》" // 标题

	//e.Text = []byte("test content bcc") // 文本文件内容，二选一
	e.HTML = []byte(fmt.Sprintf("<h1>您的验证码：%s</h1><br><h6>仅在三分钟之内有效哦～</h6>", code)) // html内容

	err := e.Send("smtp.qq.com:587",
		smtp.PlainAuth("", username, password, "smtp.qq.com"))
	if err != nil {
		return false
	}
	//_, err = e.AttachFile("test.txt")
	//if err != nil {
	// return
	//} // 附加文件
	return true
}
