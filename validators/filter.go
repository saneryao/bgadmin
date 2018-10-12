package validators

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

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

/* 功能：Form表单数据转化成结构体，并校验数据有效性
 * 参数：paramsTable是转换后的结果（结构体），ctrl是控制器
 * 返回值：result为true表示转换成功，err不为nil表示转换失败
 */
func ParseFilterParams(paramsInfo *FilterParams, ctrl *beego.Controller) (result bool, err error) {
	// 格式转换
	if err = ctrl.ParseForm(paramsInfo); err != nil {
		result = false
		return
	}
	logs.Debug("%#v", paramsInfo)

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
