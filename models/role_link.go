package models

// RoleLink 定义角色和链接映射关系（一个角色可以访问哪些链接）
type RoleLink struct {
	ID     int64
	RoleID int64
	LinkID int64
}

// func (rp *RoleLink) TableName() string {
// 	return "role_link"
// }
