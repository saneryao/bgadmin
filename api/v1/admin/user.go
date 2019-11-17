package admin

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/utils"
	"github.com/saneryao/bgadmin/validators"
	"strings"
)

// UserAPI 定义一个用户API（用于用户的增删改查）
type UserAPI struct {
	baseAPI
}

// Get 执行http请求GET方法（查询用户列表）
func (api *UserAPI) Get() {
	// 定义变量
	params := validators.FilterParams{}
	var total int64
	var users []*models.User

	// 包装并处理返回结果
	defer func() {
		others := make(map[string]interface{})
		others["draw"] = params.Draw
		if api.Error == nil {
			others["recordsTotal"] = total
			others["recordsFiltered"] = total
		}
		api.PackResultData(users, others)
	}()

	// 获取并校验输入信息
	if _, api.Error = validators.ParseFilterParams(&params, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.User))
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
		_, api.Error = qs.Limit(params.Limit, params.Offset).RelatedSel().All(&users)
	}
}

// Post 执行http请求POST方法（新建用户）
func (api *UserAPI) Post() {
	// 包装并处理返回结果
	profile := models.Profile{}
	user := models.User{Profile: &profile}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(user, others)
	}()

	// 获取并校验输入信息
	if _, api.Error = validators.ParseUserInfo(&user, &api.Controller); api.Error != nil {
		return
	}
	user.Pwd = utils.EncryptPassword(user.Pwd)

	// 查询数据
	db := orm.NewOrm()
	userRead := models.User{Name: user.Name}
	errRead := db.Read(&userRead, "Name")
	if errRead == nil {
		api.Error = errors.New(api.Tr("user") + "[" + user.Name + "]" + api.Tr("is_already_exists"))
		return
	}

	// 插入数据
	_, api.Error = db.Insert(user.Profile)
	if api.Error != nil {
		return
	}
	_, api.Error = db.Insert(&user)
}

// Put 执行http请求PUT方法（更新某个用户的所有信息）
func (api *UserAPI) Put() {
	// 包装并处理返回结果
	profile := models.Profile{}
	user := models.User{Profile: &profile}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(user, others)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}

	// 获取并校验输入信息
	if _, api.Error = validators.ParseUserInfo(&user, &api.Controller); api.Error != nil {
		return
	}
	user.Pwd = utils.EncryptPassword(user.Pwd)

	// 查询数据
	db := orm.NewOrm()
	userRead := models.User{ID: nID}
	if api.Error = db.Read(&userRead, "ID"); api.Error != nil {
		return
	}
	if user.Name != "" && !strings.EqualFold(userRead.Name, user.Name) {
		api.Error = errors.New(api.Tr("name_cannot_be_modified"))
		return
	}

	// 更新数据
	profile.ID = userRead.Profile.ID
	if _, api.Error = db.Update(&profile); api.Error != nil {
		return
	}
	user.ID = nID
	_, api.Error = db.Update(&user)
}

// Patch 执行http请求PATCH方法（更新某个用户的一部分信息）
func (api *UserAPI) Patch() {
	// 包装并处理返回结果
	userRead := models.User{}
	defer func() {
		var others map[string]interface{}
		if api.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = api.Tr("save") + api.Tr("success")
		}
		api.PackResultData(userRead, others)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}

	// 获取并校验输入信息
	profile := models.Profile{}
	user := models.User{Profile: &profile}
	if _, api.Error = validators.ParseUserInfo(&user, &api.Controller); api.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.User))
	if api.Error = qs.Filter("ID__exact", nID).RelatedSel().One(&userRead); api.Error != nil {
		return
	}
	if user.Name != "" && !strings.EqualFold(userRead.Name, user.Name) {
		api.Error = errors.New(api.Tr("name_cannot_be_modified"))
		return
	}

	// 合并数据
	if api.Error = utils.MergeData(&userRead, user); api.Error != nil {
		return
	}

	// 更新数据
	if _, api.Error = db.Update(userRead.Profile); api.Error != nil {
		return
	}
	_, api.Error = db.Update(&userRead)
}

// Delete 执行http请求DELETE方法（删除某个用户）
func (api *UserAPI) Delete() {
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
	userRead := models.User{ID: nID}
	if api.Error = db.Read(&userRead, "ID"); api.Error != nil {
		return
	}

	// 删除数据
	if _, api.Error = db.Delete(&models.Profile{ID: userRead.Profile.ID}); api.Error != nil {
		return
	}
	_, api.Error = db.Delete(&models.User{ID: nID})
}
