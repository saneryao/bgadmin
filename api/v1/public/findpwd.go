package public

import (
	"github.com/astaxie/beego/logs"
	"github.com/saneryao/bgadmin/api/v1"
)

// FindPwdAPI 定义一个找回密码API（用于用户找回密码等操作）
type FindPwdAPI struct {
	v1.CommonAPI
}

// Post 执行http请求POST方法（处理用户找回密码操作）
func (api *FindPwdAPI) Post() {
	logs.Info("!!!!!!!!!!POST!!!!!!!!!!")
	// json := make(map[string]interface{})
	// json["code"] = true
	// json["msg"] = api.Tr("redirecting")
	// json['url'] = api.

	api.Data["json"] = map[string]interface{}{"code": true, "msg": api.Tr("redirecting"), "url": api.URLFor("ResetPwdController.Get")}
	api.ServeJSON()
}
