package public

import (
	"github.com/saneryao/bgadmin/controllers"
	"github.com/saneryao/bgadmin/service"
)

// LogoutController 注销控制器（用于用户注销操作）
type LogoutController struct {
	controllers.CommonController
}

// Get 执行http请求GET方法（处理用户注销操作及注销后页面跳转）
func (api *LogoutController) Get() {
	service.CleanLoginInfo(&api.Controller)
	api.Redirect(api.URLFor("LoginController.Get"), 302)
	api.StopRun()
}
