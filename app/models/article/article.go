package article

import (
	"goblog/app/models"
)

type Article struct {
	models.BaseModel

	Title string
	Body  string
}
