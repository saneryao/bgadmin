package admin

import (
	"github.com/saneryao/bgadmin/controllers"
	"github.com/saneryao/bgadmin/service"
)

// baseController 定义一个后台基类控制器（admin中其他控制器都继承于它）
type baseController struct {
	controllers.CommonController
}

// Prepare 在执行http请求Method方法前执行（检查用户是否登录及登录后的访问权限）
func (api *baseController) Prepare() {
	api.CommonController.Prepare() // 执行父类同名函数

	// 检查用户是否已登录（后台admin不允许没有登录就进行访问）
	if !service.CheckLogin(&api.Controller) {
		// !!! ace_ajax请求子页面，若直接跳转到登录页面，会导致登录页显示在子页面中，并循环刷新
		if api.Ctx.Input.IsAjax() {
			api.Abort("Ajax401")
		}
		api.Redirect(api.URLFor("LoginController.Get"), 302)
		api.StopRun()
	}

	// 设置用户信息
	api.UserInfo = service.GetLoginInfo(&api.Controller)
	api.Data["UserName"] = api.UserInfo.Name
	if api.UserInfo.Profile == nil || api.UserInfo.Profile.Nick == "" {
		api.Data["UserNick"] = api.UserInfo.Name
	} else {
		api.Data["UserNick"] = api.UserInfo.Profile.Nick
	}

	// 检查已登录用户的权限
	if !service.CheckPermission(&api.Controller) {
		api.Abort("403")
	}
}
