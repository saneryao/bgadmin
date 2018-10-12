package public

import (
	"github.com/saneryao/bgadmin/controllers"
)

// resetpwd控制器，继承于controllers下的common控制
// 主要负责用户重置密码页面的显示
type ResetPwdController struct {
	controllers.CommonController
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处显示用户重置密码的页面
 * 参数：空
 * 返回值：空
 */
func (this *ResetPwdController) Get() {
	this.SetTpl("public/resetpwd.tpl", "public/header_resetpwd.tpl", "public/footer_resetpwd.tpl")
}
