package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/types"
	"goblog/pkg/model"
)

func Get(idstr string) (Article, error) {
	article := Article{}
	id := types.StringToInt(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

func GetAll() ([]Article, error)  {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func (article *Article) Create() (err error) {
	if err =  model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}
