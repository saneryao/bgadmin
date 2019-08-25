package models

// Role 定义角色信息
type Role struct {
	ID    int64  `form:"id" json:"id"`
	Name  string `form:"name" json:"name" valid:"MaxSize(32);AlphaDash"`
	Desc  string `form:"desc" json:"desc" valid:"MaxSize(255)"`
	State int    `form:"state" json:"state" valid:"Range(0,1)"`
}

// func (role *Role) TableName() string {
// 	return "role"
// }
