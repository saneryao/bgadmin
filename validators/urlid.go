package validators

import (
	"errors"
	"github.com/astaxie/beego"
	"strconv"
)

/* 功能：从访问的URL中获取ID，并校验其有效性
 * 参数：ctrl是控制器
 * 返回值：id表示转换后的ID，err不为nil表示转换失败
 */
func ParseIdFromUrl(ctrl *beego.Controller, name string) (id int64, err error) {
	strId := ctrl.Ctx.Input.Param(name)
	if strId == "" {
		err = errors.New("The field `ID` is not found int request URL")
		return
	}

	id, err = strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return
	}

	return
}
