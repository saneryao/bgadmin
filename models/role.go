package models

type Role struct {
	Id    int64  `form:"id" json:"id"`
	Name  string `form:"name" json:"name" valid:"MinSize(4);MaxSize(32);AlphaDash"`
	Desc  string `form:"desc" json:"desc" valid:"MaxSize(255)"`
	State int    `form:"state" json:"state" valid:"Range(0,1)"`
}

// func (role *Role) TableName() string {
// 	return "role"
// }
