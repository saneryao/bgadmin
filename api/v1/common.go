package v1

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/utils"
)

// CommonAPI 定义一个基类API（其他API都继承于它）
type CommonAPI struct {
	beego.Controller
	i18n.Locale
	UserInfo models.User
	Error    error
}

// Prepare 在执行http请求Method方法前执行（API接口）
// 此处进行国际化语言设置，并获取控制器信息给子控制器使用
func (api *CommonAPI) Prepare() {
	_, api.Lang = utils.SetLangVer(&api.Controller)
	ctrlName, actionName := api.GetControllerAndAction()
	// api.Data["CtrlName"] = ctrlName
	// api.Data["ActionName"] = actionName
	logs.Debug("=====" + ctrlName + "(" + actionName + "):" + api.Lang + "=====")
}

// PackResultData 包装并处理返回结果
// 返回给前台json格式的结果信息
func (api *CommonAPI) PackResultData(data interface{}, others map[string]interface{}) {
	jsResult := make(map[string]interface{})
	if api.Error != nil {
		jsResult["code"] = false
		jsResult["error"] = api.Error.Error()
	} else {
		jsResult["code"] = true
		if data == nil {
			jsResult["data"] = make(map[string]interface{})
		} else {
			jsResult["data"] = data
		}
	}
	if others != nil {
		for k, v := range others {
			jsResult[k] = v
		}
	}
	api.Data["json"] = jsResult
	api.ServeJSON()
	api.StopRun()
}
