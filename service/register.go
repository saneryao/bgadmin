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
	"strings"
	"time"
)

// Register 用户注册的处理逻辑（判断邮箱、用户名、密码和验证码）
func Register(params *validators.RegisterParams, ctrl *beego.Controller, lang string) (err error) {
	// 校验验证码
	if !sysinit.VerifyCode.Verify(params.IDCaptcha, params.Captcha) {
		err = errors.New(i18n.Tr(lang, "captcha_error"))
		return
	}

	// 判断两次密码一致性
	if !strings.EqualFold(params.PwdConfirm, params.Password) {
		err = errors.New(i18n.Tr(lang, "two_passwords_differ"))
		return
	}

	// 检查邮箱是否已注册
	profile := models.Profile{Email: params.Email}
	db := orm.NewOrm()
	if err = db.Read(&profile, "Email"); err == nil {
		err = errors.New(i18n.Tr(lang, "email") + i18n.Tr(lang, "is_already_exists"))
		return
	}

	// 检查用户名是否已注册
	user := models.User{Name: params.Username}
	if err = db.Read(&user, "Name"); err == nil {
		err = errors.New(i18n.Tr(lang, "username") + i18n.Tr(lang, "is_already_exists"))
		return
	}

	// 保存用户信息
	if _, err = db.Insert(&profile); err != nil {
		return
	}
	user.Profile = &profile
	user.Pwd = utils.EncryptPassword(params.Password)
	if _, err = db.Insert(&user); err != nil {
		return
	}

	// 生成并保存激活码
	code := models.Code{UserID: user.ID, Code: utils.GetRandomString(64), CreateTime: time.Now().Unix()}
	if _, err = db.Insert(&code); err != nil {
		return
	}

	// 发送激活邮件
	url := ctrl.Ctx.Request.URL.EscapedPath()
	url += ctrl.URLFor("ActiveController.Get", ":id", user.ID, ":code", code)
	if err = utils.SendEmailActiveUser(params.Email, url, lang); err != nil {
		return
	}

	// 操作成功，返回相应数据
	return
}
