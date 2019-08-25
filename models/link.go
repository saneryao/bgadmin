package models

// Link 定义链接信息
type Link struct {
	ID   int64  `form:"id" json:"id"`
	Name string `form:"name" json:"name" valid:"MaxSize(32)"`
	URL  string `form:"url" json:"url" valid:"MaxSize(255)"`
}

// func (link *Link) TableName() string {
// 	return "link"
// }
