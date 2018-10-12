package public

import (
	"github.com/saneryao/bgadmin/api/v1"
	"github.com/saneryao/bgadmin/service"
)

type LogoutApi struct {
	v1.CommonApi
}

func (this *LogoutApi) Post() {
	service.SetLogoutInfo(&this.Controller)
	this.Data["json"] = map[string]interface{}{"code": true, "msg": this.Tr("redirecting"), "url": this.URLFor("HomeController.Get")}
	this.ServeJSON()
}
