package admin

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/utils"
	"github.com/saneryao/bgadmin/validators"
	"strings"
)

// RoleAPI 定义一个角色API（用于角色的增删改查）
type RoleAPI struct {
	baseAPI
}

// Get 执行http请求GET方法（查询角色列表）
func (api *RoleAPI) Get() {
	// 定义变量
	params := validators.FilterParams{}
	var total int64
	var roles []*models.Role

	// 包装并处理返回结果
	defer func() {
		others := make(map[string]interface{})
		others["draw"] = params.Draw
		if api.Error == nil {
			others["recordsTotal"] = total
			others["recordsFiltered"] = total
		}
		api.PackResultData(roles, others)
	}()

	// 获取并校验输入信息
	if _, api.Error = validators.ParseFilterParams(&params, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.Role))
	if params.SearchColumn != "" && params.SearchValue != "" {
		qs = qs.Filter(params.SearchColumn+"__contains", params.SearchValue)
	}
	total, api.Error = qs.Count() // 查询总数
	if api.Error != nil {
		return
	}
	if total > 0 { // 查询数据表分页信息
		if params.SortBy != "" {
			if strings.EqualFold(params.Order, "DESC") {
				qs = qs.OrderBy("-" + params.SortBy)
			} else {
				qs = qs.OrderBy(params.SortBy)
			}
		}
		_, api.Error = qs.Limit(params.Limit, params.Offset).All(&roles)
	}
}

// Post 执行http请求POST方法（新建角色）
func (api *RoleAPI) Post() {
	// 包装并处理返回结果
	role := models.Role{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(role, others)
	}()

	// 获取并校验输入信息
	if _, api.Error = validators.ParseRoleInfo(&role, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	roleRead := models.Role{Name: role.Name}
	errRead := db.Read(&roleRead, "Name")
	if errRead == nil {
		api.Error = errors.New(api.Tr("role") + "[" + roleRead.Name + "]" + api.Tr("is_already_exists"))
		return
	}

	// 插入数据
	_, api.Error = db.Insert(&role)
}

// Put 执行http请求PUT方法（更新某个角色的所有信息）
func (api *RoleAPI) Put() {
	// 包装并处理返回结果
	role := models.Role{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(role, others)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}

	// 获取并校验输入信息
	if _, api.Error = validators.ParseRoleInfo(&role, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	roleRead := models.Role{ID: int(nID)}
	if api.Error = db.Read(&roleRead, "ID"); api.Error != nil {
		return
	}
	if role.Name != "" && !strings.EqualFold(roleRead.Name, role.Name) {
		api.Error = errors.New(api.Tr("name_cannot_be_modified"))
		return
	}

	// 更新数据
	role.ID = int(nID)
	_, api.Error = db.Update(&role)
}

// Patch 执行http请求PATCH方法（更新某个角色的一部分信息）
func (api *RoleAPI) Patch() {
	// 包装并处理返回结果
	roleRead := models.Role{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(roleRead, others)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}
	roleRead.ID = int(nID)

	// 获取并校验输入信息
	role := models.Role{}
	if _, api.Error = validators.ParseRoleInfo(&role, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.Role))
	if api.Error = qs.Filter("ID__exact", nID).One(&roleRead); api.Error != nil {
		return
	}
	if role.Name != "" && !strings.EqualFold(roleRead.Name, role.Name) {
		api.Error = errors.New(api.Tr("name_cannot_be_modified"))
		return
	}

	// 合并数据
	if api.Error = utils.MergeData(&roleRead, role); api.Error != nil {
		return
	}

	// 更新数据
	_, api.Error = db.Update(&roleRead)
}

// Delete 执行http请求DELETE方法（删除某个角色）
func (api *RoleAPI) Delete() {
	// 包装并处理返回结果
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("del") + api.Tr("success")
		}
		api.PackResultData(nil, others)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}

	// 删除数据
	db := orm.NewOrm()
	_, api.Error = db.Delete(&models.Role{ID: int(nID)})
}
