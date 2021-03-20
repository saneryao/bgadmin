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
	"os"
)

// Login 用户登录的处理逻辑（验证用户名、密码和验证码）
func Login(params *validators.LoginParams, ctrl *beego.Controller, lang string) (err error) {
	// 校验验证码
	if !sysinit.VerifyCode.Verify(params.IDCaptcha, params.Captcha) {
		err = errors.New(i18n.Tr(lang, "captcha_error"))
		return
	}

	// 校验用户名和密码
	user := models.User{}
	db := orm.NewOrm()
	user.Name = params.Username
	err = db.Read(&user, "Name")
	if err != nil {
		if err == os.ErrNotExist {
			err = errors.New(i18n.Tr(lang, "uname_or_pwd_error"))
		}
		return
	}
	if user.State != 1 {
		err = errors.New(i18n.Tr(lang, "user_is_disabled"))
		return
	}
	if err != nil || !utils.VerifyPassword(params.Password, user.Pwd) {
		err = errors.New(i18n.Tr(lang, "uname_or_pwd_error"))
		return
	}

	// 设置登录Session信息
	SetLoginUser(ctrl, user)
	return
}
