package models

// import (
// 	"github.com/astaxie/beego/validation"
// 	"github.com/beego/i18n"
// 	"strings"
// )

type Profile struct {
	Id      int64  `form:"id" json:"id"`
	Sex     int    `form:"sex" json:"sex" valid:"Range(0,2)"`
	Nick    string `form:"nick"  json:"nick" valid:"MaxSize(32);"`
	Email   string `form:"email"  json:"email" valid:"Email; MaxSize(128)"`
	Mobile  string `form:"mobile"  json:"mobile" valid:"Mobile"`
	Tel     string `form:"tel"  json:"tel" valid:"Tel"`
	Birth   string `form:"birth" json:"birth" valid:"Length(10)"`
	Address string `form:"address" json:"address" valid:"MaxSize(255)"`
}

// func (profile *Profile)TableName() string {
// 	return "profile"
// }
