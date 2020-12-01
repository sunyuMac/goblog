package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/user"
)

// ValidateRegistrationForm 验证表单，返回 errs 长度等于零即通过
func ValidateRegistrationForm(data user.User) (errs map[string][]string) {
	// 设置表单规则
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"password":         []string{"required", "min:6"},
		"password_comfirm": []string{"required"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"password": []string{
			"required:密码为必填项",
			"min:长度需大于 6",
		},
		"password_comfirm": []string{
			"required:确认密码框为必填项",
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

	if data.Password != data.PasswordComfirm {
		errs["password_cpmfirm"] = append(errs["password_comfirm"], "两次输入密码不匹配！")
	}

	return
}
