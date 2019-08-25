package validators

import (
	"errors"
	"github.com/astaxie/beego"
	"strconv"
)

// ParseIDFromURL 从访问的URL中获取ID，并校验其有效性
func ParseIDFromURL(ctrl *beego.Controller, name string) (id int64, err error) {
	strID := ctrl.Ctx.Input.Param(name)
	if strID == "" {
		err = errors.New("The field `ID` is not found int request URL")
		return
	}

	id, err = strconv.ParseInt(strID, 10, 64)
	if err != nil {
		return
	}

	return
}
