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
	Content string `gorm:"size:2555500;" json:"content"`
}

// 新冠感染人数
type CovidCases struct {
	Date        time.Time `json:"date"`
	CountryName string    `gorm:"size:255;" json:"name"`
	Info        uint64    `json:"value"`
}

// 新冠感染人数 [临时表]
type CovidCasesNoDate struct {
	CountryName string `gorm:"size:255;" json:"name"`
	Info        uint64 `json:"value"`
}

// 新冠感染人数 [根据时间分组]
type CovidCasesResponse struct {
	Date  time.Time          `json:"date"`
	Value []CovidCasesNoDate `json:"value"`
}

// 新冠死亡人数
type CovidDeaths struct {
	Date        time.Time `json:"date"`
	CountryName string    `gorm:"size:255;" json:"name"`
	Info        uint64    `json:"value"`
}

// 新冠死亡人数 [临时表]
type CovidDeathsNoDate struct {
	CountryName string `gorm:"size:255;" json:"name"`
	Info        uint64 `json:"value"`
}

// 新冠死亡人数 [根据时间分组]
type CovidDeathsResponse struct {
	Date  time.Time           `json:"date"`
	Value []CovidDeathsNoDate `json:"value"`
}

// 新冠治愈人数
type CovidRecovered struct {
	Date        time.Time `json:"date"`
	CountryName string    `gorm:"size:255;" json:"name"`
	Info        uint64    `json:"value"`
}

// 新冠治愈人数 [临时表]
type CovidRecoveredNoDate struct {
	CountryName string `gorm:"size:255;" json:"name"`
	Info        uint64 `json:"value"`
}

// 新冠治愈人数 [根据时间分组]
type CovidRecoveredResponse struct {
	Date  time.Time              `json:"date"`
	Value []CovidRecoveredNoDate `json:"value"`
}

// 新冠疫苗接种人数
type CovidVaccine struct {
	Date        time.Time `json:"date"`
	CountryName string    `gorm:"size:255;" json:"name"`
	Info        uint64    `json:"value"`
}

// 新冠疫苗接种人数 [临时表]
type CovidVaccineNoDate struct {
	CountryName string `gorm:"size:255;" json:"name"`
	Info        uint64 `json:"value"`
}

// 新冠疫苗接种人数 [根据时间分组]
type CovidVaccineResponse struct {
	Date  time.Time            `json:"date"`
	Value []CovidVaccineNoDate `json:"value"`
}

// 新冠感染/死亡/治愈/疫苗接种人数【信息综合】
type CovidCDRV struct {
	Case      []CovidCases     `json:"cases"`
	Deaths    []CovidDeaths    `json:"deaths"`
	Recovered []CovidRecovered `json:"recovered"`
	Vaccine   []CovidVaccine   `json:"vaccine"`
}

// 新冠感染/死亡/治愈/疫苗接种人数【信息综合】 [根据时间分组]
type CovidCDRVResponse struct {
	Case      []CovidCasesResponse     `json:"cases"`
	Deaths    []CovidDeathsResponse    `json:"deaths"`
	Recovered []CovidRecoveredResponse `json:"recovered"`
	Vaccine   []CovidVaccineResponse   `json:"vaccine"`
}
