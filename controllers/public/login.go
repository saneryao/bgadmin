package public

import (
	"github.com/saneryao/bgadmin/controllers"
	"github.com/saneryao/bgadmin/service"
)

// LoginController 登录控制器（用于显示用户登录页面等）
type LoginController struct {
	controllers.CommonController
}

// Prepare 在执行http请求Method方法前执行（beego定义的接口， 检查用户是否登录，已登录的用户跳转到后台首页）
func (api *LoginController) Prepare() {
	api.CommonController.Prepare()
	if service.CheckLogin(&api.Controller) {
		api.Redirect(api.URLFor("HomeController.Get"), 302)
		api.StopRun()
	}
}

// Get 执行http请求GET方法（beego定义的接口，显示用户登录页面）
func (api *LoginController) Get() {
	api.SetTpl("public/login.tpl", "public/header_login.tpl", "public/footer_login.tpl")
}
