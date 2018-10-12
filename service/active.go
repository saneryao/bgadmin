package service

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"github.com/saneryao/bgadmin/models"
	"time"
)

/* 功能：激活逻辑（通过用户ID、激活码进行验证）
 * 参数：params登录Form表单数据，ctrl控制器
 * 返回值：err错误值（登录成功时值为nil，登录失败时给出错误信息）
 */
func Active(id int64, code string, ctrl *beego.Controller, lang string) (err error) {
	// 获取已存储的激活码
	db := orm.NewOrm()
	ac := models.Code{UserId: id}
	if err = db.Read(ac, "UserId"); err != nil {
		return
	}

	// 判断有效期
	if ac.CreateTime > time.Now().Unix()-86400 {
		err = errors.New(i18n.Tr(lang, "active_code") + i18n.Tr(lang, "expired"))
		return
	}

	return
}
