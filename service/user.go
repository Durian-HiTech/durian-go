package service

import (
	"github.com/TualatinX/durian-go/global"
	"github.com/TualatinX/durian-go/model"
)

func CreateAUser(user *model.User) (err error) {
	if err = global.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}
