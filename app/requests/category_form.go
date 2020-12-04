package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/category"
)

// ValidateCategoryForm 验证创建分类表单
func ValidateCategoryForm(data category.Category) map[string][]string {
	// 设置表单规则
	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
		},
	}
	// 配置选项
	opt := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}
	// 开始认证
	return govalidator.New(opt).ValidateStruct()
}
