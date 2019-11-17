package models

import "time"

// Code 定义验证码信息
type Code struct {
	ID         int64     `form:"id" json:"id" orm:"column(id);auto"`
	UserID     int64     `form:"user_id" json:"user_id" orm:"column(user_id)"`
	Code       string    `form:"code" json:"code" orm:"column(code)"`
	CreateTime time.Time `form:"create_time" json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
}

// func (code *Code) TableName() string {
// 	return "code"
// }
