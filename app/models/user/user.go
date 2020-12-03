package user

import (
	"goblog/app/models"
	"goblog/pkg/password"
	"goblog/pkg/route"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`
	// gorm:"-" —— 设置 GORM 在读写时略过此字段
	PasswordComfirm string ` gorm:"-" valid:"password_comfirm"`
}

// ComparePassword 对比密码是否匹配
func (u User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, u.Password)
}

// Link 生成用户连接
func (u User) Link() string {
	return route.Name2URL("users.show", "id", u.GetStringID())
}
