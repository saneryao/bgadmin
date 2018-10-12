package sysinit

import (
	"github.com/astaxie/beego"
)

type SmtpConfig struct {
	ServerAddr string
	ServerPort int
	From       string
	Account    string
	Password   string
}

var SmtpCfg SmtpConfig

/* 功能：初始化发送邮件的邮箱设置
 * 参数：空
 * 返回值：空
 */
func initEmail() {
	SmtpCfg.ServerAddr = beego.AppConfig.String("smtp::server_addr")
	SmtpCfg.ServerPort, _ = beego.AppConfig.Int("smtp::server_port")
	SmtpCfg.From = beego.AppConfig.String("smtp::from")
	SmtpCfg.Account = beego.AppConfig.String("smtp::account")
	SmtpCfg.Password = beego.AppConfig.String("smtp::password")
}
