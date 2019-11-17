package validators

import (
	"fmt"
	"github.com/astaxie/beego"
)

// ParseUserRoleInfo 解析表单数据中的用户角色信息，并校验数据有效性
func ParseUserRoleInfo(userID *int64, roleID *int32, ctrl *beego.Controller) (result bool, err error) {
	// 获取数据
	*userID, err = ctrl.GetInt64("user_id")
	if err != nil {
		return
	}
	*roleID, err = ctrl.GetInt32("role_id")
	if err != nil {
		return
	}

	// 数据校验
	if *userID <= 0 {
		err = fmt.Errorf("Invalid user id: %d", *userID)
		return
	}
	if *roleID <= 0 {
		err = fmt.Errorf("Invalid role id: %d", *roleID)
		return
	}

	result = true
	return
}
