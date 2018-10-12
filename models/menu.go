package models

type Menu struct {
	Id       int64  `form:"id" json:"id"`
	ParentId int64  `form:"parent_id" json:"parent_id"`
	Name     string `form:"name" json:"name" valid:"MaxSize(32)"`
	Icon     string `form:"icon" json:"icon" valid:"MaxSize(32);AlphaDash"`
	Url      string `form:"url" json:"url" valid:"MaxSize(255)"`
	State    int    `form:"state" json:"state" valid:"Range(0,1)"`
}

// func (menu *Menu) TableName() string {
// 	return "menu"
// }
