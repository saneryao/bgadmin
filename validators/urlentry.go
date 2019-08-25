package validators

import (
	"errors"
	"github.com/astaxie/beego"
)

// ParseEntryFromURL 从访问的URL中获取Entry，并校验其有效性
func ParseEntryFromURL(ctrl *beego.Controller, name string) (entry string, err error) {
	entry = ctrl.Ctx.Input.Param(name)
	if entry == "" {
		err = errors.New("The field `entry` is not found int request URL")
		return
	}
	return
}
