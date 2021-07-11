package global

import (
	"gorm.io/gorm"
)

var (
	// 全局变量 DB
	DB *gorm.DB
)

// 省份名字与拼音的映射
func MapPinYin(provinceName string) (provinceNamePinYin string) {
	nameMap := make(map[string]string)

	nameMap["安徽省"] = "Anhui"
	nameMap["北京市"] = "Beijing"
	nameMap["重庆市"] = "Chongqing"
	nameMap["福建省"] = "Fujian"
	nameMap["甘肃省"] = "Gansu"
	nameMap["广东省"] = "Guangdong"
	nameMap["广西壮族自治区"] = "Guangxi"
	nameMap["贵州省"] = "Guizhou"
	nameMap["海南省"] = "Hainan"
	nameMap["河北省"] = "Hebei"
	nameMap["黑龙江省"] = "Heilongjiang"
	nameMap["河南省"] = "Henan"
	nameMap["香港"] = "Hong Kong"
	nameMap["湖北省"] = "Hubei"
	nameMap["湖南省"] = "Hunan"
	nameMap["江苏省"] = "Jiangsu"
	nameMap["江西省"] = "Jiangxi"
	nameMap["吉林省"] = "Jilin"
	nameMap["辽宁省"] = "Liaoning"
	nameMap["澳门"] = "Macau"
	nameMap["内蒙古自治区"] = "Neimenggu"
	nameMap["宁夏回族自治区"] = "Ningxia"
	nameMap["青海省"] = "Qinghai"
	nameMap["陕西省"] = "Shaanxi"
	nameMap["山东省"] = "Shandong"
	nameMap["上海市"] = "Shanghai"
	nameMap["山西省"] = "Shanxi"
	nameMap["四川省"] = "Sichuan"
	nameMap["台湾省"] = "Taiwan"
	nameMap["天津市"] = "Tianjin"
	nameMap["新疆维吾尔自治区"] = "Xinjiang"
	nameMap["西藏自治区"] = "Xizang"
	nameMap["云南省"] = "Yunnan"
	nameMap["浙江省"] = "Zhejiang"

	provinceNamePinYin = nameMap[provinceName]
	return provinceNamePinYin
}
