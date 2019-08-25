package admin

import (
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/validators"
)

// UserRoleAPI 定义一个角色菜单映射关系的API（用于用户和角色映射关系的增删改查）
type UserRoleAPI struct {
	baseAPI
}

// Get 执行http请求GET方法（beego定义的接口，查询用户和角色映射关系列表）
func (api *UserRoleAPI) Get() {
	// 定义变量
	var roles []*models.Role

	// 包装并处理返回结果
	defer func() {
		api.PackResultData(roles, nil)
	}()

	// 从URL获取并校验ID
	var nID int64
	if nID, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}

	// 查询菜单列表
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.Role))
	qs = qs.Filter("state__exact", 1) // 只查询状态为1的角色
	qs = qs.OrderBy("id")
	_, api.Error = qs.All(&roles)
	if api.Error != nil {
		return
	}

	// 查询映射关系
	mapping := make(map[int64]bool)
	var relation []*models.UserRole
	qs2 := db.QueryTable(new(models.UserRole))
	qs2 = qs2.Filter("user_id__exact", nID) // 只查询用户nId的映射关系
	var count int64
	count, api.Error = qs2.Count()
	if api.Error != nil {
		return
	}
	if count > 0 {
		_, api.Error = qs2.All(&relation)
		if api.Error != nil {
			return
		}
		mapping := make(map[int64]bool)
		for _, v := range relation {
			mapping[v.RoleID] = true
		}
	}

	// 遍历菜单，将不存在映射关系的菜单状态置为0（状态为1表示存在映射关系）
	for k, v := range roles {
		if _, exists := mapping[v.ID]; !exists {
			roles[k].State = 0
		}
	}
}
