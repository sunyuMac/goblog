package category

import (
	"goblog/app/models"
	"goblog/pkg/route"
)

type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

// Link 方法用来生成分类链接
func (c Category) Link() string {
	return route.Name2URL("categories.show", "id", c.GetStringID())
}