package public

import (
	"github.com/astaxie/beego/logs"
	"github.com/saneryao/bgadmin/api/v1"
)

type ResetPwdApi struct {
	v1.CommonApi
}

func (this *ResetPwdApi) Post() {
	logs.Info("!!!!!!!!!!POST!!!!!!!!!!")
	// json := make(map[string]interface{})
	// json["code"] = true
	// json["msg"] = this.Tr("redirecting")
	// json['url'] = this.

	this.Data["json"] = map[string]interface{}{"code": true, "msg": this.Tr("redirecting"), "url": this.URLFor("LoginController.Get")}
	this.ServeJSON()
}
