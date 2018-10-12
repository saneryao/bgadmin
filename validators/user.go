package validators

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/saneryao/bgadmin/models"
)

/* 功能：Form表单数据转化成结构体，并校验数据有效性
 * 参数：info是转换后的结果（结构体），ctrl是控制器
 * 返回值：result为true表示转换成功，err不为nil表示转换失败
 */
func ParseUserInfo(info *models.User, ctrl *beego.Controller) (result bool, err error) {
	if info.Profile == nil {
		info.Profile = &models.Profile{}
	}
	// 格式转换
	if err = ctrl.ParseForm(info); err != nil {
		result = false
		return
	}
	if err = ctrl.ParseForm(info.Profile); err != nil {
		result = false
		return
	}
	logs.Debug("%#v-----%#v", info, info.Profile)

	// 数据校验
	valid := validation.Validation{RequiredFirst: true}
	if result, err = valid.RecursiveValid(info); err != nil {
		return
	}
	if !result {
		strErrMsg := ""
		for _, errMsg := range valid.Errors {
			strErrMsg += errMsg.Key + ":" + errMsg.Message + "; "
		}
		err = errors.New(strErrMsg)
		return
	}

	result = true
	return
}
