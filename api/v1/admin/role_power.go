package admin

import (
	"github.com/astaxie/beego/orm"
	"github.com/saneryao/bgadmin/models"
	"github.com/saneryao/bgadmin/service"
	"github.com/saneryao/bgadmin/validators"
)

// 定义一个角色菜单映射关系的API（用于角色和权限映射关系的增删改查）
type RolePowerAPI struct {
	baseAPI
}

// Get 执行http请求GET方法（beego定义的接口，查询角色和权限映射关系列表）
func (api *RolePowerAPI) Get() {
	// 定义变量
	var menus []*models.Menu

	// 包装并处理返回结果
	defer func() {
		api.PackResultData(menus, nil)
	}()

	// 从URL获取并校验ID
	var nId int64
	if nId, api.Error = validators.ParseIDFromURL(&api.Controller, ":id"); api.Error != nil {
		return
	}

	// 查询菜单列表
	db := orm.NewOrm()
	qs := db.QueryTable(new(models.Menu))
	qs = qs.OrderBy("id")
	_, api.Error = qs.All(&menus)
	if api.Error != nil {
		return
	}
	menus = service.SortMenusByRelation(menus)
	menus = service.HideDisabledMenus(menus)

	// 查询映射关系
	mapping := make(map[int64]bool)
	var relation []*models.RoleMenu
	qs2 := db.QueryTable(new(models.RoleMenu))
	qs2 = qs2.Filter("role_id__exact", nId) // 只查询角色nId的映射关系
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
		for _, v := range relation {
			mapping[v.MenuID] = true
		}
	}

	// 遍历菜单，将不存在映射关系的菜单状态置为0（状态为1表示存在映射关系）
	for k, v := range menus {
		if _, exists := mapping[v.ID]; !exists {
			menus[k].State = 0
		}
	}
}
