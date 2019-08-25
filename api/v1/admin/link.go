package admin

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/utils"
	"github.com/saneryao/bgadmin/validators"
	"strings"
)

// LinkAPI 定义一个链接API（用于链接的增删改查）
type LinkAPI struct {
	baseAPI
}

// Get 执行http请求GET方法（beego定义的接口，查询链接列表）
func (api *LinkAPI) Get() {
	// 定义变量
	params := validators.FilterParams{}
	var total int64
	var links []*models.Link

	// 包装并处理返回结果
	defer func() {
		others := make(map[string]interface{})
		others["draw"] = params.Draw
		if api.Error == nil {
			others["recordsTotal"] = total
			others["recordsFiltered"] = total
		}
		api.PackResultData(links, others)
	}()

	// 获取并校验输入信息
	if _, api.Error = validators.ParseFilterParams(&params, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.Link))
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
		_, api.Error = qs.Limit(params.Limit, params.Offset).All(&links)
		if api.Error != nil {
			return
		}
	}
}

// Post 执行http请求POST方法（beego定义的接口，新建链接）
func (api *LinkAPI) Post() {
	// 包装并处理返回结果
	link := models.Link{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(link, others)
	}()

	// 获取并校验输入信息
	if _, api.Error = validators.ParseLinkInfo(&link, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	linkRead := models.Link{Name: link.Name}
	errRead := db.Read(&linkRead, "Name")
	if errRead == nil {
		api.Error = errors.New(api.Tr("link") + "[" + linkRead.Name + "]" + api.Tr("is_already_exists"))
		return
	}

	// 插入数据
	_, api.Error = db.Insert(&link)
}

// Put 执行htttp请求PUT方法（beego定义的接口，更新某个链接的所有信息）
func (api *LinkAPI) Put() {
	// 包装并处理返回结果
	link := models.Link{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(link, others)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}

	// 获取并校验输入信息
	if _, api.Error = validators.ParseLinkInfo(&link, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	// linkRead := models.Link{Id: nID}
	// if api.Error = db.Read(&linkRead, "Id"); api.Error != nil {
	// 	return
	// }
	// if link.Name != "" && !strings.EqualFold(linkRead.Name, link.Name) {
	// 	api.Error = errors.New(api.Tr("name_cannot_be_modified"))
	// 	return
	// }

	// 更新数据
	link.ID = nID
	_, api.Error = db.Update(&link)
}

// Patch 执行http请求PATCH方法（beego定义的接口，更新某个链接的一部分信息）
func (api *LinkAPI) Patch() {
	// 包装并处理返回结果
	linkRead := models.Link{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(linkRead, others)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}
	linkRead.ID = nID

	// 获取并校验输入信息
	link := models.Link{}
	if _, api.Error = validators.ParseLinkInfo(&link, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	// qs := db.QueryTable(new(models.Link))
	// if api.Error = qs.Filter("Id__exact", nID).One(&linkRead); api.Error != nil {
	// 	return
	// }
	// if link.Name != "" && !strings.EqualFold(linkRead.Name, link.Name) {
	// 	api.Error = errors.New(api.Tr("name_cannot_be_modified"))
	// 	return
	// }

	// 合并数据
	if api.Error = utils.MergeData(&linkRead, link); api.Error != nil {
		return
	}

	// 更新数据
	_, api.Error = db.Update(&linkRead)
}

// Delete 执行http请求DELETE方法（beego定义的接口，删除某个链接）
func (api *LinkAPI) Delete() {
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
	_, api.Error = db.Delete(&models.Link{ID: nID})
}
