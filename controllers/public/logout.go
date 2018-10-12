package public

import (
	"github.com/saneryao/bgadmin/controllers"
	"github.com/saneryao/bgadmin/service"
)

// logout控制器，继承于controllers下的common控制
// 主要负责用户注销操作
type LogoutController struct {
	controllers.CommonController
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处进行用户注销操作（注销后进行页面跳转）
 * 参数：空
 * 返回值：空
 */
func (this *LogoutController) Get() {
	service.SetLogoutInfo(&this.Controller)
	this.Redirect(this.URLFor("LoginController.Get"), 302)
	this.StopRun()
}
