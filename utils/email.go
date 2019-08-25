package utils

import (
	"encoding/json"
	"github.com/astaxie/beego/utils"
	"github.com/beego/i18n"
	"github.com/saneryao/bgadmin/sysinit"
	"gopkg.in/gomail.v2"
	"strings"
)

// SendEmailActiveUser 发送激活用户的邮件，
// 参数to表示收件人，参数url表示激活链接，参数lang表示国际化语言，
// 返回值error为nil表示发送成功，否则给出错误信息
func SendEmailActiveUser(to, url, lang string) error {
	subject := i18n.Tr(lang, "site_name") + " -- " + i18n.Tr(lang, "active") + i18n.Tr(lang, "user")
	body := i18n.Tr(lang, "dear") + i18n.Tr(lang, "user") + "(" + to + "):<br/>"
	body += "&nbsp;&nbsp;&nbsp;&nbsp;" + i18n.Tr(lang, "welcome_to") + i18n.Tr(lang, "site_name") + i18n.Tr(lang, "site_name") + "!<br/>"
	body += "&nbsp;&nbsp;&nbsp;&nbsp;" + i18n.Tr(lang, "active_tip") + "<br/>"
	body += "&nbsp;&nbsp;&nbsp;&nbsp;" + "<hr/><br/>"
	body += "&nbsp;&nbsp;&nbsp;&nbsp;" + "<a target=_blank href=\"" + url + "\">" + url + "</a>"
	return SendEmailViaConf(to, subject, body)
}

// SendEmailViaConf 发送邮件（通过conf文件的smtp配置进行发送），
// 参数to表示收件人，参数subject表示邮件主题，参数body表示邮件内容，
// 返回值error为nil表示发送成功，否则给出错误信息
func SendEmailViaConf(to, subject, body string) (err error) {
	// // 发送邮件
	// err = SendEmailByBeego(sysinit.SmtpCfg.ServerAddr,
	// 	sysinit.SmtpCfg.ServerPort,
	// 	sysinit.SmtpCfg.Account,
	// 	sysinit.SmtpCfg.Password,
	// 	sysinit.SmtpCfg.From,
	// 	to, subject, body,
	// 	"html")

	// 发送邮件
	email := gomail.NewMessage()
	email.SetHeader("From", sysinit.SMTPCfg.ServerAddr)
	email.SetHeader("To", to)
	email.SetHeader("Subject", subject)
	email.SetBody("text/html", body)
	dialer := gomail.NewDialer(sysinit.SMTPCfg.ServerAddr, sysinit.SMTPCfg.ServerPort, sysinit.SMTPCfg.Account, sysinit.SMTPCfg.Password)
	err = dialer.DialAndSend(email)
	return
}

// SendEmailByParameters 发送邮件（通过参数设置smtp配置），
// 参数host表示smtp服务器地址，参数port表示smtp服务器端口，
// 参数pwd表示smtp密码，参数from表示发件人（相当于smtp账号的昵称），
// 参数to表示收件人，参数subject表示邮件主题，参数body表示邮件内容，
// 返回值error为nil表示发送成功，否则给出错误信息
func SendEmailByParameters(host string, port int, accout, pwd, from, to, subject, body, mailtype string) (err error) {
	config := make(map[string]interface{})
	config["username"] = accout
	config["password"] = pwd
	config["host"] = host
	config["port"] = port
	var jsCfg []byte
	jsCfg, err = json.Marshal(config)
	if err != nil {
		return
	}

	email := utils.NewEMail(string(jsCfg))
	email.To = []string{to}
	email.From = from
	email.Subject = subject
	email.Text = body
	if strings.EqualFold(mailtype, "html") {
		email.HTML = body
	} else {
		email.Text = body
	}
	err = email.Send()
	return
}
