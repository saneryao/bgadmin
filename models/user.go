package models

import "time"

// import (
// 	"github.com/astaxie/beego/validation"
// 	"github.com/beego/i18n"
// 	"strings"
// )

// User 定义用户信息
type User struct {
	ID         int64     `form:"id" json:"id" orm:"column(id);auto"`
	Name       string    `form:"name" json:"name" orm:"column(name)" valid:"MaxSize(32);AlphaDash"`
	Pwd        string    `form:"pwd" json:"pwd" orm:"column(pwd)" valid:"MaxSize(32)"`
	State      int       `form:"state" json:"state" orm:"column(state)" valid:"Range(0,1)"`
	CreateTime time.Time `form:"create_time" json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime time.Time `form:"update_time" json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
	Profile    *Profile  `orm:"rel(one)"` // 一对一映射
	Roles      []*Role   `orm:"rel(m2m);rel_through(github.com/saneryao/bgadmin/models.UserRole)"`
}

// func (user *User) TableName() string {
// 	return "user"
// }
