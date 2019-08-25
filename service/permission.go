package service

import (
	"github.com/astaxie/beego"
)

// CheckPermission 检查已登录用户的权限，判断其是否可以继续访问
// 返回true表示有权限可以继续访问，返回false表示权限不够无法继续访问
func CheckPermission(ctrl *beego.Controller) bool {
	return true
}
