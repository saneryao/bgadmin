package models

// import (
// 	"github.com/astaxie/beego/validation"
// 	"github.com/beego/i18n"
// 	"strings"
// )

type User struct {
	Id      int64    `form:"id" json:"id"`
	Name    string   `form:"name" json:"name" valid:"MinSize(4);MaxSize(32);AlphaDash"`
	Pwd     string   `form:"pwd" json:"pwd" valid:"MinSize(4);MaxSize(32)"`
	State   int      `form:"state" json:"state" valid:"Range(0,1)"`
	Profile *Profile `orm:"rel(one)"` // 一对一映射
}

// func (user *User) TableName() string {
// 	return "user"
// }
