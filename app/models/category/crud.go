package category

import (
	"goblog/pkg/model"
	"goblog/pkg/types"
)

// Create 创建分类
func (category *Category) Create() (err error) {
	err = model.DB.Create(&category).Error
	return
}

func All() (categories []Category,err error)  {
	err = model.DB.Find(&categories).Error
	return
}

func Get(idStr string) (category Category,err error)  {
	err = model.DB.Where("id", types.StringToInt(idStr)).First(&category).Error
	return
}