package model

import "time"

// 用户
type User struct {
	UserID      uint64 `gorm:"primary_key; not null;" json:"user_id"`
	Username    string `gorm:"size:25; not null; unique" json:"username"`
	Password    string `gorm:"size:25; not null" json:"password"`
	UserType    uint64 `gorm:"default:0" json:"user_type"` // 0: 普通用户，1: 认证机构用户
	Affiliation string `gorm:"size:25;" json:"affiliation"`
}

// 订阅
type Subscription struct {
	SubscriptionID uint64 `gorm:"primary_key;" json:"subscription_id"`
	UserID         uint64 `gorm:"not null" json:"user_id"`
	CityName       string `gorm:"size:25;not null" json:"city_name"`
}

// 新闻
type News struct {
	NewsID          uint64    `gorm:"primary_key;" json:"news_id"`
	NewsTitle       string    `gorm:"size:55; not null" json:"news_title"`
	NewsContent     string    `gorm:"size:2550; not null" json:"news_content"`
	NewsCreatedTime time.Time `json:"news_created_time"`
}

// 知识版块中的问题
type Question struct {
	QuestionID      uint64    `gorm:"primary_key;" json:"question_id"`
	UserID          uint64    `gorm:"not null" json:"user_id"`
	QuestionTitle   string    `gorm:"size:200;not null" json:"question_title"`
	QuestionContent string    `gorm:"size:200;not null" json:"question_content"`
	QuestionTime    time.Time `json:"question_time"`
}

// 知识版块中对问题的评论
type Comment struct {
	CommentID      uint64    `gorm:"primary_key;" json:"comment_id"`
	UserID         uint64    `gorm:"not null" json:"user_id"`
	QuestionID     uint64    `gorm:"not null" json:"question_id"`
	CommentContent string    `gorm:"size:200;not null" json:"comment_content"`
	CommentTime    time.Time `json:"comment_time"`
	UserType       uint64    `gorm:"default:0" json:"user_type"`
}

// 高风险地区
type HighRiskArea struct {
	Type     string `gorm:"size:15;" json:"type"`
	Province string `gorm:"size:25;" json:"province"`
	District string `gorm:"size:25;" json:"district"`
	Name     string `gorm:"size:105;" json:"name"`
}

// 直接传输数据
type DirectData struct {
	Name    string `gorm:"size:25;" json:"name"`
	Content string `gorm:"size:255550;" json:"content"`
}
