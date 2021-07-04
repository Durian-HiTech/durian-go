package model

import "time"

type User struct {
	UserID   uint64 `gorm:"primary_key; not null;" json:"user_id"`
	Username string `gorm:"size:25; not null; unique" json:"username"`
	Password string `gorm:"size:25; not null" json:"password"`
}
type Subscription struct {
	SubscriptionID uint64 `gorm:"primary_key;" json:"subscription_id"`
	UserID         uint64 `gorm:"not null" json:"user_id"`
	CityName       string `gorm:"size:25;not null" json:"city_name"`
}

type News struct {
	NewsID          uint64    `gorm:"primary_key;" json:"news_id"`
	NewsTitle       string    `gorm:"size:55; not null" json:"news_title"`
	NewsContent     string    `gorm:"size:2550; not null" json:"news_content"`
	NewsCreatedTime time.Time `json:"news_created_time"`
}
