package service

import (
	"errors"

	"github.com/TualatinX/durian-go/global"
	"github.com/TualatinX/durian-go/model"
	"gorm.io/gorm"
)

func CreateAUser(user *model.User) (err error) {
	if err = global.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func QueryAUserByID(userID uint64) (user model.User, notFound bool) {
	err := global.DB.Where("user_id = ?", userID).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return user, false
	}
}

func QueryAUserByUsername(username string) (user model.User, notFound bool) {
	err := global.DB.Where("username = ?", username).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return user, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return user, false
	}
}

func UpdateAUser(user *model.User, username string, password string) error {
	user.Username = username
	user.Password = password
	err := global.DB.Save(user).Error
	return err
}

func CreateASubscription(userID uint64, cityName string) (err error) {
	subscription := model.Subscription{UserID: userID, CityName: cityName}
	if err = global.DB.Create(&subscription).Error; err != nil {
		return err
	}
	return
}

// func DeleteAFavorite(favorID uint64) (err error) {
// 	var favorite model.Favorite
// 	err = global.DB.First(&favorite, favorID).Error
// 	_ = global.DB.Delete(&favorite).Error
// 	return err
// }

// func QueryAllFavorites(userID uint64) (favorites []model.Favorite) {
// 	global.DB.Where("user_id = ?", userID).Find(&favorites)
// 	return favorites
// }
