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

func QueryASubscriptionByID(subscriptionID uint64) (subscription model.Subscription, notFound bool) {
	err := global.DB.Where("subscription_id = ?", subscriptionID).First(&subscription).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return subscription, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return subscription, false
	}
}

func QueryASubscriptionByUserIDAndCityName(userID uint64, cityName string) (subscription model.Subscription, notFound bool) {
	err := global.DB.Where("user_id = ? AND city_name = ?", userID, cityName).First(&subscription).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return subscription, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return subscription, false
	}
}

func DeleteASubscription(subscriptionID uint64) (err error) {
	var subscription model.Subscription
	err = global.DB.First(&subscription, subscriptionID).Error
	_ = global.DB.Delete(&subscription).Error
	return err
}

func QueryAllSubscriptions(userID uint64) (subscriptions []model.Subscription) {
	global.DB.Where("user_id = ?", userID).Find(&subscriptions)
	return subscriptions
}

func QueryANewsByID(NewsID uint64) (news model.News, notFound bool) {
	err := global.DB.Where("news_id = ?", news).First(&news).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return news, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return news, false
	}
}

func QueryAllNews() (subscriptions []model.Subscription) {
	global.DB.Find(&subscriptions)
	return subscriptions
}
