package public

import (
	"github.com/saneryao/bgadmin/controllers"
	"github.com/saneryao/bgadmin/service"
	"github.com/saneryao/bgadmin/validators"
)

// ActiveController 定义一个激活控制器（用于显示用户激活页面等）
type ActiveController struct {
	controllers.CommonController
}

// Get 执行http请求GET方法（显示用户激活页面）
func (api *ActiveController) Get() {
	var content string
	var err error

	// 获取ID
	var id int64
	if err == nil {
		id, err = validators.ParseIDFromURL(&api.Controller, ":id")
		if err != nil {
			content = err.Error()
		}
	}

	// 获取Code
	var code string
	if err == nil {
		code, err = validators.ParseEntryFromURL(&api.Controller, ":code")
		if err != nil {
			content = err.Error()
		}
	}

	// 激活逻辑（在service进行处理）
	if err == nil {
		if err = service.Active(id, code, &api.Controller, api.Lang); err != nil {
			content = err.Error()
		}
	}

	// 设置页面信息
	if err == nil {
		content = api.Tr("active") + api.Tr("success") + ", " + api.Tr("redirect_to") + ":<br/>"
		content += "<a href=\"" + api.URLFor("HomeCtroller.Get") + "\">" + api.Tr("home") + "</a>"
	}

	api.SetTpl("public/active.tpl")
}
