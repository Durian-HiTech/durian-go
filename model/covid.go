package model

import "time"

// 中高风险地区
type HighRiskArea struct {
	Type       string `gorm:"size:15;" json:"type"`
	Province   string `gorm:"size:25;" json:"province"`
	District   string `gorm:"size:25;" json:"district"`
	Name       string `gorm:"size:105;" json:"name"`
	Coordinate string `gorm:"size:255;" json:"coordinate"`
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
	Date            time.Time `json:"date"`
	CountryName     string    `gorm:"size:255;" json:"name"`
	Info            uint64    `json:"value"`
	TotalPerHundred uint64    `json:"totalperhundred"`
	DailyPerMillion uint64    `json:"dailypermillion"`
}

// 各省的疫苗接种数据
type VaccineChina struct {
	Id       uint64  `json:"id"`
	Time     string  `json:"time"`
	Name     string  `gorm:"size:255;" json:"name"`
	Count    uint64  `json:"count"`
	IsParent uint64  `json:"isparent"`
	Parent   string  `gorm:"size:255;" json:"parent"`
	Rate     float64 `json:"rate"`
}

// [临时表]
type ProvinceVaccineData struct {
	ParentCity  string                 `gorm:"size:255;" json:"parent_city"`
	ParentCount uint64                 `json:"parent_count"`
	Rate        float64                `json:"rate"`
	ChildCity   []ChildCityVaccineData `json:"child_city"`
}

// [临时表]
type ChildCityVaccineData struct {
	Child string `gorm:"size:255;" json:"child"`
	Count uint64 `json:"count"`
}

// 新冠疫苗接种人数
type CovidGlobalVaccine struct {
	Date            time.Time `json:"date"`
	Total           uint64    `json:"total"`
	Daily           uint64    `json:"daily"`
	TotalPerHundred uint64    `json:"totalperhundred"`
	DailyPerMillion uint64    `json:"dailypermillion"`
}

// 新冠疫苗接种人数 [临时表]
type CovidVaccineNoDate struct {
	CountryName     string `gorm:"size:255;" json:"name"`
	Info            uint64 `json:"value"`
	TotalPerHundred uint64 `json:"totalperhundred"`
	DailyPerMillion uint64 `json:"dailypermillion"`
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
// ---------------------------区级行政单位的overview数据--------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
type CovidHangzhouCases struct {
	Date         time.Time `json:"date"`
	CityName     string    `gorm:"size:255;" json:"cityname"`
	ProvinceName string    `gorm:"size:255;" json:"provincename"`
	Info         uint64    `json:"value"`
}

type CovidHangzhouDeaths struct {
	Date         time.Time `json:"date"`
	CityName     string    `gorm:"size:255;" json:"cityname"`
	ProvinceName string    `gorm:"size:255;" json:"provincename"`
	Info         uint64    `json:"value"`
}

type CovidHangzhouRecovered struct {
	Date         time.Time `json:"date"`
	CityName     string    `gorm:"size:255;" json:"cityname"`
	ProvinceName string    `gorm:"size:255;" json:"provincename"`
	Info         uint64    `json:"value"`
}

// 某区级行政单位的overview列表 [临时表]
type DistrictOverview struct {
	Date     time.Time `json:"date"`
	Overview Overview  `json:"overview"`
}

// 某区级行政单位的overview和detail列表 [临时表]
type DistrictOverviewAndDetail struct {
	Date     time.Time                `json:"date"`
	Overview Overview                 `json:"overview"`
	Detailed []CovidDetailCDRDistrict `json:"detailed"`
}

// 某省下的市/直辖市的区某一天的详细信息（多了新增死亡、新增治愈这几个数据） [临时表]
type CovidDetailCDRDistrict struct {
	DistrictName string `gorm:"size:255;" json:"name"`
	Cases        uint64 `json:"cases"`
	NowCases     uint64 `json:"nowcases"`
	NewCases     uint64 `json:"newcases"`
	Vaccine      uint64 `json:"vaccine"`    // 其实是全空
	NewVaccine   uint64 `json:"newvaccine"` // 其实是全空
	Recovered    uint64 `json:"recovered"`
	NewRecovered uint64 `json:"newrecovered"`
	Deaths       uint64 `json:"deaths"`
	NewDeaths    uint64 `json:"newdeaths"`
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ---------------------------某个国家的overview数据-----------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

// 某国家的overview+detail列表 [临时表]
type CountryOverviewAndDetails struct {
	Date     time.Time          `json:"date"`
	Overview Overview           `json:"overview"`
	Detailed []CovidCDRProvince `json:"detailed"`
}

// 某国家的overview数据 [临时表]
type Overview struct {
	NowCases  NowCases  `json:"nowcases"`
	Cases     Cases     `json:"cases"`
	Deaths    Deaths    `json:"deaths"`
	Vaccine   Vaccine   `json:"vaccine"`
	Recovered Recovered `json:"recovered"`
}

// 某一天的确诊数（现存确诊、新增确诊） [临时表]
type NowCases struct {
	NowNum uint64 `json:"nownum"`
	NewNum uint64 `json:"newnum"`
}

// 某一天的累积确诊数 [临时表]
type Cases struct {
	NowNum uint64 `json:"nownum"`
	NewNum uint64 `json:"newnum"`
}

// 某一天的死亡数（累计死亡、新增死亡） [临时表]
type Deaths struct {
	NowNum uint64 `json:"nownum"`
	NewNum uint64 `json:"newnum"`
}

// 某一天的接种数（累计接种、新增接种） [临时表]
type Vaccine struct {
	NowNum uint64 `json:"nownum"`
	NewNum uint64 `json:"newnum"`
}

// 某一天的治愈数（累计死亡、新增死亡） [临时表]
type Recovered struct {
	NowNum uint64 `json:"nownum"`
	NewNum uint64 `json:"newnum"`
}

// 省份某一天的信息 [临时表]
type CovidCDRProvince struct {
	ProvinceName string `gorm:"size:255;" json:"name"`
	NowCases     uint64 `json:"nowcases"`
	Cases        uint64 `json:"cases"`
	Deaths       uint64 `json:"deaths"`
	Recovered    uint64 `json:"recovered"`
}

// 世界的overview+detail列表 用于HomeData.json[临时表]
type GlobalOverviewAndDetails struct {
	Overview Overview                `json:"overview"`
	Detailed []CovidDetailCDRCountry `json:"detailed"`
}

// 中国的overview+detail列表 用于HomeData.json[临时表]
type ChinaOverviewAndDetails struct {
	Overview Overview                 `json:"overview"`
	Detailed []CovidDetailCDRProvince `json:"detailed"`
}

// 世界的overview+detail列表 用于世界每一天的所有国家的数据[临时表]
type GlobalOverviewAndDetailsWithDate struct {
	Date     time.Time               `json:"date"`
	Overview Overview                `json:"overview"`
	Detailed []CovidDetailCDRCountry `json:"detailed"`
}

// 某省份某一天的详细信息（多了新增死亡、新增治愈这几个数据） [临时表]
type CovidDetailCDRProvince struct {
	ProvinceName string `gorm:"size:255;" json:"name"`
	Cases        uint64 `json:"cases"`
	NowCases     uint64 `json:"nowcases"`
	NewCases     uint64 `json:"newcases"`
	Vaccine      uint64 `json:"vaccine"`    // 其实是全空
	NewVaccine   uint64 `json:"newvaccine"` // 其实是全空
	Recovered    uint64 `json:"recovered"`
	NewRecovered uint64 `json:"newrecovered"`
	Deaths       uint64 `json:"deaths"`
	NewDeaths    uint64 `json:"newdeaths"`
}

// 某国家某一天的详细信息（多了新增死亡、新增治愈这几个数据） [临时表]
type CovidDetailCDRCountry struct {
	CountryName  string `gorm:"size:255;" json:"name"`
	Cases        uint64 `json:"cases"`
	NowCases     uint64 `json:"nowcases"`
	NewCases     uint64 `json:"newcases"`
	Vaccine      uint64 `json:"vaccine"`    // 其实是全空
	NewVaccine   uint64 `json:"newvaccine"` // 其实是全空
	Recovered    uint64 `json:"recovered"`
	NewRecovered uint64 `json:"newrecovered"`
	Deaths       uint64 `json:"deaths"`
	NewDeaths    uint64 `json:"newdeaths"`
}

// 新冠感染/死亡/治愈/疫苗接种人数 [信息综合] [根据时间分组]
type CovidCDRVResponseProvince struct {
	Case      []CovidProvinceCases     `json:"cases"`
	Deaths    []CovidProvinceDeaths    `json:"deaths"`
	Recovered []CovidProvinceRecovered `json:"recovered"`
	Vaccine   []CovidProvinceRecovered `json:"vaccine"` // 其实是全空
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ---------------------------疫苗的overview数据---------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// 某国家某一天的详细信息 [临时表]
type CovidVaccineCountry struct {
	CountryName     string `gorm:"size:255;" json:"name"`
	Vaccine         uint64 `json:"vaccine"`
	NewVaccine      uint64 `json:"newvaccine"`
	TotalPerHundred uint64 `json:"totalperhundred"`
	DailyPerMillion uint64 `json:"dailypermillion"`
}

// 全球某一天的接种信息 [临时表]
type CovidVaccineGlobalOview struct {
	Vaccine         uint64 `json:"vaccine"`
	NewVaccine      uint64 `json:"newvaccine"`
	TotalPerHundred uint64 `json:"totalperhundred"`
	DailyPerMillion uint64 `json:"dailypermillion"`
}

// 世界疫苗接种的overview+detail列表 用于世界每一天的所有国家的接种数据[临时表]
type GlobalVaccineOverviewAndDetailsWithDate struct {
	Date     time.Time               `json:"date"`
	Overview CovidVaccineGlobalOview `json:"overview"`
	Detailed []CovidVaccineCountry   `json:"detailed"`
}
