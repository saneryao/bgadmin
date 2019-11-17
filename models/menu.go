package models

import "time"

// Menu 定义菜单信息
type Menu struct {
	ID         int       `form:"id" json:"id" orm:"column(id);auto"`
	ParentID   int       `form:"parent_id" json:"parent_id" orm:"column(parent_id)"`
	Name       string    `form:"name" json:"name" orm:"column(name)" valid:"MaxSize(32)"`
	Icon       string    `form:"icon" json:"icon" orm:"column(icon)" valid:"MaxSize(32);AlphaDash"`
	URL        string    `form:"url" json:"url" orm:"column(url)" valid:"MaxSize(255)"`
	State      int       `form:"state" json:"state" orm:"column(state)" valid:"Range(0,1)"`
	CreateTime time.Time `form:"create_time" json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime time.Time `form:"update_time" json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
}

// func (menu *Menu) TableName() string {
// 	return "menu"
// }
