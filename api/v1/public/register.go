package public

import (
	"github.com/saneryao/bgadmin/api/v1"
	"github.com/saneryao/bgadmin/service"
	"github.com/saneryao/bgadmin/validators"
)

type RegisterApi struct {
	v1.CommonApi
}

func (this *RegisterApi) Post() {
	// 包装并处理返回结果
	others := make(map[string]interface{})
	defer func() {
		this.PackResultData(nil, others)
	}()

	// 获取并校验输入信息
	params := validators.RegisterParams{}
	if _, this.Error = validators.ParseRegisterParams(&params, &this.Controller); this.Error != nil {
		return
	}

	// 注册逻辑（在service层进行处理）
	if this.Error = service.Register(&params, &this.Controller, this.Lang); this.Error != nil {
		return
	}

	// 返回数据给客户端
	others["url"] = this.URLFor("HomeController.Get")
	return
}
