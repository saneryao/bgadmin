package models

type UserRole struct {
	Id     int64
	UserId int64
	RoleId int64
}

// func (ur *UserRole) TableName() string {
// 	return "user_role"
// }
