package validators

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

// FilterParams 数据分页信息等参数
type FilterParams struct {
	Draw         int    `form:"draw" valid:"Min(0)"`
	Limit        int    `form:"limit" valid:"Range(1, 1000)"`
	Offset       int    `form:"offset" valid:"Min(0)"`
	Page         int    `form:"page" valid:"Min(0)"`
	PerPage      int    `form:"per_page" valid:"Min(1)"`
	SortBy       string `form:"sortby""`
	Order        string `form:"order"`
	SearchColumn string `form:"search_column"`
	SearchValue  string `form:"search_value"`
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
// func (this *FilterParams) Valid(valid *validation.Validation) {
// }

// ParseFilterParams 转换表单数据为结构体，并校验数据有效性
func ParseFilterParams(paramsInfo *FilterParams, ctrl *beego.Controller) (result bool, err error) {
	// 格式转换
	if err = ctrl.ParseForm(paramsInfo); err != nil {
		result = false
		return
	}
	logs.Debug("%#v", paramsInfo)
	if paramsInfo.Limit == 0 {
		paramsInfo.Limit = 1000
	}
	if paramsInfo.PerPage == 0 {
		paramsInfo.PerPage = 1000
	}

	// 数据校验
	valid := validation.Validation{}
	if result, err = valid.Valid(paramsInfo); err != nil {
		return
	}
	if !result {
		strErrMsg := ""
		for _, errMsg := range valid.Errors {
			strErrMsg += errMsg.Key + ":" + errMsg.Message + "; "
		}
		err = errors.New(strErrMsg)
		return
	}

	result = true
	return
}
