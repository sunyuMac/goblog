package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

type UserController struct {
	BaseController
}

func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_, err := user.Get(id)
	if err != nil {
		uc.ResponseForSQLError(w, err)
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
