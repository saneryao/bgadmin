package validators

import (
	"errors"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/saneryao/bgadmin/utils"
)

// LoginParams 登录信息等参数
type LoginParams struct {
	Username  string `form:"username" json:"username" valid:"Required;MinSize(4);MaxSize(32);AlphaDash"`
	Password  string `form:"password" json:"password" valid:"Required;MinSize(4);MaxSize(32);"`
	IDCaptcha string `form:"captcha_id" json:"captcha_id" valid:"Required;MinSize(8);MaxSize(32);AlphaNumeric"`
	Captcha   string `form:"captcha" json:"captcha" valid:"Required;Length(4);Numeric"`
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
// func (this *LoginParams) Valid(valid *validation.Validation) {
// }

// ParseLoginParams 解析表单数据中的登录信息，并校验数据有效性
func ParseLoginParams(paramsInfo *LoginParams, ctrl *beego.Controller) (result bool, err error) {
	// 格式转换
	if err = ctrl.ParseForm(paramsInfo); err != nil {
		return
	}
	// logs.Debug("%#v", paramsInfo)
	var bytesPwd []byte
	if bytesPwd, err = utils.RsaDecryptByDefaultPrikey(paramsInfo.Password); err != nil {
		return
	}
	paramsInfo.Password = string(bytesPwd)

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
