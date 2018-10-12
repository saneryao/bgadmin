package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/captcha"
)

var VerifyCode *captcha.Captcha

/* 功能：初始化验证码设置
 * 参数：空
 * 返回值：空
 */
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
