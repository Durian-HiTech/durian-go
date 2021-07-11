package model

import "time"

// 用户
type User struct {
	UserID      uint64 `gorm:"primary_key; not null;" json:"user_id"`
	Username    string `gorm:"size:25; not null; unique" json:"username"`
	Password    string `gorm:"size:25; not null" json:"password"`
	UserInfo    string `gorm:"size:255;" json:"user_info"`
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

// 公告
type Notice struct {
	NoticeID          uint64    `gorm:"primary_key;" json:"notice_id"`
	NoticeTitle       string    `gorm:"size:55; not null" json:"notice_title"`
	NoticeContent     string    `gorm:"size:2550; not null" json:"notice_content"`
	NoticeCreatedTime time.Time `json:"notice_created_time"`
}

// 谣言
type Rumor struct {
	RumorID          uint64    `gorm:"primary_key;" json:"rumor_id"`
	RumorTitle       string    `gorm:"size:55; not null" json:"rumor_title"`
	RumorContent     string    `gorm:"size:2550; not null" json:"rumor_content"`
	RumorType        uint64    `gorm:"default:0; not null" json:"rumor_type"` // 0: 谣言, 1: 事实, 2: 辟谣, 3: 误区
	RumorCreatedTime time.Time `json:"rumor_created_time"`
}

// 防疫小知识
type Knowledge struct {
	KnowledgeID      uint64 `gorm:"primary_key;" json:"knowledge_id"`
	KnowledgeTitle   string `gorm:"size:55; not null" json:"knowledge_title"`
	KnowledgeContent string `gorm:"size:2550; not null" json:"knowledge_content"`
}

// 知识版块中的问题
type Question struct {
	QuestionID      uint64    `gorm:"primary_key;" json:"question_id"`
	UserID          uint64    `gorm:"not null" json:"user_id"`
	QuestionTitle   string    `gorm:"size:200;not null" json:"question_title"`
	QuestionContent string    `gorm:"size:200;not null" json:"question_content"`
	QuestionTime    time.Time `json:"question_time"`
}

type QuestionWithUserInfo struct {
	Question
	Username string `json:"username"`
	UserInfo string `json:"user_info"`
}

// 知识版块中对问题的评论
type Comment struct {
	CommentID      uint64    `gorm:"primary_key;" json:"comment_id"`
	UserID         uint64    `gorm:"not null" json:"user_id"`
	QuestionID     uint64    `gorm:"not null" json:"question_id"`
	CommentContent string    `gorm:"size:255;not null" json:"comment_content"`
	CommentTime    time.Time `json:"comment_time"`
	UserType       uint64    `gorm:"default:0" json:"user_type"`
}

type CommentWithUserInfo struct {
	Comment
	Username string `json:"username"`
	UserInfo string `json:"user_info"`
}

type FlightDomestic struct {
	AirlineName          string `gorm:"size:255;" json:"airline_name"`
	FlightNumber         string `gorm:"size:255;" json:"flight_number"`
	DepartureDate        string `gorm:"size:255;" json:"departure_date"`
	ArrivalDate          string `gorm:"size:255;" json:"arrival_date"`
	DepartureCityName    string `gorm:"size:255;" json:"departure_city_name"`
	DepartureAirportName string `gorm:"size:255;" json:"departure_airport_name"`
	ArrivalCityName      string `gorm:"size:255;" json:"arrival_city_name"`
	ArrivalAirportName   string `gorm:"size:255;" json:"arrival_airport_name"`
}

type FlightDomesticWithStatus struct {
	FlightDomestic
	Status string `gorm:"size:255;" json:"status"`
}

type TrainDomestic struct {
	DepartureCity    string `gorm:"size:255;" json:"departure_city"`
	DepartureTime    string `gorm:"size:255;" json:"departure_time"`
	ArrivalCity      string `gorm:"size:255;" json:"arrival_city"`
	ArrivalTime      string `gorm:"size:255;" json:"arrival_time"`
	TrainNumber      string `gorm:"size:255;" json:"train_number"`
	DepartureStation string `gorm:"size:255;" json:"departure_station"`
	ArrivalStation   string `gorm:"size:255;" json:"arrival_station"`
}

type TrainDomesticWithStatus struct {
	TrainDomestic
	Status string `gorm:"size:255;" json:"status"`
}
