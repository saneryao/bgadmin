package admin

import (
	"strings"
)

// Page控制器，继承于admin下的base控制
// 主要负责后台管理页面的显示
type PageController struct {
	BaseController
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处默认先显示后台管理html框架（加载html、css、js和导航栏、左侧菜单），
 *             再通过page参数显示左侧菜单对应的子页面
 * 参数：空
 * 返回值：空
 */
func (this *PageController) Get() {
	strPage := this.GetString("page")
	switch {
	case strPage == "":
		this.SetTpl("admin/page.tpl",
			"admin/header_admin.tpl",
			"admin/footer_admin.tpl",
			"admin/navbar_admin.tpl",
			"admin/sidebar_admin.tpl")
	case strings.EqualFold(strPage, "index"):
		fallthrough
	case strings.EqualFold(strPage, "user"):
		fallthrough
	case strings.EqualFold(strPage, "role"):
		fallthrough
	case strings.EqualFold(strPage, "menu"):
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
		this.Layout = ""
		this.TplName = "admin/" + strPage + ".tpl"
	default:
		this.Abort("Ajax404")
	}
}
