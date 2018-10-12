package admin

import (
	"github.com/saneryao/bgadmin/controllers"
)

// modifypwd控制器，继承于controllers下的common控制
// 主要负责用户信息的查看和修改
type ModifyPwdController struct {
	controllers.CommonController
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处显示用户信息的页面
 * 参数：空
 * 返回值：空
 */
func (this *ModifyPwdController) Get() {
	this.SetTpl("admin/modifypwd.tpl")
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处显示用户信息的页面
 * 参数：空
 * 返回值：空
 */
func (this *ModifyPwdController) Post() {
}
