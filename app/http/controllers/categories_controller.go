package controllers

import (
	"fmt"
	"goblog/app/models/category"
	"goblog/app/requests"
	"goblog/pkg/view"
	"net/http"
)

type CategoriesController struct {
	BaseController
}

func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	errors := requests.ValidateCategoryForm(_category)
	if len(errors) != 0 {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   errors,
		}, "categories.create")

		return
	}

	err := _category.Create()
	if _category.ID > 0 && err == nil {
		fmt.Fprint(w, "创建成功")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "创建文章分类失败，请联系管理员")
	}
}

func (*CategoriesController) Show(w http.ResponseWriter, r *http.Request)  {
	
}
