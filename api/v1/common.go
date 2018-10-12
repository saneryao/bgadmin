package v1

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/utils"
)

// 定义一个基类API（其他API都继承于它）
type CommonApi struct {
	beego.Controller
	i18n.Locale
	CtrlName   string
	ActionName string
	UserInfo   models.User
	Error      error
}

/* 功能：beego定义的接口，在执行html请求Method方法前执行，
 *             此处进行国际化语言设置，并获取控制器信息给子控制器使用
 * 参数：空
 * 返回值：空
 */
func (this *CommonApi) Prepare() {
	_, this.Lang = utils.SetLangVer(&this.Controller)
	this.CtrlName, this.ActionName = this.GetControllerAndAction()
	this.Data["CtrlName"] = this.CtrlName
	this.Data["ActionName"] = this.ActionName
	logs.Debug("=====" + this.CtrlName + "(" + this.ActionName + "):" + this.Lang + "=====")
}

/* 功能：包装并处理返回结果
 * 参数：data是数据集合，others是其他参数
 * 返回值：result为true表示转换成功，err不为nil表示转换失败
 */
func (this *CommonApi) PackResultData(data interface{}, others map[string]interface{}) {
	jsResult := make(map[string]interface{})
	if this.Error != nil {
		jsResult["code"] = false
		jsResult["error"] = this.Error.Error()
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
	this.Data["json"] = jsResult
	this.ServeJSON()
	this.StopRun()
}
