package category

import "goblog/pkg/model"

// Create 创建分类
func (category *Category) Create() (err error) {
	err = model.DB.Create(&category).Error
	return
}
