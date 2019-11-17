package admin

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/validators"
)

// UserRoleAPI 定义一个角色菜单映射关系的API（用于用户和角色映射关系的增删改查）
type UserRoleAPI struct {
	baseAPI
}

// Get 执行http请求GET方法（查询用户和角色映射关系列表）
func (api *UserRoleAPI) Get() {
	// 定义变量
	var totalRoles int64
	var user models.User

	// 包装并处理返回结果
	defer func() {
		others := make(map[string]interface{})
		if api.Error == nil {
			others["totalRoles"] = totalRoles
		}
		api.PackResultData(user, others)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}
	user.ID = nID

	// 查询数据
	db := orm.NewOrm()
	if totalRoles, api.Error = db.LoadRelated(&user, "Roles"); api.Error != nil {
		return
	}
}

// Post 执行http请求POST方法（新建用户角色）
func (api *UserRoleAPI) Post() {
	// 包装并处理返回结果
	userRole := models.UserRole{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(userRole, others)
	}()

	// 获取并校验输入信息
	var userID int64
	var roleID int32
	if _, api.Error = validators.ParseUserRoleInfo(&userID, &roleID, &api.Controller); api.Error != nil {
		return
	}
	userRole.User = &models.User{ID: userID}
	userRole.Role = &models.Role{ID: int(roleID)}

	// 查询数据
	db := orm.NewOrm()
	if api.Error = db.Read(&userRole, "user_id", "role_id"); api.Error == nil {
		roleInfo := fmt.Sprintf("[%d-%d]", userID, roleID)
		api.Error = errors.New(api.Tr("user_role") + roleInfo + api.Tr("is_already_exists"))
		return
	}
	if api.Error == orm.ErrNoRows {
		api.Error = nil
	}
	if api.Error != nil {
		return
	}

	// 插入数据
	_, api.Error = db.Insert(&userRole)
}

// Delete 执行http请求DELETE方法（删除某个用户的角色）
func (api *UserRoleAPI) Delete() {
	// 包装并处理返回结果
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("del") + api.Tr("success")
		}
		api.PackResultData(nil, others)
	}()

	var userID int64
	var roleID int32
	if _, api.Error = validators.ParseUserRoleInfo(&userID, &roleID, &api.Controller); api.Error != nil {
		return
	}
	userRole := models.UserRole{}
	userRole.User = &models.User{ID: userID}
	userRole.Role = &models.Role{ID: int(roleID)}

	// 查询数据
	db := orm.NewOrm()
	if api.Error = db.Read(&userRole, "user_id", "role_id"); api.Error != nil {
		return
	}

	// 删除数据
	_, api.Error = db.Delete(&userRole)
}
