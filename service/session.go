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

// SetLoginInfo 登录后设置用户信息到session
func SetLoginInfo(ctrl *beego.Controller, info interface{}) {
	ctrl.SetSession("USER_INFO", info)
}

// GetLoginInfo：从session中获取用户信息
func GetLoginInfo(ctrl *beego.Controller) (user models.User) {
	info := ctrl.GetSession("USER_INFO")

	// 类型断言
	if info != nil {
		user, _ = info.(models.User)
	}
	return
}

// SetLogoutInfo 注销时清空session信息
func SetLogoutInfo(ctrl *beego.Controller) {
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
