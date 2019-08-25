package public

import (
	"github.com/saneryao/bgadmin/api/v1"
	"github.com/saneryao/bgadmin/service"
	"github.com/saneryao/bgadmin/validators"
)

// RegisterAPI 定义一个用户注册API（用于用户注册等操作）
type RegisterAPI struct {
	v1.CommonAPI
}

// Post 执行http请求POST方法（beego定义的接口，处理用户注册操作）
func (api *RegisterAPI) Post() {
	// 包装并处理返回结果
	others := make(map[string]interface{})
	defer func() {
		api.PackResultData(nil, others)
	}()

	// 获取并校验输入信息
	params := validators.RegisterParams{}
	if _, api.Error = validators.ParseRegisterParams(&params, &api.Controller); api.Error != nil {
		return
	}

	// 注册逻辑（在service层进行处理）
	if api.Error = service.Register(&params, &api.Controller, api.Lang); api.Error != nil {
		return
	}

	// 返回数据给客户端
	others["url"] = api.URLFor("HomeController.Get")
	return
}
