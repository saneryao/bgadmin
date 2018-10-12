package admin

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/utils"
	"github.com/saneryao/bgadmin/validators"
	"strings"
)

// 定义一个用户API（用于用户的增删改查）
type UsersApi struct {
	BaseApi
}

/* 功能：beego定义的接口，执行html请求GET方法，
 *             此处查询用户列表
 * 参数：空
 * 返回值：空
 */
func (this *UsersApi) Get() {
	// 定义变量
	params := validators.FilterParams{}
	var total int64
	var users []*models.User

	// 包装并处理返回结果
	defer func() {
		others := make(map[string]interface{})
		others["draw"] = params.Draw
		if this.Error == nil {
			others["recordsTotal"] = total
			others["recordsFiltered"] = total
		}
		this.PackResultData(users, others)
	}()

	// 获取并校验输入信息
	if _, this.Error = validators.ParseFilterParams(&params, &this.Controller); this.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.User))
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
		_, this.Error = qs.Limit(params.Limit, params.Offset).RelatedSel().All(&users)
	}
}

/* 功能：beego定义的接口，执行html请求POST方法，
 *             此处新建单个用户信息
 * 参数：空
 * 返回值：空
 */
func (this *UsersApi) Post() {
	// 包装并处理返回结果
	profile := models.Profile{}
	user := models.User{Profile: &profile}
	defer func() {
		var others map[string]interface{}
		if this.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = this.Tr("save") + this.Tr("success")
		}
		this.PackResultData(user, others)
	}()

	// 获取并校验输入信息
	if _, this.Error = validators.ParseUserInfo(&user, &this.Controller); this.Error != nil {
		return
	}
	user.Pwd = utils.EncryptPassword(user.Pwd)

	// 查询数据
	db := orm.NewOrm()
	userRead := models.User{Name: user.Name}
	errRead := db.Read(&userRead, "Name")
	if errRead == nil {
		this.Error = errors.New(this.Tr("user") + "[" + user.Name + "]" + this.Tr("is_already_exists"))
		return
	}

	// 插入数据
	_, this.Error = db.Insert(user.Profile)
	if this.Error != nil {
		return
	}
	_, this.Error = db.Insert(&user)
}

/* 功能：beego定义的接口，执行html请求PUT方法，
 *             此处更新单个用户的所有信息
 * 参数：空
 * 返回值：空
 */
func (this *UsersApi) Put() {
	// 包装并处理返回结果
	profile := models.Profile{}
	user := models.User{Profile: &profile}
	defer func() {
		var others map[string]interface{}
		if this.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = this.Tr("save") + this.Tr("success")
		}
		this.PackResultData(user, others)
	}()

	// 从URL获取并校验ID
	var nId int64
	if nId, this.Error = validators.ParseIdFromUrl(&this.Controller, ":id"); this.Error != nil {
		return
	}

	// 获取并校验输入信息
	if _, this.Error = validators.ParseUserInfo(&user, &this.Controller); this.Error != nil {
		return
	}
	user.Pwd = utils.EncryptPassword(user.Pwd)

	// 查询数据
	db := orm.NewOrm()
	userRead := models.User{Id: nId}
	if this.Error = db.Read(&userRead, "Id"); this.Error != nil {
		return
	}
	if user.Name != "" && !strings.EqualFold(userRead.Name, user.Name) {
		this.Error = errors.New(this.Tr("name_cannot_be_modified"))
		return
	}

	// 更新数据
	profile.Id = userRead.Profile.Id
	if _, this.Error = db.Update(&profile); this.Error != nil {
		return
	}
	user.Id = nId
	_, this.Error = db.Update(&user)
}

/* 功能：beego定义的接口，执行html请求PATCH方法，
 *             此处更新单个用户的某一部分信息
 * 参数：空
 * 返回值：空
 */
func (this *UsersApi) Patch() {
	// 包装并处理返回结果
	userRead := models.User{}
	defer func() {
		var others map[string]interface{}
		if this.Error == nil {
			others = make(map[string]interface{})
			others["msg"] = this.Tr("save") + this.Tr("success")
		}
		this.PackResultData(userRead, others)
	}()

	// 从URL获取并校验ID
	var nId int64
	if nId, this.Error = validators.ParseIdFromUrl(&this.Controller, ":id"); this.Error != nil {
		return
	}

	// 获取并校验输入信息
	profile := models.Profile{}
	user := models.User{Profile: &profile}
	if _, this.Error = validators.ParseUserInfo(&user, &this.Controller); this.Error != nil {
		return
	}

	// 查询数据
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.User))
	if this.Error = qs.Filter("Id__exact", nId).RelatedSel().One(&userRead); this.Error != nil {
		return
	}
	if user.Name != "" && !strings.EqualFold(userRead.Name, user.Name) {
		this.Error = errors.New(this.Tr("name_cannot_be_modified"))
		return
	}

	// 合并数据
	if this.Error = utils.MergeData(&userRead, user); this.Error != nil {
		return
	}

	// 更新数据
	if _, this.Error = db.Update(userRead.Profile); this.Error != nil {
		return
	}
	_, this.Error = db.Update(&userRead)
}

/* 功能：beego定义的接口，执行html请求DELETE方法，
 *             此处删除单个用户
 * 参数：空
 * 返回值：空
 */
func (this *UsersApi) Delete() {
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
	userRead := models.User{Id: nId}
	if this.Error = db.Read(&userRead, "Id"); this.Error != nil {
		return
	}

	// 删除数据
	if _, this.Error = db.Delete(&models.Profile{Id: userRead.Profile.Id}); this.Error != nil {
		return
	}
	_, this.Error = db.Delete(&models.User{Id: nId})
}
