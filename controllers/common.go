package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/utils"
	"strings"
)

// CommonController 定义一个基类控制器（其他控制器都继承于它）
type CommonController struct {
	beego.Controller
	i18n.Locale
	UserInfo models.User
}

// SetTpl 定义一个给子控制器通用的接口，用来设置模板布局、布局依赖的文件、模板文件
// 参数：0-5个字符串
// 第1个是模板文件，若没有设置此参数会使用默认值，默认值是“控制器名/操作.tpl”，
// 第2个是自定义头信息，第3个是自定义脚信息，
// 第4个是自定义导航栏信息，第5个是自定义左侧菜单信息）
func (api *CommonController) SetTpl(template ...string) {
	api.Data["PageTitle"] = api.Tr("site_name")
	api.Layout = "layout/base.tpl"
	if len(template) > 1 {
		api.LayoutSections = make(map[string]string)
		api.LayoutSections["CustomHeader"] = template[1]
		if len(template) > 2 {
			api.LayoutSections["CustomFooter"] = template[2]
		}
		if len(template) > 3 {
			api.LayoutSections["CustomNavbar"] = template[3]
		}
		if len(template) > 4 {
			api.LayoutSections["CustomSidebar"] = template[4]
		}
	}
	if len(template) > 0 {
		api.TplName = template[0]
	} else {
		ctrlName, actionName := api.GetControllerAndAction()
		ctrlName = strings.ToLower(ctrlName)
		actionName = strings.ToLower(actionName)
		api.TplName = ctrlName + "/" + actionName + ".tpl"
	}
}

// Prepare 在执行http请求Method方法前执行（API接口）
// 此处进行国际化语言设置，并获取控制器信息给子控制器使用
func (api *CommonController) Prepare() {
	_, api.Lang = utils.SetLangVer(&api.Controller)
	ctrlName, actionName := api.GetControllerAndAction()
	// api.Data["CtrlName"] = ctrlName
	// api.Data["ActionName"] = actionName
	logs.Debug("-----" + ctrlName + "(" + actionName + "):" + api.Lang + "-----")
}
