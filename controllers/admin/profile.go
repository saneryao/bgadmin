package admin

import (
	"github.com/saneryao/bgadmin/controllers"
)

// profile控制器，继承于controllers下的common控制
// 主要负责用户信息的查看和修改
type ProfileController struct {
	controllers.CommonController
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处显示用户信息的页面
 * 参数：空
 * 返回值：空
 */
func (this *ProfileController) Get() {
	this.SetTpl("admin/profile.tpl")
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处显示用户信息的页面
 * 参数：空
 * 返回值：空
 */
func (this *ProfileController) Post() {
}
