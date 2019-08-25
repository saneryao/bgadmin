package admin

import (
	"github.com/saneryao/bgadmin/controllers"
)

// ModifyPwdController 定义一个修改密码控制器（用于显示修改密码页面等）
type ModifyPwdController struct {
	controllers.CommonController
}

// Get 执行http请求GET方法（beego定义的接口，显示修改密码页面）
func (api *ModifyPwdController) Get() {
	api.SetTpl("admin/modifypwd.tpl")
}
