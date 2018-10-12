package models

type Code struct {
	Id         int64
	UserId     int64
	Code       string
	CreateTime int64
}

// func (code *Code) TableName() string {
// 	return "code"
// }
