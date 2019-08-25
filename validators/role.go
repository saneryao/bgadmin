package validators

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/saneryao/bgadmin/models"
)

// ParseRoleInfo 解析表单数据中的角色信息，并校验数据有效性
func ParseRoleInfo(info *models.Role, ctrl *beego.Controller) (result bool, err error) {
	// 格式转换
	if err = ctrl.ParseForm(info); err != nil {
		result = false
		return
	}
	logs.Debug("%#v", info)

	// 数据校验
	valid := validation.Validation{}
	if result, err = valid.Valid(info); err != nil {
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
