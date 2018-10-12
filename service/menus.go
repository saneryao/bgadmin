package service

import (
	// "github.com/astaxie/beego/logs"
	"github.com/saneryao/bgadmin/models"
)

/* 功能：通过菜单的继承关系把菜单重新进行排序
 * 参数：menus（要排序的菜单列表）
 * 返回值：sorted（排序后的菜单列表）
 */
func SortMenusByRelation(menus []*models.Menu) (sorted []*models.Menu) {
	sorted = make([]*models.Menu, 0)
	for k1, _ := range menus { // 遍历菜单列表
		if menus[k1].ParentId == 0 { // 判断是否是一级菜单
			sorted = append(sorted, menus[k1])
			// logs.Debug("**********%#v", menus[k1])
		} else { // 子菜单
			var findChildren bool = false
			var indexPos int = 0
			for k2, _ := range sorted {
				if findChildren {
					if sorted[k2].ParentId == menus[k1].ParentId {
						continue
					}
					if sorted[k2].Id != menus[k1].ParentId {
						indexPos = k2
						break
					}
				} else {
					if sorted[k2].Id == menus[k1].ParentId {
						findChildren = true
					}
				}
			}
			// logs.Debug("++++++++++%#v", indexPos)
			if indexPos == 0 {
				sorted = append(sorted, menus[k1])
			} else {
				tail := append([]*models.Menu{}, sorted[indexPos:]...)
				sorted = append(sorted[0:indexPos], menus[k1])
				sorted = append(sorted, tail...)
			}
			// logs.Debug("**********%#v", menus[k1])
		}
	}
	return
}
