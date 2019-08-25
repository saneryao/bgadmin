package public

import (
	"github.com/astaxie/beego/logs"
	"github.com/saneryao/bgadmin/api/v1"
)

// ResetPwdAPI 定义一个用户重置密码API（用于用户重置密码等操作）
type ResetPwdAPI struct {
	v1.CommonAPI
}

// Post 执行http请求POST方法（beego定义的接口，处理用户重置密码操作）
func (api *ResetPwdAPI) Post() {
	logs.Info("!!!!!!!!!!POST!!!!!!!!!!")
	// json := make(map[string]interface{})
	// json["code"] = true
	// json["msg"] = api.Tr("redirecting")
	// json['url'] = api.

	api.Data["json"] = map[string]interface{}{"code": true, "msg": api.Tr("redirecting"), "url": api.URLFor("LoginController.Get")}
	api.ServeJSON()
}
