package models

import "time"

// UserRole 定义用户和角色映射关系
type UserRole struct {
	ID         int64     `form:"id" json:"id" orm:"column(id);auto"`
	User       *User     `orm:"rel(fk)"`
	Role       *Role     `orm:"rel(fk)"`
	CreateTime time.Time `form:"create_time" json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
}

// func (ur *UserRole) TableName() string {
// 	return "user_role"
// }
