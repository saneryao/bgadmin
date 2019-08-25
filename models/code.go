package models

// Code 定义验证码信息
type Code struct {
	ID         int64  `form:"id" json:"id"`
	UserID     int64  `form:"user_id" json:"user_id"`
	Code       string `form:"code" json:"code"`
	CreateTime int64  `form:"create_time" json:"create_time"`
}

// func (code *Code) TableName() string {
// 	return "code"
// }
