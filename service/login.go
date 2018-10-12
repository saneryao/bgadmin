package service

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/sysinit"
	"github.com/saneryao/bgadmin/utils"
	"github.com/saneryao/bgadmin/validators"
)

/* 功能：登录逻辑（判断用户名、密码和验证码）
 * 参数：params登录Form表单数据，ctrl控制器
 * 返回值：err错误值（登录成功时值为nil，登录失败时给出错误信息）
 */
func Login(params *validators.LoginParams, ctrl *beego.Controller, lang string) (err error) {
	// 校验验证码
	if !sysinit.VerifyCode.Verify(params.IdCaptcha, params.Captcha) {
		err = errors.New(i18n.Tr(lang, "captcha_error"))
		return
	}

	// 校验用户名和密码
	user := models.User{}
	db := orm.NewOrm()
	user.Name = params.Username
	err = db.Read(&user, "Name")
	if err != nil || !utils.VerifyPassword(params.Password, user.Pwd) {
		err = errors.New(i18n.Tr(lang, "uname_or_pwd_error"))
		return
	}

	// 设置登录Session信息
	SetLoginInfo(ctrl, user)
	return
}
