package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/pkg/view"
	"net/http"
)

// AuthController 处理静态页面
type AuthController struct {
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	_user := user.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	_user.Create()
	if _user.ID > 0 {
		fmt.Fprint(w, "插入成功，ID 为"+_user.GetStringID())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "创建用户失败，请联系管理员")
	}
}
