package models

// UserRole 定义用户和角色映射关系
type UserRole struct {
	ID     int64
	UserID int64
	RoleID int64
}

// func (ur *UserRole) TableName() string {
// 	return "user_role"
// }
