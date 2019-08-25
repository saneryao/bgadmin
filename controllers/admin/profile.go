package admin

import (
	"github.com/saneryao/bgadmin/controllers"
)

// ProfileController 定义一个用户信息控制器（用于显示用户信息页面等）
type ProfileController struct {
	controllers.CommonController
}

// Get 执行http请求GET方法（beego定义的接口，显示用户信息页面）
func (api *ProfileController) Get() {
	api.SetTpl("admin/profile.tpl")
}
