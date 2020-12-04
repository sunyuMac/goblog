package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/models/category"
	"goblog/app/requests"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

type CategoriesController struct {
	BaseController
}

// Create 创建分类页面
func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

// Store 保存分类
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

// Show 展示分类下文章
func (cc *CategoriesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的数据
	_, err := category.Get(id)
	if err != nil {
		cc.ResponseForSQLError(w, err)
		return
	}

	// 3. 获取结果集
	articles, pagerData, err := article.GetByCategoryID(id, r, 5)

	if err != nil {
		cc.ResponseForSQLError(w, err)
		return
	}

	// ---  2. 加载模板 ---
	view.Render(w, view.D{
		"Articles":  articles,
		"PagerData": pagerData,
	}, "articles.index", "articles._article_meta")

}
