package model

import "time"

type User struct {
	UserID      uint64 `gorm:"primary_key; not null;" json:"user_id"`
	Username    string `gorm:"size:25; not null; unique" json:"username"`
	Password    string `gorm:"size:25; not null" json:"password"`
	UserType    uint64 `gorm:"default:0" json:"user_type"` // 0: 普通用户，1: 认证机构用户
	Affiliation string `gorm:"size:25;" json:"affiliation"`
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

type Question struct {
	QuestionID      uint64    `gorm:"primary_key;" json:"question_id"`
	UserID          uint64    `gorm:"not null" json:"user_id"`
	QuestionContent string    `gorm:"size:200;not null" json:"question_content"`
	QuestionTime    time.Time `json:"question_time"`
}

type Comment struct {
	CommentID      uint64    `gorm:"primary_key;" json:"comment_id"`
	UserID         uint64    `gorm:"not null" json:"user_id"`
	QuestionID     uint64    `gorm:"not null" json:"question_id"`
	CommentContent string    `gorm:"size:200;not null" json:"comment_content"`
	CommentTime    time.Time `json:"comment_time"`
	Valid          bool      `gorm:"default:false" json:"valid"`
	UserType       uint64    `gorm:"default:0" json:"user_type"`
}
