package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/captcha"
)

// VerifyCode 定义全局可用的验证码校验器
var VerifyCode *captcha.Captcha

// initCaptcha 初始化验证码设置
func initCaptcha() {
	var err error
	VerifyCode = captcha.NewWithFilter("/captcha/", CacheRedis)
	VerifyCode.ChallengeNums, err = beego.AppConfig.Int("captcha::length")
	if err != nil {
		logs.Error(err)
	}
	VerifyCode.StdWidth, err = beego.AppConfig.Int("captcha::width")
	if err != nil {
		logs.Error(err)
	}
	VerifyCode.StdHeight, err = beego.AppConfig.Int("captcha::height")
	if err != nil {
		logs.Error(err)
	}
}
