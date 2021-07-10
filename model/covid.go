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

// 新冠感染/死亡/治愈/疫苗接种人数 [信息综合]
type CovidCDRV struct {
	Case      []CovidCases     `json:"cases"`
	Deaths    []CovidDeaths    `json:"deaths"`
	Recovered []CovidRecovered `json:"recovered"`
	Vaccine   []CovidVaccine   `json:"vaccine"`
}

// 新冠感染/死亡/治愈/疫苗接种人数 [信息综合] [根据时间分组]
type CovidCDRVResponse struct {
	Case      []CovidCasesResponse     `json:"cases"`
	Deaths    []CovidDeathsResponse    `json:"deaths"`
	Recovered []CovidRecoveredResponse `json:"recovered"`
	Vaccine   []CovidVaccineResponse   `json:"vaccine"`
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ---------------------------中国各省份数据--------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
type CovidChinaCases struct {
	Date         time.Time `json:"date"`
	ProvinceName string    `gorm:"size:255;" json:"name"`
	Info         uint64    `json:"value"`
}

type CovidChinaDeaths struct {
	Date         time.Time `json:"date"`
	ProvinceName string    `gorm:"size:255;" json:"name"`
	Info         uint64    `json:"value"`
}

type CovidChinaRecovered struct {
	Date         time.Time `json:"date"`
	ProvinceName string    `gorm:"size:255;" json:"name"`
	Info         uint64    `json:"value"`
}

// [临时表]
type CovidChinaCasesNoDate struct {
	ProvinceName string `gorm:"size:255;" json:"name"`
	Info         uint64 `json:"value"`
}

// [临时表]
type CovidChinaDeathsNoDate struct {
	ProvinceName string `gorm:"size:255;" json:"name"`
	Info         uint64 `json:"value"`
}

// [临时表]
type CovidChinaRecoveredNoDate struct {
	ProvinceName string `gorm:"size:255;" json:"name"`
	Info         uint64 `json:"value"`
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

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ---------------------------某个国家的overview数据-----------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

// 某国家的overvie+detail列表 [临时表]
type CountryOverviewAndDetails struct {
	Date     time.Time          `json:"date"`
	Overview Overview           `json:"overview"`
	Detailed []CovidCDRProvince `json:"detailed"`
}

// 某国家的overview数据 [临时表]
type Overview struct {
	NowCases  NowCases  `json:"now_cases"`
	Cases     Cases     `json:"cases"`
	Deaths    Deaths    `json:"deaths"`
	Vaccine   Vaccine   `json:"vaccine"`
	Recovered Recovered `json:"recovered"`
}

// 某一天的确诊数（现存确诊、新增确诊） [临时表]
type NowCases struct {
	NowNum uint64 `json:"now_num"`
	NewNum uint64 `json:"new_num"`
}

// 某一天的累积确诊数 [临时表]
type Cases struct {
	NowNum uint64 `json:"now_num"`
}

// 某一天的死亡数（累计死亡、新增死亡） [临时表]
type Deaths struct {
	NowNum uint64 `json:"now_num"`
	NewNum uint64 `json:"new_num"`
}

// 某一天的接种数（累计接种、新增接种） [临时表]
type Vaccine struct {
	NowNum uint64 `json:"now_num"`
	NewNum uint64 `json:"new_num"`
}

// 某一天的治愈数（累计死亡、新增死亡） [临时表]
type Recovered struct {
	NowNum uint64 `json:"now_num"`
	NewNum uint64 `json:"new_num"`
}

// 省份某一天的信息 [临时表]
type CovidCDRProvince struct {
	ProvinceName string `gorm:"size:255;" json:"name"`
	NowCases     uint64 `json:"now_cases"`
	Cases        uint64 `json:"cases"`
	Deaths       uint64 `json:"deaths"`
	Recovered    uint64 `json:"recovered"`
}

// 世界的overvie+detail列表 用于HomeData.json[临时表]
type GlobalOverviewAndDetails struct {
	Overview Overview                `json:"overview"`
	Detailed []CovidDetailCDRCountry `json:"detailed"`
}

// 中国的overvie+detail列表 用于HomeData.json[临时表]
type ChinaOverviewAndDetails struct {
	Overview Overview                 `json:"overview"`
	Detailed []CovidDetailCDRProvince `json:"detailed"`
}

// 某省份某一天的详细信息（多了新增死亡、新增治愈这几个数据） [临时表]
type CovidDetailCDRProvince struct {
	ProvinceName string `gorm:"size:255;" json:"name"`
	Cases        uint64 `json:"cases"`
	NowCases     uint64 `json:"now_cases"`
	NewCases     uint64 `json:"new_cases"`
	Vaccine      uint64 `json:"vaccine"`     // 其实是全空
	NewVaccine   uint64 `json:"new_vaccine"` // 其实是全空
	Recovered    uint64 `json:"recovered"`
	NewRecovered uint64 `json:"new_recovered"`
	Deaths       uint64 `json:"deaths"`
	NewDeaths    uint64 `json:"new_deaths"`
}

// 某国家某一天的详细信息（多了新增死亡、新增治愈这几个数据） [临时表]
type CovidDetailCDRCountry struct {
	CountryName  string `gorm:"size:255;" json:"name"`
	Cases        uint64 `json:"cases"`
	NowCases     uint64 `json:"now_cases"`
	NewCases     uint64 `json:"new_cases"`
	Vaccine      uint64 `json:"vaccine"`     // 其实是全空
	NewVaccine   uint64 `json:"new_vaccine"` // 其实是全空
	Recovered    uint64 `json:"recovered"`
	NewRecovered uint64 `json:"new_recovered"`
	Deaths       uint64 `json:"deaths"`
	NewDeaths    uint64 `json:"new_deaths"`
}

// 新冠感染/死亡/治愈/疫苗接种人数 [信息综合] [根据时间分组]
type CovidCDRVResponseProvince struct {
	Case      []CovidProvinceCases     `json:"cases"`
	Deaths    []CovidProvinceDeaths    `json:"deaths"`
	Recovered []CovidProvinceRecovered `json:"recovered"`
	Vaccine   []CovidProvinceRecovered `json:"vaccine"` // 其实是全空
}
