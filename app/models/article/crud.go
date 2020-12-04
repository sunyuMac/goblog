package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
)

func Get(idstr string) (Article, error) {
	article := Article{}
	id := types.StringToInt(idstr)
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

// GetAll 获取全部文章
func GetAll(r *http.Request, perPage int) (articles []Article, viewData pagination.ViewData, err error) {
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)
	viewData = _pager.Paging()
	_pager.Results(&articles)

	return
}

func (article *Article) Create() (err error) {
	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
	}
	return
}

func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

func GetByUserId(userId string) (articles []Article, err error) {
	err = model.DB.Preload("User").Where("user_id = ?", userId).Find(&articles).Error

	return
}
