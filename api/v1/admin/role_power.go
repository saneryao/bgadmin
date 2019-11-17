package admin

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/validators"
)

// 定义一个角色和菜单或链接映射关系的API（用于角色和权限映射关系的增删改查）
type RolePowerAPI struct {
	baseAPI
}

// Get 执行http请求GET方法（查询角色和权限映射关系列表）
func (api *RolePowerAPI) Get() {
	// 定义变量
	var totalMenus int64
	var totalLinks int64
	var role models.Role

	// 包装并处理返回结果
	defer func() {
		others := make(map[string]interface{})
		if api.Error == nil {
			others["totalMenus"] = totalMenus
			others["totalLinks"] = totalLinks
		}
		api.PackResultData(role, others)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}
	role.ID = int(nID)

	// 查询数据
	db := orm.NewOrm()
	if totalMenus, api.Error = db.LoadRelated(&role, "Menus"); api.Error != nil {
		return
	}
	if totalLinks, api.Error = db.LoadRelated(&role, "Links"); api.Error != nil {
		return
	}
}

// Post 执行http请求POST方法（新建角色权限）
func (api *RolePowerAPI) Post() {
	// 包装并处理返回结果
	rolePower := models.RolePower{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(rolePower, others)
	}()

	// 获取并校验输入信息
	var roleID, menuID, linkID int32
	if _, api.Error = validators.ParseRolePowerInfo(&roleID, &menuID, &linkID, &api.Controller); api.Error != nil {
		return
	}
	rolePower.Role = &models.Role{ID: int(roleID)}
	rolePower.Menu = &models.Menu{ID: int(menuID)}
	rolePower.Link = &models.Link{ID: int(linkID)}

	// 查询数据
	db := orm.NewOrm()
	if api.Error = db.Read(&rolePower, "role_id", "menu_id", "link_id"); api.Error == nil {
		roleInfo := fmt.Sprintf("[%d-%d-%d]", roleID, menuID, linkID)
		api.Error = errors.New(api.Tr("role_power") + roleInfo + api.Tr("is_already_exists"))
		return
	}
	if api.Error == orm.ErrNoRows {
		api.Error = nil
	}
	if api.Error != nil {
		return
	}

	// 插入数据
	_, api.Error = db.Insert(&rolePower)
}

// Delete 执行http请求DELETE方法（删除某个角色的权限）
func (api *RolePowerAPI) Delete() {
	// 包装并处理返回结果
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("del") + api.Tr("success")
		}
		api.PackResultData(nil, others)
	}()

	// 获取并校验输入信息
	var roleID, menuID, linkID int32
	if _, api.Error = validators.ParseRolePowerInfo(&roleID, &menuID, &linkID, &api.Controller); api.Error != nil {
		return
	}
	rolePower := models.RolePower{}
	rolePower.Role = &models.Role{ID: int(roleID)}
	rolePower.Menu = &models.Menu{ID: int(menuID)}
	rolePower.Link = &models.Link{ID: int(linkID)}

	// 查询数据
	db := orm.NewOrm()
	if api.Error = db.Read(&rolePower, "role_id", "menu_id", "link_id"); api.Error != nil {
		return
	}

	// 删除数据
	_, api.Error = db.Delete(&rolePower)
}
