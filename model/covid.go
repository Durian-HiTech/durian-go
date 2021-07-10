package model

import "time"

// 中高风险地区
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

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ---------------------------省份数据--------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

// 新冠感染人数
type CovidProvinceCases struct {
	Date        time.Time `json:"date"`
	CountryName string    `gorm:"size:255;" json:"name"`
	Info        string    `gorm:"size:2555550;" json:"value"`
}

// 新冠死亡人数 [Province]
type CovidProvinceDeaths struct {
	Date        time.Time `json:"date"`
	CountryName string    `gorm:"size:255;" json:"name"`
	Info        string    `gorm:"size:2555550;" json:"value"`
}

// 新冠治愈人数 [Province]
type CovidProvinceRecovered struct {
	Date        time.Time `json:"date"`
	CountryName string    `gorm:"size:255;" json:"name"`
	Info        string    `gorm:"size:2555550;" json:"value"`
}

// 新冠感染/死亡/治愈/疫苗接种人数【信息综合】 [根据时间分组]
type CovidCDRVResponseProvince struct {
	Case      []CovidProvinceCases     `json:"cases"`
	Deaths    []CovidProvinceDeaths    `json:"deaths"`
	Recovered []CovidProvinceRecovered `json:"recovered"`
	Vaccine   []CovidProvinceRecovered `json:"vaccine"` // 其实是全空
}
