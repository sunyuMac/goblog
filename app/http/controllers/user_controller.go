package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
}

func (u *UserController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_, err := user.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 用户未找到")
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		articles, err := article.GetByUserId(id)
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "500 服务器内部错误")
		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index", "articles._article_meta")
		}
	}

}
