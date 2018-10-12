package service

import (
	"github.com/astaxie/beego"
	"github.com/saneryao/bgadmin/models"
)

/* 功能：定义一个检查是否登录的接口给子控制器使用
 * 参数：ctrl 控制器
 * 返回值：result 是否已登录（true表示已登录，false表示没有登录）
 *                 user 登录后返回的用户信息
 */
func CheckLogin(ctrl *beego.Controller) (result bool) {
	info := ctrl.GetSession("USER_INFO")

	// 类型断言
	if info != nil {
		_, result = info.(models.User)
	}
	return
}

/* 功能：登录成功后，把用户信息设置到session
 * 参数：user 用户信息
 *             info 信息
 * 返回值：空
 */
func SetLoginInfo(ctrl *beego.Controller, info interface{}) {
	ctrl.SetSession("USER_INFO", info)
}

/* 功能：获取登录成功后设置到session的用户信息
 * 参数：user 用户信息
 *             info 信息
 * 返回值：空
 */
func GetLoginInfo(ctrl *beego.Controller) (user models.User) {
	info := ctrl.GetSession("USER_INFO")

	// 类型断言
	if info != nil {
		user, _ = info.(models.User)
	}
	return
}

/* 功能：注销时，清空session信息
 * 参数：空
 * 返回值：空
 */
func SetLogoutInfo(ctrl *beego.Controller) {
	ctrl.DestroySession()
}

/* 功能：把信息设置到session
 * 参数：user 用户信息
 *             info 信息
 * 返回值：空
 */
func SetSessionInfo(ctrl *beego.Controller, name, info interface{}) {
	ctrl.SetSession(name, info)
}

/* 功能：从session获取信息
 * 参数：name 键值对中的关键字key
 * 返回值：info 键值对中的值
 */
func GetSessionInfo(ctrl *beego.Controller, name interface{}) (info interface{}) {
	info = ctrl.GetSession(name)
	return
}
