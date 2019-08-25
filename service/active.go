package service

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"github.com/saneryao/bgadmin/models"
	"time"
)

// Active 用户激活的处理逻辑（通过用户ID、激活码进行验证）
func Active(id int64, code string, ctrl *beego.Controller, lang string) (err error) {
	// 获取已存储的激活码
	db := orm.NewOrm()
	ac := models.Code{UserID: id}
	if err = db.Read(ac, "UserID"); err != nil {
		return
	}

	// 判断有效期
	if ac.CreateTime > time.Now().Unix()-86400 {
		err = errors.New(i18n.Tr(lang, "active_code") + i18n.Tr(lang, "expired"))
		return
	}

	return
}
