package public

import (
	"github.com/saneryao/bgadmin/api/v1"
	"github.com/saneryao/bgadmin/service"
	"github.com/saneryao/bgadmin/validators"
)

// LoginAPI 定义一个用户登录API（用于用户登录时用户名和密码校验等）
type LoginAPI struct {
	v1.CommonAPI
}

// Post 执行http请求POST方法（处理用户登录操作）
func (api *LoginAPI) Post() {
	// 包装并处理返回结果
	others := make(map[string]interface{})
	defer func() {
		api.PackResultData(nil, others)
	}()

	// 获取并校验输入信息
	params := validators.LoginParams{}
	if _, api.Error = validators.ParseLoginParams(&params, &api.Controller); api.Error != nil {
		return
	}

	// 登录逻辑（在service层进行处理）
	if api.Error = service.Login(&params, &api.Controller, api.Lang); api.Error != nil {
		return
	}

	// 返回数据给客户端
	others["url"] = api.URLFor("HomeController.Get")
	return
}
