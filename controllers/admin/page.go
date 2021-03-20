package admin

import (
	"github.com/saneryao/bgadmin/service"
	"strings"
)

// PageController 定义一个页面控制器（用于后台管理页面的渲染等）
type PageController struct {
	baseController
}

// Get 执行http请求GET方法（渲染页面）
// 默认先显示后台管理html框架（加载html、css、js和导航栏、左侧菜单），
// 再通过page参数显示左侧菜单对应的子页面
func (api *PageController) Get() {
	strPage := api.GetString("page")
	switch {
	case strPage == "":
		api.SetTpl("admin/page.tpl",
			"admin/header_admin.tpl",
			"admin/footer_admin.tpl",
			"admin/navbar_admin.tpl",
			"admin/sidebar_admin.tpl")
		menus := service.GetUserMenus(&api.Controller)
		if len(menus) <= 0 {
			api.Abort("Ajax401")
		}
		api.Data["Menus"] = menus
	case strings.EqualFold(strPage, "index"):
		fallthrough
	case strings.EqualFold(strPage, "user"):
		fallthrough
	case strings.EqualFold(strPage, "role"):
		fallthrough
	case strings.EqualFold(strPage, "menu"):
		fallthrough
	case strings.EqualFold(strPage, "link"):
		fallthrough
	case strings.EqualFold(strPage, "demo"):
		fallthrough
	case strings.EqualFold(strPage, "grid"):
		fallthrough
	case strings.EqualFold(strPage, "button"):
		fallthrough
	case strings.EqualFold(strPage, "icon"):
		fallthrough
	case strings.EqualFold(strPage, "resetpwd"):
		fallthrough
	case strings.EqualFold(strPage, "profile"):
		api.Layout = ""
		api.TplName = "admin/" + strPage + ".tpl"
		menus := service.GetUserMenus(&api.Controller)
		if len(menus) <= 0 {
			api.Abort("Ajax401")
		}
		api.Data["Menus"] = menus
	default:
		api.Abort("Ajax404")
	}
}
