package service

import (
	"github.com/astaxie/beego"
	"github.com/saneryao/bgadmin/models"
)

// CheckLogin 从session中检测用户是否已登录
// 返回true表示已登录，返回false表示没有登录
func CheckLogin(ctrl *beego.Controller) (result bool) {
	info := ctrl.GetSession("USER_INFO")

	// 类型断言
	if info != nil {
		_, result = info.(models.User)
	}
	return
}

// SetLoginUser 登录后设置用户信息到session
func SetLoginUser(ctrl *beego.Controller, info models.User) {
	ctrl.SetSession("USER_INFO", info)
}

// GetLoginUser：从session中获取用户信息
func GetLoginUser(ctrl *beego.Controller) (user models.User) {
	info := ctrl.GetSession("USER_INFO")

	// 类型断言
	if info != nil {
		user, _ = info.(models.User)
	}
	return
}

// SetUserMenus 登录后设置用户的菜单信息到session
func SetUserMenus(ctrl *beego.Controller, info []models.Menu) {
	ctrl.SetSession("USER_MENUS", info)
}

// GetUserMenus：从session中获取用户信息
func GetUserMenus(ctrl *beego.Controller) (menus []models.Menu) {
	info := ctrl.GetSession("USER_MENUS")

	// 类型断言
	if info != nil {
		menus, _ = info.([]models.Menu)
	}
	return
}

// SetUserLinks 登录后设置用户信息到session
func SetUserLinks(ctrl *beego.Controller, info []models.Link) {
	ctrl.SetSession("USER_LINKS", info)
}

// GetUserLinks：从session中获取用户信息
func GetUserLinks(ctrl *beego.Controller) (links []models.Link) {
	info := ctrl.GetSession("USER_LINKS")

	// 类型断言
	if info != nil {
		links, _ = info.([]models.Link)
	}
	return
}

// CleanLoginInfo 注销时清空session信息
func CleanLoginInfo(ctrl *beego.Controller) {
	ctrl.DestroySession()
}

// SetSessionInfo 把信息设置到session
func SetSessionInfo(ctrl *beego.Controller, name, info interface{}) {
	ctrl.SetSession(name, info)
}

// GetSessionInfo 从session中获取信息
func GetSessionInfo(ctrl *beego.Controller, name interface{}) (info interface{}) {
	info = ctrl.GetSession(name)
	return
}
