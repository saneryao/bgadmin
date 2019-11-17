package models

import "time"

// RolePower 定义角色和菜单的映射关系（一个角色可以看到的菜单或操作的API）
type RolePower struct {
	ID         int       `form:"id" json:"id" orm:"column(id);auto"`
	Role       *Role     `orm:"rel(fk);column(role_id)"`
	Menu       *Menu     `orm:"rel(fk);column(menu_id)"`
	Link       *Link     `orm:"rel(fk);column(link_id)"`
	LinkPower  int       `form:"link_power" json:"link_power" orm:"column(link_power)"`
	CreateTime time.Time `form:"create_time" json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
}

// func (rm *RolePower) TableName() string {
// 	return "role_power"
// }
