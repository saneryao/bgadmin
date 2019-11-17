package validators

import (
	"fmt"
	"github.com/astaxie/beego"
)

// ParseRolePowerInfo 解析表单数据中的角色权限信息，并校验数据有效性
func ParseRolePowerInfo(roleID, menuID, linkID *int32, ctrl *beego.Controller) (result bool, err error) {
	// 获取数据
	*roleID, err = ctrl.GetInt32("role_id")
	if err != nil {
		return
	}
	*menuID, err = ctrl.GetInt32("menu_id")
	if err != nil {
		return
	}
	*linkID, err = ctrl.GetInt32("link_id")
	if err != nil {
		return
	}

	// 数据校验
	if *roleID <= 0 {
		err = fmt.Errorf("Invalid role id: %d", *roleID)
		return
	}
	if *menuID < 0 {
		err = fmt.Errorf("Invalid menu id: %d", *menuID)
		return
	}
	if *linkID < 0 {
		err = fmt.Errorf("Invalid link id: %d", *linkID)
		return
	}
	if *menuID == 0 && *linkID == 0 {
		err = fmt.Errorf("Invalid param menu id or link id: %d, %d", *menuID, *linkID)
		return
	}

	result = true
	return
}
