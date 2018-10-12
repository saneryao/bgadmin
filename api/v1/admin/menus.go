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

// 定义一个菜单API（用于菜单的增删改查）
type MenusApi struct {
	BaseApi
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处查询菜单列表
 * 参数：空
 * 返回值：空
 */
func (this *MenusApi) Get() {
	// 定义变量
	params := validators.FilterParams{}
	var total int64
	var menus []*models.Menu

	// 包装并处理返回结果
	defer func() {
		others := make(map[string]interface{})
		others["draw"] = params.Draw
		if this.Error == nil {
			others["recordsTotal"] = total
			others["recordsFiltered"] = total
		}
		this.PackResultData(menus, others)
	}()

	// 获取并校验输入信息
	if _, this.Error = validators.ParseFilterParams(&params, &this.Controller); this.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.Menu))
	if params.SearchColumn != "" && params.SearchValue != "" {
		qs = qs.Filter(params.SearchColumn+"__contains", params.SearchValue)
	}
	total, this.Error = qs.Count() // 查询总数
	if this.Error != nil {
		return
	}
	if total > 0 { // 查询数据表分页信息
		qs = qs.Limit(params.Limit).Offset(params.Offset)
		if params.SortBy != "" {
			if strings.EqualFold(params.Order, "DESC") {
				qs = qs.OrderBy("-" + params.SortBy)
			} else {
				qs = qs.OrderBy(params.SortBy)
			}
		}
		_, this.Error = qs.Limit(params.Limit, params.Offset).All(&menus)
		if this.Error != nil {
			return
		}
		menus = service.SortMenusByRelation(menus)
	}
}

/* 功能：beego定义的接口，执行html请求POST方法，
 *             此处新建单个菜单信息
 * 参数：空
 * 返回值：空
 */
func (this *MenusApi) Post() {
	// 包装并处理返回结果
	menu := models.Menu{}
	defer func() {
		var others map[string]interface{}
		if this.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = this.Tr("save") + this.Tr("success")
		}
		this.PackResultData(menu, others)
	}()

	// 获取并校验输入信息
	if _, this.Error = validators.ParseMenuInfo(&menu, &this.Controller); this.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	menuRead := models.Menu{Name: menu.Name}
	errRead := db.Read(&menuRead, "Name")
	if errRead == nil {
		this.Error = errors.New(this.Tr("menu") + "[" + menuRead.Name + "]" + this.Tr("is_already_exists"))
		return
	}

	// 插入数据
	_, this.Error = db.Insert(&menu)
}

/* 功能：beego定义的接口，执行html请求PUT方法，
 *             此处更新单个菜单的所有信息
 * 参数：空
 * 返回值：空
 */
func (this *MenusApi) Put() {
	// 包装并处理返回结果
	menu := models.Menu{}
	defer func() {
		var others map[string]interface{}
		if this.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = this.Tr("save") + this.Tr("success")
		}
		this.PackResultData(menu, others)
	}()

	// 从URL获取并校验ID
	var nId int64
	if nId, this.Error = validators.ParseIdFromUrl(&this.Controller, ":id"); this.Error != nil {
		return
	}

	// 获取并校验输入信息
	if _, this.Error = validators.ParseMenuInfo(&menu, &this.Controller); this.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	// menuRead := models.Menu{Id: nId}
	// if this.Error = db.Read(&menuRead, "Id"); this.Error != nil {
	// 	return
	// }
	// if menu.Name != "" && !strings.EqualFold(menuRead.Name, menu.Name) {
	// 	this.Error = errors.New(this.Tr("name_cannot_be_modified"))
	// 	return
	// }

	// 更新数据
	menu.Id = nId
	_, this.Error = db.Update(&menu)
}

/* 功能：beego定义的接口，执行html请求PATCH方法，
 *             此处更新单个菜单的某一部分信息
 * 参数：空
 * 返回值：空
 */
func (this *MenusApi) Patch() {
	// 包装并处理返回结果
	menuRead := models.Menu{}
	defer func() {
		var others map[string]interface{}
		if this.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = this.Tr("save") + this.Tr("success")
		}
		this.PackResultData(menuRead, others)
	}()

	// 从URL获取并校验ID
	var nId int64
	if nId, this.Error = validators.ParseIdFromUrl(&this.Controller, ":id"); this.Error != nil {
		return
	}
	menuRead.Id = nId

	// 获取并校验输入信息
	menu := models.Menu{}
	if _, this.Error = validators.ParseMenuInfo(&menu, &this.Controller); this.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	// qs := db.QueryTable(new(models.Menu))
	// if this.Error = qs.Filter("Id__exact", nId).One(&menuRead); this.Error != nil {
	// 	return
	// }
	// if menu.Name != "" && !strings.EqualFold(menuRead.Name, menu.Name) {
	// 	this.Error = errors.New(this.Tr("name_cannot_be_modified"))
	// 	return
	// }

	// 合并数据
	if this.Error = utils.MergeData(&menuRead, menu); this.Error != nil {
		return
	}

	// 更新数据
	_, this.Error = db.Update(&menuRead)
}

/* 功能：beego定义的接口，执行html请求DELETE方法，
 *             此处删除单个菜单
 * 参数：空
 * 返回值：空
 */
func (this *MenusApi) Delete() {
	// 包装并处理返回结果
	defer func() {
		var others map[string]interface{}
		if this.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = this.Tr("del") + this.Tr("success")
		}
		this.PackResultData(nil, others)
	}()

	// 从URL获取并校验ID
	var nId int64
	if nId, this.Error = validators.ParseIdFromUrl(&this.Controller, ":id"); this.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	menuRead := models.Menu{ParentId: nId}
	if this.Error = db.Read(&menuRead, "ParentId"); this.Error == nil {
		this.Error = errors.New(this.Tr("contain_submenus"))
		return
	}

	// 删除数据
	_, this.Error = db.Delete(&models.Menu{Id: nId})
}
