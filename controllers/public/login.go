package public

import (
	"github.com/saneryao/bgadmin/controllers"
	"github.com/saneryao/bgadmin/service"
)

// login控制器，继承于controllers下的common控制
// 主要负责用户登录页面的显示
type LoginController struct {
	controllers.CommonController
}

/* 功能：beego定义的接口，在执行html请求Method方法前执行，
 *             此处检查用户是否登录（已登录的用户跳转到后台首页）
 * 参数：空
 * 返回值：空
 */
func (this *LoginController) Prepare() {
	this.CommonController.Prepare()
	if service.CheckLogin(&this.Controller) {
		this.Redirect(this.URLFor("HomeController.Get"), 302)
		this.StopRun()
	}
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处显示用户登录的页面
 * 参数：空
 * 返回值：空
 */
func (this *LoginController) Get() {
	this.SetTpl("public/login.tpl", "public/header_login.tpl", "public/footer_login.tpl")
}
