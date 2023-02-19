package server

import (
	"awesomeProject0511/common"
	"awesomeProject0511/util"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

// SendVerificationCode server发送验证码
func SendVerificationCode(em string, code string) bool {

	e := email.NewEmail()
	e.From = "lee" // 发件人 这里的收件人邮件要和下面的一样才可以
	{
		/*
			CC全称是Carbon Copy，意为抄送，BCC全称Blind Carbon Copy，意为暗抄送，收件人看不到被暗抄送给了谁。
		*/
		e.Bcc = []string{em} // 密送 (Blind carbon copy)
	}

	e.Subject = "布比卡因《 BUP.PUB 》" // 标题

	//e.Text = []byte("test content bcc") // 文本文件内容，二选一
	e.HTML = []byte(fmt.Sprintf("<h1>您的验证码：%s</h1><br><h6>仅在三分钟之内有效哦～</h6>", code)) // html内容

	err := e.Send("smtp.qq.com:587",
		smtp.PlainAuth("", "emal@emal.com", "key", "smtp.qq.com"))
	if err != nil {
		return false
	}
	//_, err = e.AttachFile("test.txt")
	//if err != nil {
	//	return
	//} // 附加文件
	return true
}

// VerificationCode redis校验验证码
func VerificationCode(em string, code string) bool {
	if len(code) == 0 {
		return false
	}
	rdb := common.InitCache()
	regEmail := util.ReEmail(em)
	dbCode, _ := rdb.Get(common.Ctx, regEmail).Result()
	if dbCode == "" || dbCode != code {
		return false
	}
	rdb.Del(common.Ctx, regEmail) // 删除key
	return true
}
