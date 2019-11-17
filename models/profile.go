package models

// import (
// 	"github.com/astaxie/beego/validation"
// 	"github.com/beego/i18n"
// 	"strings"
// )

// Profile 定义个人信息（从用户信息拆分而来，保存个人的详细信息）
type Profile struct {
	ID      int64  `form:"id" json:"id" orm:"column(id);auto"`
	Sex     int    `form:"sex" json:"sex" orm:"column(sex)" valid:"Range(0,2)"`
	Nick    string `form:"nick"  json:"nick" orm:"column(nick)" valid:"MaxSize(32);"`
	Email   string `form:"email"  json:"email" orm:"column(email)" valid:"Email; MaxSize(128)"`
	Mobile  string `form:"mobile"  json:"mobile" orm:"column(mobile)" valid:"Mobile"`
	Tel     string `form:"tel"  json:"tel" orm:"column(tel)" valid:"Tel"`
	Birth   string `form:"birth" json:"birth" orm:"column(birth)" valid:"Length(10)"`
	Address string `form:"address" json:"address" orm:"column(adress)" valid:"MaxSize(255)"`
}

// func (profile *Profile)TableName() string {
// 	return "profile"
// }
