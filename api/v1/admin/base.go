package admin

import (
	"github.com/saneryao/bgadmin/api/v1"
	"github.com/saneryao/bgadmin/service"
)

// 定义一个后台基类控制器（admin中其他控制器都继承于它）
type BaseApi struct {
	v1.CommonApi
}

/* 功能：beego定义的接口，在执行html请求Method方法前执行，
 *             此处检查用户是否登录，及登录后的访问权限
 * 参数：空
 * 返回值：空
 */
func (this *BaseApi) Prepare() {
	this.CommonApi.Prepare() // 执行父类同名函数

	// 检查用户是否已登录（后台admin不允许没有登录就进行访问）
	if !service.CheckLogin(&this.Controller) {
		// 定义返回值
		jsResult := make(map[string]interface{})
		jsResult["code"] = false
		jsResult["error"] = this.Tr("error_401")
		this.Data["json"] = jsResult

		this.ServeJSON()
		this.StopRun()
	}

	// 检查已登录用户的权限
	if !service.CheckPermission(&this.Controller) {
		// 定义返回值
		jsResult := make(map[string]interface{})
		jsResult["code"] = false
		jsResult["error"] = this.Tr("error_403")
		this.Data["json"] = jsResult

		this.ServeJSON()
		this.StopRun()
	}
}
