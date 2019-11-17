package models

import "time"

// Role 定义角色信息
type Role struct {
	ID         int       `form:"id" json:"id" orm:"column(id);auto"`
	Name       string    `form:"name" json:"name" orm:"column(name)" valid:"MaxSize(32);AlphaDash"`
	Desc       string    `form:"desc" json:"desc" orm:"column(desc)" valid:"MaxSize(255)"`
	State      int       `form:"state" json:"state" orm:"column(state)" valid:"Range(0,1)"`
	CreateTime time.Time `form:"create_time" json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime time.Time `form:"update_time" json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
	Menus      []*Menu   `orm:"rel(m2m);rel_through(github.com/saneryao/bgadmin/models.RolePower)"`
	Links      []*Link   `orm:"rel(m2m);rel_through(github.com/saneryao/bgadmin/models.RolePower)"`
}

// func (role *Role) TableName() string {
// 	return "role"
// }
