package admin

import (
	"fmt"
	"github.com/saneryao/bgadmin/api/v1"
	"github.com/saneryao/bgadmin/service"
)

// baseAPI 定义一个后台基类控制器（admin中其他控制器都继承于它）
type baseAPI struct {
	v1.CommonAPI
}

// Prepare 在执行http请求Method方法前执行（检查用户是否登录及登录后的访问权限）
func (api *baseAPI) Prepare() {
	api.CommonAPI.Prepare() // 执行父类同名函数

	// 检查用户是否已登录（后台admin不允许没有登录就进行访问，未登录直接返回401未授权）
	if !service.CheckLogin(&api.Controller) {
		api.Error = fmt.Errorf(api.Tr("error_401"))
		api.PackResultData(nil, nil)
	}

	// 检查已登录用户的权限（权限不允许时，直接返回403拒绝访问）
	if !service.CheckPermission(&api.Controller) {
		api.Error = fmt.Errorf(api.Tr("error_403"))
		api.PackResultData(nil, nil)
	}
}
