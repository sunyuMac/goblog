package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/article"
)

// ValidateArticleFormData 验证文章表单信息
func ValidateArticleFormData(data article.Article) (errs map[string][]string) {
	// 设置表单规则
	rules := govalidator.MapData{
		"title": []string{"required", "min:3", "max:40"},
		"body":  []string{"required", "min:10"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:文章标题为必填项",
			"min:标题长度需大于3",
			"max:标题长度需小于40",
		},
		"body": []string{
			"required:文章内容为必填项",
			"min:文章长度必须大于10",
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
	errs = govalidator.New(opt).ValidateStruct()

	return
}
