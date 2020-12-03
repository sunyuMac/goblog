package article

import (
	"goblog/app/models"
	"goblog/app/models/user"
	"goblog/pkg/route"
)

type Article struct {
	models.BaseModel

	Title  string `gorm:"column:title;type:varchar(255);not null" valid:"title"`
	Body   string `gorm:"column:body;type:text;not null" valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	User   user.User
}

// Link 获取文章连接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

// CreatedAtDate 创建日期
func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}
