package models

// RoleMenu 定义角色和菜单映射关系（一个角色可以看到的菜单）
type RoleMenu struct {
	ID     int64
	RoleID int64
	MenuID int64
}

// func (rm *RoleMenu) TableName() string {
// 	return "role_menu"
// }
