package public

import (
	"github.com/saneryao/bgadmin/controllers"
)

// ResetPwdController 重置密码控制器（用于显示用户重置密码页面等）
type ResetPwdController struct {
	controllers.CommonController
}

// Get 执行http请求GET方法（显示用户重置密码页面）
func (api *ResetPwdController) Get() {
	api.SetTpl("public/resetpwd.tpl", "public/header_resetpwd.tpl", "public/footer_resetpwd.tpl")
}
