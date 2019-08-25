package admin

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/service"
	"github.com/saneryao/bgadmin/utils"
	"github.com/saneryao/bgadmin/validators"
	"strings"
)

// MenuAPI 定义一个菜单API（用于菜单的增删改查）
type MenuAPI struct {
	baseAPI
}

// Get 执行http请求GET方法（beego定义的接口，查询菜单列表）
func (api *MenuAPI) Get() {
	// 定义变量
	params := validators.FilterParams{}
	var total int64
	var menus []*models.Menu

	// 包装并处理返回结果
	defer func() {
		others := make(map[string]interface{})
		others["draw"] = params.Draw
		if api.Error == nil {
			others["recordsTotal"] = total
			others["recordsFiltered"] = total
		}
		api.PackResultData(menus, others)
	}()

	// 获取并校验输入信息
	if _, api.Error = validators.ParseFilterParams(&params, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.Menu))
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
		_, api.Error = qs.Limit(params.Limit, params.Offset).All(&menus)
		if api.Error != nil {
			return
		}
		menus = service.SortMenusByRelation(menus)
	}
}

// Post 执行http请求POST方法（beego定义的接口，新建菜单）
func (api *MenuAPI) Post() {
	// 包装并处理返回结果
	menu := models.Menu{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(menu, others)
	}()

	// 获取并校验输入信息
	if _, api.Error = validators.ParseMenuInfo(&menu, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	menuRead := models.Menu{Name: menu.Name}
	errRead := db.Read(&menuRead, "Name")
	if errRead == nil {
		api.Error = errors.New(api.Tr("menu") + "[" + menuRead.Name + "]" + api.Tr("is_already_exists"))
		return
	}

	// 插入数据
	_, api.Error = db.Insert(&menu)
}

// Put 执行http请求PUT方法（beego定义的接口，更新某个菜单的所有信息）
func (api *MenuAPI) Put() {
	// 包装并处理返回结果
	menu := models.Menu{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(menu, others)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}

	// 获取并校验输入信息
	if _, api.Error = validators.ParseMenuInfo(&menu, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	// menuRead := models.Menu{Id: nID}
	// if api.Error = db.Read(&menuRead, "Id"); api.Error != nil {
	// 	return
	// }
	// if menu.Name != "" && !strings.EqualFold(menuRead.Name, menu.Name) {
	// 	api.Error = errors.New(api.Tr("name_cannot_be_modified"))
	// 	return
	// }

	// 更新数据
	menu.ID = nID
	_, api.Error = db.Update(&menu)
}

// Patch 执行http请求PATCH方法（beego定义的接口，更新某个菜单的一部分信息）
func (api *MenuAPI) Patch() {
	// 包装并处理返回结果
	menuRead := models.Menu{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(menuRead, others)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}
	menuRead.ID = nID

	// 获取并校验输入信息
	menu := models.Menu{}
	if _, api.Error = validators.ParseMenuInfo(&menu, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	// qs := db.QueryTable(new(models.Menu))
	// if api.Error = qs.Filter("Id__exact", nID).One(&menuRead); api.Error != nil {
	// 	return
	// }
	// if menu.Name != "" && !strings.EqualFold(menuRead.Name, menu.Name) {
	// 	api.Error = errors.New(api.Tr("name_cannot_be_modified"))
	// 	return
	// }

	// 合并数据
	if api.Error = utils.MergeData(&menuRead, menu); api.Error != nil {
		return
	}

	// 更新数据
	_, api.Error = db.Update(&menuRead)
}

// Delete 执行http请求DELETE方法（beego定义的接口，删除某个菜单）
func (api *MenuAPI) Delete() {
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

	// 查询数据
	db := orm.NewOrm()
	menuRead := models.Menu{ParentID: nID}
	if api.Error = db.Read(&menuRead, "ParentId"); api.Error == nil {
		api.Error = errors.New(api.Tr("contain_submenus"))
		return
	}

	// 删除数据
	_, api.Error = db.Delete(&models.Menu{ID: nID})
}
