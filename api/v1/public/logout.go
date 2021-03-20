package public

import (
	"github.com/saneryao/bgadmin/api/v1"
	"github.com/saneryao/bgadmin/service"
)

// LogoutAPI 定义一个用户注销API（用于用户注销时清理会话信息等操作）
type LogoutAPI struct {
	v1.CommonAPI
}

// Post 执行http请求POST方法（处理用户注销操作）
func (api *LogoutAPI) Post() {
	service.CleanLoginInfo(&api.Controller)
	api.Data["json"] = map[string]interface{}{"code": true, "msg": api.Tr("redirecting"), "url": api.URLFor("HomeController.Get")}
	api.ServeJSON()
}
