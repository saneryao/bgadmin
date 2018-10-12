package public

import (
	"github.com/saneryao/bgadmin/controllers"
	"github.com/saneryao/bgadmin/service"
	"github.com/saneryao/bgadmin/validators"
)

// active控制器，继承于controllers下的common控制
// 主要负责激活用户页面的显示
type ActiveController struct {
	controllers.CommonController
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处显示用户激活的页面
 * 参数：空
 * 返回值：空
 */
func (this *ActiveController) Get() {
	var content string
	var err error

	// 获取ID
	var id int64
	if err == nil {
		id, err = validators.ParseIdFromUrl(&this.Controller, ":id")
		if err != nil {
			content = err.Error()
		}
	}

	// 获取Code
	var code string
	if err == nil {
		code, err = validators.ParseEntryFromUrl(&this.Controller, ":code")
		if err != nil {
			content = err.Error()
		}
	}

	// 激活逻辑（在service进行处理）
	if err == nil {
		if err = service.Active(id, code, &this.Controller, this.Lang); err != nil {
			content = err.Error()
		}
	}

	// 设置页面信息
	if err == nil {
		content = this.Tr("active") + this.Tr("success") + ", " + this.Tr("redirect_to") + ":<br/>"
		content += "<a href=\"" + this.URLFor("HomeCtroller.Get") + "\">" + this.Tr("home") + "</a>"
	}

	this.SetTpl("public/active.tpl")
}
