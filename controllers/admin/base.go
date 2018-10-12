package admin

import (
	"github.com/saneryao/bgadmin/controllers"
	"github.com/saneryao/bgadmin/service"
)

// 定义一个后台基类控制器（admin中其他控制器都继承于它）
type BaseController struct {
	controllers.CommonController
}

/* 功能：beego定义的接口，在执行html请求Method方法前执行，
 *             此处检查用户是否登录，及登录后的访问权限
 * 参数：空
 * 返回值：空
 */
func (this *BaseController) Prepare() {
	this.CommonController.Prepare() // 执行父类同名函数

	// 检查用户是否已登录（后台admin不允许没有登录就进行访问）
	if !service.CheckLogin(&this.Controller) {
		// !!! ace_ajax请求子页面，若直接跳转到登录页面，会导致登录页显示在子页面中，并循环刷新
		if this.Ctx.Input.IsAjax() {
			this.Abort("Ajax401")
		}
		this.Redirect(this.URLFor("LoginController.Get"), 302)
		this.StopRun()
	}

	// 设置用户信息
	this.UserInfo = service.GetLoginInfo(&this.Controller)
	this.Data["UserName"] = this.UserInfo.Name
	if this.UserInfo.Profile == nil || this.UserInfo.Profile.Nick == "" {
		this.Data["UserNick"] = this.UserInfo.Name
	} else {
		this.Data["UserNick"] = this.UserInfo.Profile.Nick
	}

	// 检查已登录用户的权限
	if !service.CheckPermission(&this.Controller) {
		this.Abort("403")
	}
}
