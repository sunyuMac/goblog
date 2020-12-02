package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

// Create
func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
	}

	return
}

// Get
func Get(idstr string) (user User, err error) {
	id := types.StringToInt(idstr)
	err = model.DB.First(&user, id).Error
	return
}

// GetByEmail
func GetByEmail(email string) (user User, err error)  {
	err = model.DB.Where("email = ?", email).First(&user).Error
	return
}
