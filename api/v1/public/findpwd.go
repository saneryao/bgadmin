package public

import (
	"github.com/astaxie/beego/logs"
	"github.com/saneryao/bgadmin/api/v1"
)

type FindPwdApi struct {
	v1.CommonApi
}

func (this *FindPwdApi) Post() {
	logs.Info("!!!!!!!!!!POST!!!!!!!!!!")
	// json := make(map[string]interface{})
	// json["code"] = true
	// json["msg"] = this.Tr("redirecting")
	// json['url'] = this.

	this.Data["json"] = map[string]interface{}{"code": true, "msg": this.Tr("redirecting"), "url": this.URLFor("ResetPwdController.Get")}
	this.ServeJSON()
}
