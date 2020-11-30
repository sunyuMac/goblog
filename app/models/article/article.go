package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

type Article struct {
	models.BaseModel

	Title string `gorm:"column:title;type:varchar(255);not null"`
	Body  string `gorm:"column:body;type:text;not null"`
}

// Link 获取文章连接
func (a *Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}
