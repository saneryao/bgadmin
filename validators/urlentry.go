package validators

import (
	"errors"
	"github.com/astaxie/beego"
)

/* 功能：从访问的URL中获取Entry，并校验其有效性
 * 参数：ctrl是控制器
 * 返回值：id表示转换后的Entry，err不为nil表示转换失败
 */
func ParseEntryFromUrl(ctrl *beego.Controller, name string) (entry string, err error) {
	entry = ctrl.Ctx.Input.Param(name)
	if entry == "" {
		err = errors.New("The field `entry` is not found int request URL")
		return
	}
	return
}
