package public

import (
	"github.com/saneryao/bgadmin/api/v1"
	"github.com/saneryao/bgadmin/service"
	"github.com/saneryao/bgadmin/validators"
)

type LoginApi struct {
	v1.CommonApi
}

func (this *LoginApi) Post() {
	// 包装并处理返回结果
	others := make(map[string]interface{})
	defer func() {
		this.PackResultData(nil, others)
	}()

	// 获取并校验输入信息
	params := validators.LoginParams{}
	if _, this.Error = validators.ParseLoginParams(&params, &this.Controller); this.Error != nil {
		return
	}

	// 登录逻辑（在service层进行处理）
	if this.Error = service.Login(&params, &this.Controller, this.Lang); this.Error != nil {
		return
	}

	// 返回数据给客户端
	others["url"] = this.URLFor("HomeController.Get")
	return
}
