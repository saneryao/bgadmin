package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/utils"
	"strings"
)

// 定义一个基类控制器（其他控制器都继承于它）
type CommonController struct {
	beego.Controller
	i18n.Locale
	CtrlName   string
	ActionName string
	UserInfo   models.User
}

/* 功能：定义一个给子控制器通用的接口，用来设置模板布局、布局依赖的文件、模板文件
 * 参数：字符串不定参数，可以设置0-5个参数
 *             （第1个是模板文件，若没有设置此参数会使用默认值，默认值是“控制器名/操作.tpl”，
 *                 第2个是自定义头信息，第3个是自定义脚信息，
 *                 第4个是自定义导航栏信息，第5个是自定义左侧菜单信息）
 * 返回值：空
 */
func (this *CommonController) SetTpl(template ...string) {
	this.Data["PageTitle"] = this.Tr("site_name")
	this.Layout = "layout/base.tpl"
	if len(template) > 1 {
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["CustomHeader"] = template[1]
		if len(template) > 2 {
			this.LayoutSections["CustomFooter"] = template[2]
		}
		if len(template) > 3 {
			this.LayoutSections["CustomNavbar"] = template[3]
		}
		if len(template) > 4 {
			this.LayoutSections["CustomSidebar"] = template[4]
		}
	}
	if len(template) > 0 {
		this.TplName = template[0]
	} else {
		strCtrlName := strings.ToLower(this.CtrlName[0 : len(this.CtrlName)-10]) // 去掉Controller这个10个字母
		strActionName := strings.ToLower(this.ActionName)
		this.TplName = strCtrlName + "/" + strActionName + ".tpl"
	}
}

/* 功能：beego定义的接口，在执行html请求Method方法前执行，
 *             此处进行国际化语言设置，并获取控制器信息给子控制器使用
 * 参数：空
 * 返回值：空
 */
func (this *CommonController) Prepare() {
	_, this.Lang = utils.SetLangVer(&this.Controller)
	this.CtrlName, this.ActionName = this.GetControllerAndAction()
	this.Data["CtrlName"] = this.CtrlName
	this.Data["ActionName"] = this.ActionName
	logs.Debug("-----" + this.CtrlName + "(" + this.ActionName + "):" + this.Lang + "-----")
}
