package article

import (
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
