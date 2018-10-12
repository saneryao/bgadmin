package admin

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/utils"
	"github.com/saneryao/bgadmin/validators"
	"strings"
)

// 定义一个角色API（用于角色的增删改查）
type RolesApi struct {
	BaseApi
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处查询角色列表
 * 参数：空
 * 返回值：空
 */
func (this *RolesApi) Get() {
	// 定义变量
	params := validators.FilterParams{}
	var total int64
	var roles []*models.Role

	// 包装并处理返回结果
	defer func() {
		others := make(map[string]interface{})
		others["draw"] = params.Draw
		if this.Error == nil {
			others["recordsTotal"] = total
			others["recordsFiltered"] = total
		}
		this.PackResultData(roles, others)
	}()

	// 获取并校验输入信息
	if _, this.Error = validators.ParseFilterParams(&params, &this.Controller); this.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.Role))
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
		_, this.Error = qs.Limit(params.Limit, params.Offset).All(&roles)
	}
}

/* 功能：beego定义的接口，执行html请求POST方法，
 *             此处新建单个角色信息
 * 参数：空
 * 返回值：空
 */
func (this *RolesApi) Post() {
	// 包装并处理返回结果
	role := models.Role{}
	defer func() {
		var others map[string]interface{}
		if this.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = this.Tr("save") + this.Tr("success")
		}
		this.PackResultData(role, others)
	}()

	// 获取并校验输入信息
	if _, this.Error = validators.ParseRoleInfo(&role, &this.Controller); this.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	roleRead := models.Role{Name: role.Name}
	errRead := db.Read(&roleRead, "Name")
	if errRead == nil {
		this.Error = errors.New(this.Tr("role") + "[" + roleRead.Name + "]" + this.Tr("is_already_exists"))
		return
	}

	// 插入数据
	_, this.Error = db.Insert(&role)
}

/* 功能：beego定义的接口，执行html请求PUT方法，
 *             此处更新单个角色的所有信息
 * 参数：空
 * 返回值：空
 */
func (this *RolesApi) Put() {
	// 包装并处理返回结果
	role := models.Role{}
	defer func() {
		var others map[string]interface{}
		if this.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = this.Tr("save") + this.Tr("success")
		}
		this.PackResultData(role, others)
	}()

	// 从URL获取并校验ID
	var nId int64
	if nId, this.Error = validators.ParseIdFromUrl(&this.Controller, ":id"); this.Error != nil {
		return
	}

	// 获取并校验输入信息
	if _, this.Error = validators.ParseRoleInfo(&role, &this.Controller); this.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	roleRead := models.Role{Id: nId}
	if this.Error = db.Read(&roleRead, "Id"); this.Error != nil {
		return
	}
	if role.Name != "" && !strings.EqualFold(roleRead.Name, role.Name) {
		this.Error = errors.New(this.Tr("name_cannot_be_modified"))
		return
	}

	// 更新数据
	role.Id = nId
	_, this.Error = db.Update(&role)
}

/* 功能：beego定义的接口，执行html请求PATCH方法，
 *             此处更新单个角色的某一部分信息
 * 参数：空
 * 返回值：空
 */
func (this *RolesApi) Patch() {
	// 包装并处理返回结果
	roleRead := models.Role{}
	defer func() {
		var others map[string]interface{}
		if this.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = this.Tr("save") + this.Tr("success")
		}
		this.PackResultData(roleRead, others)
	}()

	// 从URL获取并校验ID
	var nId int64
	if nId, this.Error = validators.ParseIdFromUrl(&this.Controller, ":id"); this.Error != nil {
		return
	}
	roleRead.Id = nId

	// 获取并校验输入信息
	role := models.Role{}
	if _, this.Error = validators.ParseRoleInfo(&role, &this.Controller); this.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.Role))
	if this.Error = qs.Filter("Id__exact", nId).One(&roleRead); this.Error != nil {
		return
	}
	if role.Name != "" && !strings.EqualFold(roleRead.Name, role.Name) {
		this.Error = errors.New(this.Tr("name_cannot_be_modified"))
		return
	}

	// 合并数据
	if this.Error = utils.MergeData(&roleRead, role); this.Error != nil {
		return
	}

	// 更新数据
	_, this.Error = db.Update(&roleRead)
}

/* 功能：beego定义的接口，执行html请求DELETE方法，
 *             此处删除单个角色
 * 参数：空
 * 返回值：空
 */
func (this *RolesApi) Delete() {
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

	// 删除数据
	db := orm.NewOrm()
	_, this.Error = db.Delete(&models.Role{Id: nId})
}
