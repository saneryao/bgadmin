package models

import "time"

// Link 定义链接信息
type Link struct {
	ID         int       `form:"id" json:"id" orm:"column(id);auto"`
	Name       string    `form:"name" json:"name" orm:"column(name)" valid:"MaxSize(32)"`
	URL        string    `form:"url" json:"url" orm:"column(url)" valid:"MaxSize(255)"`
	CreateTime time.Time `form:"create_time" json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	UpdateTime time.Time `form:"update_time" json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
}

// func (link *Link) TableName() string {
// 	return "link"
// }
