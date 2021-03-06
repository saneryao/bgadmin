package service

import (
	// "github.com/astaxie/beego/logs"
	"github.com/saneryao/bgadmin/models"
)

// SortMenusByRelation 通过菜单的继承关系把菜单重新进行排序以便页面显示
func SortMenusByRelation(menus []*models.Menu) (sorted []*models.Menu) {
	sorted = findMenusByParentID(menus, 0)
	if sorted == nil {
		sorted = make([]*models.Menu, 0)
	}
	return
}
func findMenusByParentID(menus []*models.Menu, parentID int) (sorted []*models.Menu) {
	for _, menu := range menus {
		if menu.ParentID == parentID {
			sorted = append(sorted, menu)
			sorted = append(sorted, findMenusByParentID(menus, menu.ID)...)
		}
	}
	return
}

// HideDisabledMenus 隐藏状态为已禁用的菜单（如果父菜单已禁用，则子菜单也禁用）
func HideDisabledMenus(all []*models.Menu) (visible []*models.Menu) {
	// logs.Debug("!!!!!!!!!!DisableChildrenMenus!!!!!!!!!!")
	visible = make([]*models.Menu, 0)

	// 菜单的父菜单和状态信息，方便建立map进行快速查询
	type menuParentState struct {
		ParentID int
		State    int
	}

	// 建立map方便快速查询
	mapping := make(map[int]menuParentState)
	for _, v := range all {
		mapping[v.ID] = menuParentState{ParentID: v.ParentID, State: v.State}
	}

	// 遍历菜单列表（如果父菜单已禁用，则子菜单也禁用）
	for k, v := range all {
		// logs.Debug("==========%#v", v)
		if v.State == 1 && v.ParentID != 0 {
			nParentID := v.ParentID
			for { // 往上遍历所有父节点，若有父节点状态不为1，则改变此菜单状态
				if nParentID == 0 {
					break
				}
				parent, exists := mapping[nParentID]
				if !exists {
					break
				}
				if parent.State != 1 {
					all[k].State = parent.State
					// logs.Debug("@@@@@@@@%#v", all[k])
				}
				nParentID = parent.ParentID
			}
		}
	}

	// 遍历菜单列表（只返回状态为1的菜单）
	for k, v := range all {
		// logs.Debug("==========%#v", v)
		if v.State == 1 {
			visible = append(visible, all[k])
		}
	}

	return
}
