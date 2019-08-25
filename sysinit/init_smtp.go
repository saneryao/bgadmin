package sysinit

import (
	"github.com/astaxie/beego"
)

// SMTPConfig 定义smtp设置
type SMTPConfig struct {
	ServerAddr string
	ServerPort int
	From       string
	Account    string
	Password   string
}

// SMTPCfg 定义全局可用的smtp设置
var SMTPCfg SMTPConfig

// initSMTP 初始化发送邮件的邮箱设置
func initSMTP() {
	SMTPCfg.ServerAddr = beego.AppConfig.String("smtp::server_addr")
	SMTPCfg.ServerPort, _ = beego.AppConfig.Int("smtp::server_port")
	SMTPCfg.From = beego.AppConfig.String("smtp::from")
	SMTPCfg.Account = beego.AppConfig.String("smtp::account")
	SMTPCfg.Password = beego.AppConfig.String("smtp::password")
}
