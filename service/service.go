package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/TualatinX/durian-go/global"
	"github.com/TualatinX/durian-go/model"
	"gorm.io/gorm"
)

// 创建用户
func CreateAUser(user *model.User) (err error) {
	if err = global.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// 根据用户 ID 查询某个用户
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

// 根据用户 username 查询某个用户
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

// 更新用户的用户名、密码、个人信息
func UpdateAUser(user *model.User, username string, password string, userInfo string) error {
	user.Username = username
	user.Password = password
	user.UserInfo = userInfo
	err := global.DB.Save(user).Error
	return err
}

// 创建用户订阅城市
func CreateASubscription(userID uint64, cityName string) (err error) {
	subscription := model.Subscription{UserID: userID, Name: cityName}
	if err = global.DB.Create(&subscription).Error; err != nil {
		return err
	}
	return
}

// 根据订阅 ID 查询某个订阅情况
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

// 根据用户 ID 和其订阅城市名查询某个订阅情况
func QueryASubscriptionByUserIDAndCityName(userID uint64, cityName string) (subscription model.Subscription, notFound bool) {
	err := global.DB.Where("user_id = ? AND name = ?", userID, cityName).First(&subscription).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return subscription, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return subscription, false
	}
}

// 根据用户 ID 和其订阅城市名删除订阅城市
func DeleteASubscription(userID uint64, cityName string) (err error) {
	var subscription model.Subscription
	err = global.DB.Where("user_id = ? AND name = ?", userID, cityName).First(&subscription).Error
	_ = global.DB.Delete(&subscription).Error
	return err
}

// 查询某用户的所有城市
func QueryAllSubscriptions(userID uint64) (subscriptions []model.Subscription) {
	global.DB.Where("user_id = ?", userID).Find(&subscriptions)
	return subscriptions
}

// 查看订阅的城市（实际上是省和直辖市）的信息，如新增确诊、累计确诊、累计死亡、累计治愈
func QuerySubcriptionsData(subscriptions []model.Subscription) (subscriptionsData []model.CovidDetailCDRProvince, length int) {
	var chinaCases []model.CovidChinaCases
	var chinaDeaths []model.CovidChinaDeaths
	var chinaRecovered []model.CovidChinaRecovered

	err1 := global.DB.Order("date desc, province_name asc").Find(&chinaCases).Error
	err2 := global.DB.Order("date desc, province_name asc").Find(&chinaDeaths).Error
	err3 := global.DB.Order("date desc, province_name asc").Find(&chinaRecovered).Error

	if (err1 != nil && errors.Is(err1, gorm.ErrRecordNotFound)) || (err2 != nil && errors.Is(err2, gorm.ErrRecordNotFound)) || (err3 != nil && errors.Is(err3, gorm.ErrRecordNotFound)) {
		return subscriptionsData, 0
	} else if err1 != nil && !errors.Is(err1, gorm.ErrRecordNotFound) {
		panic(err1)
	} else if err2 != nil && !errors.Is(err2, gorm.ErrRecordNotFound) {
		panic(err2)
	} else if err3 != nil && !errors.Is(err3, gorm.ErrRecordNotFound) {
		panic(err3)
	} else {
		lenSubcriptions := len(subscriptions)
		// 循环获取订阅的省或直辖市的信息
		for i := 0; i < lenSubcriptions; i++ {
			provinceNamePinYin := global.MapPinYin(subscriptions[i].Name)
			j := 0
			for ; j < 34; j++ { // 查询省或直辖市相应的下标
				if chinaCases[j].ProvinceName == provinceNamePinYin {
					break
				}
			}
			fmt.Println(subscriptions[i].Name)
			fmt.Println(chinaCases[j].ProvinceName)
			subscriptionsData = append(subscriptionsData, model.CovidDetailCDRProvince{ProvinceName: subscriptions[i].Name,
				NowCases:     chinaCases[j].Info - chinaDeaths[j].Info - chinaRecovered[j].Info,
				Cases:        chinaCases[j].Info,
				NewCases:     chinaCases[j].Info - chinaCases[j+34].Info,
				Deaths:       chinaDeaths[j].Info,
				NewDeaths:    chinaDeaths[j].Info - chinaDeaths[j+34].Info,
				Recovered:    chinaRecovered[j].Info,
				NewRecovered: chinaRecovered[j].Info - chinaRecovered[j+34].Info,
				Vaccine:      0, NewVaccine: 0})
		}
		return subscriptionsData, len(subscriptionsData)
	}
}

// 根据辟谣 ID 查询新闻详情
func QueryARumorByID(rumorID uint64) (rumor model.Rumor, notFound bool) {
	err := global.DB.Where("rumor_id = ?", rumorID).First(&rumor).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return rumor, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return rumor, false
	}
}

// 查询所有辟谣
func QueryAllRumor() (rumor []model.Rumor) {
	global.DB.Find(&rumor)
	return rumor
}

// 根据新闻 ID 查询新闻详情
func QueryANewsByID(NewsID uint64) (news model.News, notFound bool) {
	err := global.DB.Where("news_id = ?", NewsID).First(&news).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return news, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return news, false
	}
}

// 查询所有新闻
func QueryAllNews() (news []model.News) {
	global.DB.Find(&news)
	return news
}

// 根据公告 ID 查询公告详情
func QueryANoticeByID(NoticeID uint64) (notice model.Notice, notFound bool) {
	err := global.DB.Where("notice_id = ?", NoticeID).First(&notice).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return notice, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return notice, false
	}
}

// 查询所有公告
func QueryAllNotice() (notice []model.Notice) {
	global.DB.Order("notice_created_time desc").Find(&notice)
	return notice
}

// 根据防疫小知识 ID 查询知识详情
func QueryAKnowledgeByID(KnowledgeID uint64) (knowledge model.Knowledge, notFound bool) {
	err := global.DB.Where("knowledge_id = ?", KnowledgeID).First(&knowledge).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return knowledge, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return knowledge, false
	}
}

// 查询所有防疫小知识
func QueryAllKnowledge() (knowledge []model.Knowledge) {
	global.DB.Find(&knowledge)
	return knowledge
}

// 创建一个知识版块下的问题
func CreateAQuestion(question *model.Question) (err error) {
	if err = global.DB.Create(&question).Error; err != nil {
		return err
	}
	return
}

// 根据问题 ID 查询一个问题
func QueryAQuestionByID(questionID uint64) (QuestionWithUserInfo model.QuestionWithUserInfo, notFound bool) {
	var question model.Question
	err := global.DB.Where("question_id = ?", questionID).First(&question).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return QuestionWithUserInfo, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		user, _ := QueryAUserByID(question.UserID)
		QuestionWithUserInfo = model.QuestionWithUserInfo{Question: question, Username: user.Username, UserInfo: user.UserInfo}
		return QuestionWithUserInfo, false
	}
}

// 查询所有问题
func QueryAllQuestions() (res []model.QuestionWithUserInfo, recommendQuestions []model.QuestionWithUserInfo) {
	var questions []model.Question
	global.DB.Order("question_time desc").Find(&questions)
	for _, e := range questions {
		user, _ := QueryAUserByID(e.UserID)
		res = append(res, model.QuestionWithUserInfo{Question: e, Username: user.Username, UserInfo: user.UserInfo})
	}
	mapSelected := make(map[int]bool)
	lenQuestions := len(questions)
	for i := 0; i < 7; i++ {
		mapSelected[int(rand.Intn(lenQuestions))] = true
	}
	for i, _ := range mapSelected {
		recommendQuestions = append(recommendQuestions, res[i])
	}
	return res, recommendQuestions
}

// 创建一个对问题的评论
func CreateAComment(comment *model.Comment) (err error) {
	if err = global.DB.Create(&comment).Error; err != nil {
		return err
	}
	return
}

// 列出某个问题的所有评论
func QueryAllComments(questionID uint64) (resWithUsername []model.CommentWithUserInfo) {
	var comments []model.Comment
	var res []model.Comment
	global.DB.Where("question_id = ?", questionID).Order("comment_time desc").Find(&comments)
	// 将置顶的评论提前，存入res
	for _, e := range comments {
		if e.UserType == 1 {
			res = append(res, e)
		}
	}
	for _, e := range comments {
		if e.UserType != 1 {
			res = append(res, e)
		}
	}
	for _, e := range res {
		user, _ := QueryAUserByID(e.UserID)
		resWithUsername = append(resWithUsername, model.CommentWithUserInfo{Comment: e, Username: user.Username, UserInfo: user.UserInfo})
	}
	return resWithUsername
}

// 查看所有国内航班
func QueryAllFlights() (flightDomesticWithStatus []model.FlightDomesticWithStatus) {
	var flights []model.FlightDomestic
	global.DB.Find(&flights)
	for _, v := range flights {
		flightDomesticWithStatus = append(flightDomesticWithStatus, model.FlightDomesticWithStatus{FlightDomestic: v, Status: "已取消"})
	}
	return flightDomesticWithStatus
}

// 查看所有国内列车
func QueryAllTrains() (trainDomesticWithStatus []model.TrainDomesticWithStatus) {
	var trains []model.TrainDomestic
	global.DB.Find(&trains)
	for _, v := range trains {
		trainDomesticWithStatus = append(trainDomesticWithStatus, model.TrainDomesticWithStatus{TrainDomestic: v, Status: "预计准点"})
	}
	return trainDomesticWithStatus
}

// 查询所有主要城市
func QueryAllMainCity() (city []model.MainCity) {
	global.DB.Find(&city)
	return city
}

// 查询城市的中心位置坐标
func QueryCenterCityCoordinate(name string) (city model.CenterCity, notFound bool) {
	err := global.DB.Where("name LIKE ?", "%"+name+"%").First(&city).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return city, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return city, false
	}
}

//----------------------------------------------------
//----------------------------------------------------
//------------------------新冠数据---------------------
//----------------------------------------------------
//----------------------------------------------------

// 查询所有高风险地区
func QueryAllHighRiskAreas() (areas []model.HighRiskArea) {
	global.DB.Find(&areas)
	return areas
}

// 根据 name 返回对应的数据 content (Json)
func QueryDataByName(name string) (directData model.DirectData, notFound bool) {
	err := global.DB.Where("name = ?", name).First(&directData).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return directData, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return directData, false
	}
}

// 查询所有地区的新冠感染人数
func QueryAllCovidCases() (cases []model.CovidCases) {
	global.DB.Find(&cases)
	return cases
}

// 查询所有地区的新冠感染人数（根据日期汇总）
func QueryAllCovidCasesResponse() (response []model.CovidCasesResponse) {
	var cases []model.CovidCases
	global.DB.Order("date asc").Find(&cases)
	lenCases := len(cases)
	if lenCases == 0 {
		return response
	}
	curDate := cases[0].Date

	for i := 0; i < lenCases; i++ {
		var tmp []model.CovidCasesNoDate
		for j := i; j < lenCases; j++ {
			if cases[j].Date == curDate {
				tmp = append(tmp, model.CovidCasesNoDate{CountryName: cases[j].CountryName, Info: cases[j].Info})
			} else {
				curDate = cases[j].Date
				i = j
				break
			}
			if j == lenCases-1 {
				i = lenCases
			}
		}
		response = append(response, model.CovidCasesResponse{Date: curDate, Value: tmp})
	}
	return response
}

// 查询所有地区的新冠死亡人数
func QueryAllCovidDeaths() (deaths []model.CovidDeaths) {
	global.DB.Find(&deaths)
	return deaths
}

// 查询所有地区的新冠死亡人数（根据日期汇总）
func QueryAllCovidDeathsResponse() (response []model.CovidDeathsResponse) {
	var cases []model.CovidDeaths
	global.DB.Order("date asc").Find(&cases)
	lenCases := len(cases)
	if lenCases == 0 {
		return response
	}
	curDate := cases[0].Date

	for i := 0; i < lenCases; i++ {
		var tmp []model.CovidDeathsNoDate
		for j := i; j < lenCases; j++ {
			if cases[j].Date == curDate {
				tmp = append(tmp, model.CovidDeathsNoDate{CountryName: cases[j].CountryName, Info: cases[j].Info})
			} else {
				curDate = cases[j].Date
				i = j
				break
			}
			if j == lenCases-1 {
				i = lenCases
			}
		}
		response = append(response, model.CovidDeathsResponse{Date: curDate, Value: tmp})
	}
	return response
}

// 查询所有地区的新冠治愈人数
func QueryAllCovidRecovereds() (recovereds []model.CovidRecovered) {
	global.DB.Find(&recovereds)
	return recovereds
}

// 查询所有地区的新冠治愈人数（根据日期汇总）
func QueryAllCovidRecoveredsResponse() (response []model.CovidRecoveredResponse) {
	var cases []model.CovidRecovered
	global.DB.Order("date asc").Find(&cases)
	lenCases := len(cases)
	if lenCases == 0 {
		return response
	}
	curDate := cases[0].Date

	for i := 0; i < lenCases; i++ {
		var tmp []model.CovidRecoveredNoDate
		for j := i; j < lenCases; j++ {
			if cases[j].Date == curDate {
				tmp = append(tmp, model.CovidRecoveredNoDate{CountryName: cases[j].CountryName, Info: cases[j].Info})
			} else {
				curDate = cases[j].Date
				i = j
				break
			}
			if j == lenCases-1 {
				i = lenCases
			}
		}
		response = append(response, model.CovidRecoveredResponse{Date: curDate, Value: tmp})
	}
	return response
}

// 查询所有地区的新冠疫苗接种人数
func QueryAllCovidVaccines() (vaccines []model.CovidVaccine) {
	global.DB.Find(&vaccines)
	return vaccines
}

// 查询所有地区的新冠疫苗接种人数（根据日期汇总）
func QueryAllCovidVaccinesResponse() (response []model.CovidVaccineResponse) {
	var cases []model.CovidVaccine
	global.DB.Order("date asc").Find(&cases)
	lenCases := len(cases)
	if lenCases == 0 {
		return response
	}
	curDate := cases[0].Date

	for i := 0; i < lenCases; i++ {
		var tmp []model.CovidVaccineNoDate
		for j := i; j < lenCases; j++ {
			if cases[j].Date == curDate {
				tmp = append(tmp, model.CovidVaccineNoDate{CountryName: cases[j].CountryName, Info: cases[j].Info})
			} else {
				curDate = cases[j].Date
				i = j
				break
			}
			if j == lenCases-1 {
				i = lenCases
			}
		}
		response = append(response, model.CovidVaccineResponse{Date: curDate, Value: tmp})
	}
	return response
}

// 查询某地区的新冠感染人数（根据日期汇总） [Province]
func QueryAllCovidCasesResponseProvince(province string) (response []model.CovidProvinceCases, notFound bool) {
	err := global.DB.Where("country_name = ?", province).Order("date asc").Find(&response).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return response, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return response, false
	}
}

// 查询某地区的新冠死亡人数（根据日期汇总） [Province]
func QueryAllCovidDeathsResponseProvince(province string) (response []model.CovidProvinceDeaths, notFound bool) {
	err := global.DB.Where("country_name = ?", province).Order("date asc").Find(&response).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return response, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return response, false
	}
}

// 查询某地区的新冠治愈人数（根据日期汇总） [Province]
func QueryAllCovidRecoveredsResponseProvince(province string) (response []model.CovidProvinceRecovered, notFound bool) {
	err := global.DB.Where("country_name = ?", province).Order("date asc").Find(&response).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return response, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return response, false
	}
}

// // 仅用于接口测试，测试二维数组
// func QueryChinaTest() (detail [][]model.CovidChinaCasesNoDate) {
// 	var cases []model.CovidChinaCases

// 	_ = global.DB.Order("date desc, province_name asc").Find(&cases).Error
// 	lenCases := len(cases)

// 	days := lenCases / 34         // 现共有534天
// 	for i := 0; i < days-1; i++ { // 最前一天（2020.1.22）没有新增，先不要了
// 		var temp []model.CovidChinaCasesNoDate
// 		for j := 0; j < 34; j++ {
// 			temp = append(temp, model.CovidChinaCasesNoDate{ProvinceName: cases[i*34+j].ProvinceName, Info: cases[i*34+j].Info})
// 		}
// 		detail = append(detail, temp)
// 	}
// 	return detail
// }

// 获取某个三级行政单位（如海淀区）的overview数据
func QueryDistrictOverview(districtName string) (districtData []model.DistrictOverview) {
	var districtCases []model.CovidHangzhouCases
	var districtDeaths []model.CovidHangzhouDeaths
	var districtRecovered []model.CovidHangzhouRecovered
	err1 := global.DB.Where("city_name = ?", districtName).Order("date asc").Find(&districtCases).Error
	err2 := global.DB.Where("city_name = ?", districtName).Order("date asc").Find(&districtDeaths).Error
	err3 := global.DB.Where("city_name = ?", districtName).Order("date asc").Find(&districtRecovered).Error

	if (err1 != nil && errors.Is(err1, gorm.ErrRecordNotFound)) || (err2 != nil && errors.Is(err2, gorm.ErrRecordNotFound)) || (err3 != nil && errors.Is(err3, gorm.ErrRecordNotFound)) {
		return districtData
	} else if err1 != nil && !errors.Is(err1, gorm.ErrRecordNotFound) {
		panic(err1)
	} else if err2 != nil && !errors.Is(err2, gorm.ErrRecordNotFound) {
		panic(err2)
	} else if err3 != nil && !errors.Is(err3, gorm.ErrRecordNotFound) {
		panic(err3)
	} else {
		lenCases := len(districtCases)
		lenDeaths := len(districtDeaths)
		lenRecovered := len(districtRecovered)
		length := lenCases
		if length < lenDeaths {
			length = lenDeaths
		}
		if length < lenRecovered { // 至此找到最短的表
			length = lenRecovered
		}
		for i := 0; i < length; i++ { // 等价于日期
			var casesNum uint64
			var casesNewNum uint64
			var casesNowNum uint64
			var deathsNum uint64
			var deathsNewNum uint64
			var recoveredNum uint64
			var recoveredNewNum uint64

			casesNum = districtCases[i].Info
			deathsNum = districtDeaths[i].Info
			recoveredNum = districtRecovered[i].Info
			casesNowNum = casesNum - deathsNum - recoveredNum
			curDate := districtCases[i].Date
			if i != 0 { // 第二天及以后
				casesNewNum = casesNum - districtData[i-1].Overview.Cases.NowNum
				deathsNewNum = deathsNum - districtData[i-1].Overview.Deaths.NowNum
				recoveredNewNum = recoveredNum - districtData[i-1].Overview.Recovered.NowNum
			} else {
				casesNewNum = 0
				deathsNewNum = 0
				recoveredNewNum = 0
			}

			nowCasesItem := model.NowCases{NowNum: casesNowNum, NewNum: casesNewNum}
			casesItem := model.Cases{NowNum: casesNum, NewNum: casesNewNum}
			deathItem := model.Deaths{NowNum: deathsNum, NewNum: deathsNewNum}
			recoveredItem := model.Recovered{NowNum: recoveredNum, NewNum: recoveredNewNum}
			vaccineItem := model.Vaccine{NowNum: 0, NewNum: 0}
			overviewItem := model.Overview{NowCases: nowCasesItem, Cases: casesItem, Deaths: deathItem, Vaccine: vaccineItem, Recovered: recoveredItem}
			districtOverviewItem := model.DistrictOverview{Date: curDate, Overview: overviewItem}
			districtData = append(districtData, districtOverviewItem)
		}
		return districtData
	}
}

// 获取中国某省份按日期构成的数据，每日数据下包含整体数据、省下的各市的数据
func QueryProvinceOverviewAndDetails(provinceName string, provinceNamePinYin string) (districtData []model.DistrictOverviewAndDetail) {
	var cases []model.CovidHangzhouCases
	var deaths []model.CovidHangzhouDeaths
	var recovered []model.CovidHangzhouRecovered

	// 用中文的省份名字进行查询
	err1 := global.DB.Where("province_name = ?", provinceName).Order("date asc, city_name asc").Find(&cases).Error
	err2 := global.DB.Where("province_name = ?", provinceName).Order("date asc, city_name asc").Find(&deaths).Error
	err3 := global.DB.Where("province_name = ?", provinceName).Order("date asc, city_name asc").Find(&recovered).Error

	if (err1 != nil && errors.Is(err1, gorm.ErrRecordNotFound)) || (err2 != nil && errors.Is(err2, gorm.ErrRecordNotFound)) || (err3 != nil && errors.Is(err3, gorm.ErrRecordNotFound)) {
		return districtData
	} else if err1 != nil && !errors.Is(err1, gorm.ErrRecordNotFound) {
		panic(err1)
	} else if err2 != nil && !errors.Is(err2, gorm.ErrRecordNotFound) {
		panic(err2)
	} else if err3 != nil && !errors.Is(err3, gorm.ErrRecordNotFound) {
		panic(err3)
	} else {
		// 不能简单用各市之和作为省份的总体数据，因此需要从省份的表格中查询
		var caseProvince []model.CovidChinaCases
		var deathsProvince []model.CovidChinaDeaths
		var recoveredProvince []model.CovidChinaRecovered

		// 用拼音进行查询
		_ = global.DB.Where("province_name = ?", provinceNamePinYin).Order("date asc").Find(&caseProvince).Error
		_ = global.DB.Where("province_name = ?", provinceNamePinYin).Order("date asc").Find(&deathsProvince).Error
		_ = global.DB.Where("province_name = ?", provinceNamePinYin).Order("date asc").Find(&recoveredProvince).Error
		provinceDateLength := len(caseProvince)

		lenCases := len(cases)
		districtLength := 0 // 查看这个省份/直辖市下有多少个市/区
		oneDate := cases[0].Date
		for i := 0; i < lenCases; i++ {
			if oneDate == cases[i].Date {
				districtLength += 1
			} else {
				break
			}
		}

		days := lenCases / districtLength // 共有多少天
		for i := 0; i < days; i++ {
			// 某一天的省份整体数据
			var detail []model.CovidDetailCDRDistrict
			var casesNum uint64
			var casesNewNum uint64
			var casesNowNum uint64
			var deathsNum uint64
			var deathsNewNum uint64
			var recoveredNum uint64
			var recoveredNewNum uint64

			curDate := cases[i*districtLength].Date // 这一天
			k := 0
			for k = 0; k < provinceDateLength; k++ { // 查找这一天的全省累计确诊、累计死亡、累计治愈
				if caseProvince[k].Date == curDate {
					casesNum = caseProvince[k].Info
					deathsNum = deathsProvince[k].Info
					recoveredNum = recoveredProvince[k].Info
					break
				}
			}
			if k == provinceDateLength { // 没查到
				casesNum = 0
				deathsNum = 0
				recoveredNum = 0
			}

			if i != 0 { // 第二天及以后
				casesNewNum = casesNum - districtData[i-1].Overview.Cases.NowNum
				deathsNewNum = deathsNum - districtData[i-1].Overview.Deaths.NowNum
				recoveredNewNum = recoveredNum - districtData[i-1].Overview.Recovered.NowNum
			} else {
				casesNewNum = 0
				deathsNewNum = 0
				recoveredNewNum = 0
			}
			casesNowNum = casesNum - deathsNum - recoveredNum

			nowCasesItem := model.NowCases{NowNum: casesNowNum, NewNum: casesNewNum}
			casesItem := model.Cases{NowNum: casesNum, NewNum: casesNewNum}
			deathItem := model.Deaths{NowNum: deathsNum, NewNum: deathsNewNum}
			recoveredItem := model.Recovered{NowNum: recoveredNum, NewNum: recoveredNewNum}
			vaccineItem := model.Vaccine{NowNum: 0, NewNum: 0}
			overviewItem := model.Overview{NowCases: nowCasesItem, Cases: casesItem, Deaths: deathItem, Vaccine: vaccineItem, Recovered: recoveredItem}

			// 每一天省份下的各市的数据
			for j := 0; j < districtLength; j++ {
				deathsDistrictNum := deaths[i*districtLength+j].Info
				recoveredDistrictNum := recovered[i*districtLength+j].Info
				casesDistrictNum := cases[i*districtLength+j].Info
				casesDistrictNowNum := casesDistrictNum - deathsDistrictNum - recoveredDistrictNum
				var casesDistrictNewNum uint64
				var deathsDistrictNewNum uint64
				var recoveredDistrictredNewNum uint64
				if i != 0 { // 第二天及以后
					casesDistrictNewNum = casesDistrictNum - cases[(i-1)*districtLength+j].Info
					deathsDistrictNewNum = deathsDistrictNum - deaths[(i-1)*districtLength+j].Info
					recoveredDistrictredNewNum = recoveredDistrictNum - recovered[(i-1)*districtLength+j].Info
				} else {
					casesDistrictNewNum = 0
					deathsDistrictNewNum = 0
					recoveredDistrictredNewNum = 0
				}
				detail = append(detail, model.CovidDetailCDRDistrict{DistrictName: cases[i*districtLength+j].CityName,
					NowCases:     casesDistrictNowNum,
					Cases:        casesDistrictNum,
					NewCases:     casesDistrictNewNum,
					Deaths:       deathsDistrictNum,
					NewDeaths:    deathsDistrictNewNum,
					Recovered:    recoveredDistrictNum,
					NewRecovered: recoveredDistrictredNewNum,
					Vaccine:      0,
					NewVaccine:   0})
			}
			districtOverviewAndDetailsItem := model.DistrictOverviewAndDetail{Date: curDate, Overview: overviewItem, Detailed: detail}

			// 开始汇总
			districtData = append(districtData, districtOverviewAndDetailsItem)
		}
		return districtData
	}
}

// 获取某个外国国家按日期构成的数据，每日数据下包含整体数据、各省数据
func QueryOtherCountryOverviewAndDetails(countryName string) (data []model.CountryOverviewAndDetails) {
	var countryCases []model.CovidProvinceCases
	var countryDeaths []model.CovidProvinceDeaths
	var countryRecovered []model.CovidProvinceRecovered
	// 获取各个省份的信息
	_ = global.DB.Where("country_name = ?", countryName).Order("date asc").Find(&countryCases).Error
	_ = global.DB.Where("country_name = ?", countryName).Order("date asc").Find(&countryDeaths).Error
	_ = global.DB.Where("country_name = ?", countryName).Order("date asc").Find(&countryRecovered).Error
	lenCases := len(countryCases)
	lenDeaths := len(countryDeaths)
	lenRecovered := len(countryRecovered)
	if (lenCases != lenDeaths) || (lenCases != lenRecovered) || (lenDeaths != lenRecovered) {
		fmt.Print(lenCases)
		return data
	}

	// 获取国家每天的整体信息：累计死亡、累计确诊、累计治愈
	var caseCountry []model.CovidCases
	var deathsCountry []model.CovidDeaths
	var recoveredCountry []model.CovidRecovered
	_ = global.DB.Where("country_name = ?", countryName).Order("date asc").Find(&caseCountry).Error
	_ = global.DB.Where("country_name = ?", countryName).Order("date asc").Find(&deathsCountry).Error
	_ = global.DB.Where("country_name = ?", countryName).Order("date asc").Find(&recoveredCountry).Error
	dateLength := len(caseCountry)
	fmt.Println(lenCases)
	fmt.Println(dateLength)

	for i := 0; i < lenCases; i++ { // 相当于日期
		var detail []model.CovidCDRProvince
		var casesNum uint64    // 今日累计确诊
		var casesNewNum uint64 // 今日新增确诊
		var casesNowNum uint64 // 今日现存确诊
		var deathsNum uint64
		var deathsNewNum uint64
		var recoveredNum uint64
		var recoveredNewNum uint64

		// 解析第i天的数据
		var casesMap map[string]uint64
		var deathsMap map[string]uint64
		var recoveredMap map[string]uint64
		json.Unmarshal([]byte(countryCases[i].Info), &casesMap)
		json.Unmarshal([]byte(countryDeaths[i].Info), &deathsMap)
		json.Unmarshal([]byte(countryRecovered[i].Info), &recoveredMap)

		for provinceName, casesProvinceNum := range casesMap {
			deathsProvinceNum := deathsMap[provinceName]
			recoveredProvinceNum := recoveredMap[provinceName]
			casesProvinceNowNum := casesProvinceNum - deathsProvinceNum - recoveredProvinceNum
			detail = append(detail, model.CovidCDRProvince{ProvinceName: provinceName, NowCases: casesProvinceNowNum,
				Cases: casesProvinceNum, Deaths: deathsProvinceNum, Recovered: recoveredProvinceNum})

		}
		curDate := countryCases[i].Date
		// _ = global.DB.Where("country_name = ? AND date = ?", countryName, curDate).First(&caseCountry).Error
		// _ = global.DB.Where("country_name = ? AND date = ?", countryName, curDate).First(&deathsCountry).Error
		// _ = global.DB.Where("country_name = ? AND date = ?", countryName, curDate).First(&recoveredCountry).Error
		k := 0
		for k = 0; k < dateLength; k++ { // 查找这一天的全国累计确诊、累计死亡、累计治愈
			if caseCountry[k].Date == curDate {
				casesNum = caseCountry[k].Info
				deathsNum = deathsCountry[k].Info
				recoveredNum = recoveredCountry[k].Info
				break
			}
		}
		if k == dateLength { // 没查到
			casesNum = 0
			deathsNum = 0
			recoveredNum = 0
		}

		casesNowNum = casesNum - deathsNum - recoveredNum // 现存确诊=今日累计确诊-今日累计死亡-今日累计治愈
		if i != 0 {                                       // 第二天及以后
			casesNewNum = casesNum - data[i-1].Overview.Cases.NowNum // 新增确诊=今日累计确诊-昨日累计确诊
			deathsNewNum = deathsNum - data[i-1].Overview.Deaths.NowNum
			recoveredNewNum = recoveredNum - data[i-1].Overview.Recovered.NowNum
		} else { // 第一天就认为全是0了
			casesNewNum = 0
			deathsNewNum = 0
			recoveredNewNum = 0
		}

		nowCasesItem := model.NowCases{NowNum: casesNowNum, NewNum: casesNewNum}
		casesItem := model.Cases{NowNum: casesNum, NewNum: casesNewNum}
		deathItem := model.Deaths{NowNum: deathsNum, NewNum: deathsNewNum}
		recoveredItem := model.Recovered{NowNum: recoveredNum, NewNum: recoveredNewNum}
		vaccineItem := model.Vaccine{NowNum: 0, NewNum: 0} // 先填0，后续有需求再添加
		overviewItem := model.Overview{NowCases: nowCasesItem, Cases: casesItem, Deaths: deathItem, Vaccine: vaccineItem, Recovered: recoveredItem}
		countryOverviewAndDetailsItem := model.CountryOverviewAndDetails{Date: curDate, Overview: overviewItem, Detailed: detail}
		data = append(data, countryOverviewAndDetailsItem)
	}

	return data
}

// 获取中国按日期构成的数据，每日数据下包含中国整体数据、各省数据
func QueryChinaOverviewAndDetails() (data []model.CountryOverviewAndDetails, notFound bool) {
	var cases []model.CovidChinaCases
	var deaths []model.CovidChinaDeaths
	var recovered []model.CovidChinaRecovered

	err1 := global.DB.Order("date asc, province_name asc").Find(&cases).Error
	err2 := global.DB.Order("date asc, province_name asc").Find(&deaths).Error
	err3 := global.DB.Order("date asc, province_name asc").Find(&recovered).Error

	if (err1 != nil && errors.Is(err1, gorm.ErrRecordNotFound)) || (err2 != nil && errors.Is(err2, gorm.ErrRecordNotFound)) || (err3 != nil && errors.Is(err3, gorm.ErrRecordNotFound)) {
		return data, true
	} else if err1 != nil && !errors.Is(err1, gorm.ErrRecordNotFound) {
		panic(err1)
	} else if err2 != nil && !errors.Is(err2, gorm.ErrRecordNotFound) {
		panic(err2)
	} else if err3 != nil && !errors.Is(err3, gorm.ErrRecordNotFound) {
		panic(err3)
	} else {
		lenCases := len(cases)
		lenDeaths := len(deaths)
		lenRecovered := len(recovered)
		if (lenCases != lenDeaths) || (lenCases != lenRecovered) || (lenDeaths != lenRecovered) {
			fmt.Print(lenCases)
			return data, true
		}
		days := lenCases / 34 // 现共有535天
		for i := 0; i < days; i++ {
			var detail []model.CovidCDRProvince
			var casesNum uint64
			var casesNewNum uint64
			var casesNowNum uint64
			var deathsNum uint64
			var deathsNewNum uint64
			var recoveredNum uint64
			var recoveredNewNum uint64
			// var vaccineNum uint64
			// var vaccineNewNum uint64
			for j := 0; j < 34; j++ {
				detail = append(detail, model.CovidCDRProvince{ProvinceName: cases[i*34+j].ProvinceName,
					NowCases:  cases[i*34+j].Info - deaths[i*34+j].Info - recovered[i*34+j].Info,
					Cases:     cases[i*34+j].Info,
					Deaths:    deaths[i*34+j].Info,
					Recovered: recovered[i*34+j].Info})
				casesNum += cases[i*34+j].Info
				deathsNum += deaths[i*34+j].Info
				recoveredNum += recovered[i*34+j].Info
			}

			casesNowNum = casesNum - deathsNum - recoveredNum
			curDate := cases[i*34].Date
			// var chinaVaccineToday model.CovidVaccine
			// err := global.DB.Where("country_name = ? AND date = ?", "China", curDate).First(&chinaVaccineToday).Error
			if i != 0 { // 第二天及以后
				casesNewNum = casesNum - data[i-1].Overview.Cases.NowNum
				deathsNewNum = deathsNum - data[i-1].Overview.Deaths.NowNum
				recoveredNewNum = recoveredNum - data[i-1].Overview.Recovered.NowNum
				// if err != nil && errors.Is(err, gorm.ErrRecordNotFound) { // covid_vaccine最新一天(7.9)的一定查不到，用前一天(7.8)的代替
				// 	fmt.Println(curDate)
				// 	fmt.Println(data[i-1].Overview.Vaccine.NowNum)
				// 	vaccineNum = data[i-1].Overview.Vaccine.NowNum
				// 	vaccineNewNum = data[i-1].Overview.Vaccine.NewNum
				// } else {
				// 	vaccineNum = chinaVaccineToday.Info
				// 	vaccineNewNum = chinaVaccineToday.Info - data[i-1].Overview.Vaccine.NowNum
				// }
			} else {
				casesNewNum = 0
				deathsNewNum = 0
				recoveredNewNum = 0
				// vaccineNewNum = 0
			}

			nowCasesItem := model.NowCases{NowNum: casesNowNum, NewNum: casesNewNum}
			casesItem := model.Cases{NowNum: casesNum, NewNum: casesNewNum}
			deathItem := model.Deaths{NowNum: deathsNum, NewNum: deathsNewNum}
			recoveredItem := model.Recovered{NowNum: recoveredNum, NewNum: recoveredNewNum}
			vaccineItem := model.Vaccine{NowNum: 0, NewNum: 0}
			overviewItem := model.Overview{NowCases: nowCasesItem, Cases: casesItem, Deaths: deathItem, Vaccine: vaccineItem, Recovered: recoveredItem}
			countryOverviewAndDetailsItem := model.CountryOverviewAndDetails{Date: curDate, Overview: overviewItem, Detailed: detail}
			data = append(data, countryOverviewAndDetailsItem)
		}
		return data, false
	}
}

// HomeData.json
// 获取全球整体数据和各国的当日的详细数据，全球整体数据为全球确诊等，各国详细数据下包含新增确诊、累计确诊、现存确诊等
func QueryOtherCountryOverviewAndDetailsForHomeData() (globalTable model.GlobalOverviewAndDetails, chinaTable model.ChinaOverviewAndDetails) {
	var cases []model.CovidCases
	var deaths []model.CovidDeaths
	var recovered []model.CovidRecovered

	// 获取各个国家的信息
	_ = global.DB.Order("date desc, country_name asc").Find(&cases).Error
	_ = global.DB.Order("date desc, country_name asc").Find(&deaths).Error
	_ = global.DB.Order("date desc, country_name asc").Find(&recovered).Error

	var globalDetail []model.CovidDetailCDRCountry
	var globalCasesNum uint64    // 今日全球累计确诊
	var globalCasesNewNum uint64 // 今日全球新增确诊
	var globalCasesNowNum uint64 // 今日全球现存确诊
	var globalDeathsNum uint64
	var globalDeathsNewNum uint64
	var globalRecoveredNum uint64
	var globalRecoveredNewNum uint64

	var chinaDetail []model.CovidDetailCDRProvince
	var chinaCasesNum uint64
	var chinaCasesNowNum uint64
	var chinaCasesNewNum uint64
	var chinaDeathsNum uint64
	var chinaDeathsNewNum uint64
	var chinaRecoveredNum uint64
	var chinaRecoveredNewNum uint64

	// 对世界各国进行统计
	for i := 0; i < 173; i++ { // 国家数目
		var casesNum uint64    // 今日累计确诊
		var casesNewNum uint64 // 今日新增确诊
		var casesNowNum uint64 // 今日现存确诊
		var deathsNum uint64
		var deathsNewNum uint64
		var recoveredNum uint64
		var recoveredNewNum uint64

		if cases[i].CountryName != "Global" {
			casesNum = cases[i].Info
			deathsNum = deaths[i].Info
			recoveredNum = recovered[i].Info

			casesNowNum = casesNum - deathsNum - recoveredNum

			casesNewNum = casesNum - cases[i+173].Info
			deathsNewNum = deathsNum - deaths[i+173].Info
			recoveredNewNum = recoveredNum - recovered[i+173].Info

			globalDetail = append(globalDetail, model.CovidDetailCDRCountry{CountryName: cases[i].CountryName,
				NowCases: casesNowNum, Cases: casesNum, NewCases: casesNewNum,
				Deaths: deathsNum, NewDeaths: deathsNewNum,
				Recovered: recoveredNum, NewRecovered: recoveredNewNum,
				Vaccine: 0, NewVaccine: 0})
			if cases[i].CountryName == "China" {
				chinaCasesNum = cases[i].Info
				chinaDeathsNum = deaths[i].Info
				chinaRecoveredNum = recovered[i].Info

				chinaRecoveredNewNum = chinaRecoveredNum - recovered[i+173].Info
				chinaDeathsNewNum = chinaDeathsNum - deaths[i+173].Info
				chinaCasesNowNum = chinaCasesNum - chinaDeathsNum - chinaRecoveredNum
				chinaCasesNewNum = chinaCasesNum - cases[i+173].Info
			}
		} else {
			globalCasesNum = cases[i].Info
			globalDeathsNum = deaths[i].Info
			globalRecoveredNum = recovered[i].Info

			globalCasesNowNum = globalCasesNum - globalDeathsNum - globalRecoveredNum
			globalCasesNewNum = globalCasesNum - cases[i+173].Info
			globalDeathsNewNum = globalDeathsNum - deaths[i+173].Info
			globalRecoveredNewNum = globalRecoveredNum - recovered[i+173].Info
		}
	}
	nowGlobalCasesItem := model.NowCases{NowNum: globalCasesNowNum, NewNum: globalCasesNewNum}
	casesGlobalItem := model.Cases{NowNum: globalCasesNum, NewNum: globalCasesNewNum}
	deathGlobalItem := model.Deaths{NowNum: globalDeathsNum, NewNum: globalDeathsNewNum}
	recoveredGlobalItem := model.Recovered{NowNum: globalRecoveredNum, NewNum: globalRecoveredNewNum}
	vaccineGlobalItem := model.Vaccine{NowNum: 0, NewNum: 0} // 先填0，后续有需求再添加

	overviewGlobalItem := model.Overview{NowCases: nowGlobalCasesItem, Cases: casesGlobalItem, Deaths: deathGlobalItem, Vaccine: vaccineGlobalItem, Recovered: recoveredGlobalItem}
	globalTable = model.GlobalOverviewAndDetails{Overview: overviewGlobalItem, Detailed: globalDetail}

	// 中国部分
	var chinaCases []model.CovidChinaCases
	var chinaDeaths []model.CovidChinaDeaths
	var chinaRecovered []model.CovidChinaRecovered

	_ = global.DB.Order("date desc, province_name asc").Find(&chinaCases).Error
	_ = global.DB.Order("date desc, province_name asc").Find(&chinaDeaths).Error
	_ = global.DB.Order("date desc, province_name asc").Find(&chinaRecovered).Error

	for i := 0; i < 34; i++ {
		chinaDetail = append(chinaDetail, model.CovidDetailCDRProvince{ProvinceName: chinaCases[i].ProvinceName,
			NowCases: chinaCases[i].Info - chinaDeaths[i].Info - chinaRecovered[i].Info, Cases: chinaCases[i].Info, NewCases: chinaCases[i].Info - chinaCases[i+34].Info,
			Deaths: chinaDeaths[i].Info, NewDeaths: chinaDeaths[i].Info - chinaDeaths[i+34].Info,
			Recovered: chinaRecovered[i].Info, NewRecovered: chinaRecovered[i].Info - chinaRecovered[i+34].Info,
			Vaccine: 0, NewVaccine: 0})
	}
	nowChinaCasesItem := model.NowCases{NowNum: chinaCasesNowNum, NewNum: chinaCasesNewNum}
	casesChinaItem := model.Cases{NowNum: chinaCasesNum, NewNum: chinaCasesNewNum}
	deathChinaItem := model.Deaths{NowNum: chinaDeathsNum, NewNum: chinaDeathsNewNum}
	recoveredChinaItem := model.Recovered{NowNum: chinaRecoveredNum, NewNum: chinaRecoveredNewNum}
	vaccineChinaItem := model.Vaccine{NowNum: 0, NewNum: 0} // 先填0，后续有需求再添加

	overviewChinaItem := model.Overview{NowCases: nowChinaCasesItem, Cases: casesChinaItem, Deaths: deathChinaItem, Vaccine: vaccineChinaItem, Recovered: recoveredChinaItem}
	chinaTable = model.ChinaOverviewAndDetails{Overview: overviewChinaItem, Detailed: chinaDetail}

	return globalTable, chinaTable
}

// 获取世界每天的overview数据和detail数据（每天的各国的数据）
func QueryGlobalOverviewAndDetailsHistory() (globalTable []model.GlobalOverviewAndDetailsWithDate) {
	var cases []model.CovidCases
	var deaths []model.CovidDeaths
	var recovered []model.CovidRecovered

	// 获取各个国家的信息
	_ = global.DB.Order("date asc, country_name asc").Find(&cases).Error
	_ = global.DB.Order("date asc, country_name asc").Find(&deaths).Error
	_ = global.DB.Order("date asc, country_name asc").Find(&recovered).Error

	countryLength := 0 // 查看有多少个国家
	oneDate := cases[0].Date
	lenCases := len(cases)
	for i := 0; i < lenCases; i++ {
		if oneDate == cases[i].Date {
			countryLength += 1
		} else {
			break
		}
	}
	dateLength := lenCases / countryLength // 一共有多少天 现在共有535天 173个国家
	// 对世界各国进行统计
	// 按照日期进行遍历
	for i := 0; i < dateLength; i++ {

		var globalDetail []model.CovidDetailCDRCountry // 记录今日全球的detail数据，也即每个国家的统计数据
		var globalCasesNum uint64                      // 今日全球累计确诊
		var globalCasesNewNum uint64                   // 今日全球新增确诊
		var globalCasesNowNum uint64                   // 今日全球现存确诊
		var globalDeathsNum uint64
		var globalDeathsNewNum uint64
		var globalRecoveredNum uint64
		var globalRecoveredNewNum uint64

		curDate := cases[i*countryLength].Date
		for j := 0; j < countryLength; j++ { // 国家数目
			var casesNum uint64    // 今日累计确诊
			var casesNewNum uint64 // 今日新增确诊
			var casesNowNum uint64 // 今日现存确诊
			var deathsNum uint64
			var deathsNewNum uint64
			var recoveredNum uint64
			var recoveredNewNum uint64

			if cases[i*countryLength+j].CountryName != "Global" {
				casesNum = cases[i*countryLength+j].Info
				deathsNum = deaths[i*countryLength+j].Info
				recoveredNum = recovered[i*countryLength+j].Info

				casesNowNum = casesNum - deathsNum - recoveredNum

				if i != 0 { // 减去前一天的
					casesNewNum = casesNum - cases[(i-1)*countryLength+j].Info
					deathsNewNum = deathsNum - deaths[(i-1)*countryLength+j].Info
					recoveredNewNum = recoveredNum - recovered[(i-1)*countryLength+j].Info
				} else {
					casesNewNum = 0
					deathsNewNum = 0
					recoveredNewNum = 0
				}

				globalDetail = append(globalDetail, model.CovidDetailCDRCountry{CountryName: cases[i*countryLength+j].CountryName,
					NowCases: casesNowNum, Cases: casesNum, NewCases: casesNewNum,
					Deaths: deathsNum, NewDeaths: deathsNewNum,
					Recovered: recoveredNum, NewRecovered: recoveredNewNum,
					Vaccine: 0, NewVaccine: 0})

			} else {
				globalCasesNum = cases[i*countryLength+j].Info
				globalDeathsNum = deaths[i*countryLength+j].Info
				globalRecoveredNum = recovered[i*countryLength+j].Info

				globalCasesNowNum = globalCasesNum - globalDeathsNum - globalRecoveredNum

				if i != 0 { // 直接减去前一天
					globalCasesNewNum = globalCasesNum - cases[(i-1)*countryLength+j].Info
					globalDeathsNewNum = globalDeathsNum - deaths[(i-1)*countryLength+j].Info
					globalRecoveredNewNum = globalRecoveredNum - recovered[(i-1)*countryLength+j].Info
				} else {
					globalCasesNewNum = 0
					globalDeathsNewNum = 0
					globalRecoveredNewNum = 0
				}
			}
		}
		nowGlobalCasesItem := model.NowCases{NowNum: globalCasesNowNum, NewNum: globalCasesNewNum}
		casesGlobalItem := model.Cases{NowNum: globalCasesNum, NewNum: globalCasesNewNum}
		deathGlobalItem := model.Deaths{NowNum: globalDeathsNum, NewNum: globalDeathsNewNum}
		recoveredGlobalItem := model.Recovered{NowNum: globalRecoveredNum, NewNum: globalRecoveredNewNum}
		vaccineGlobalItem := model.Vaccine{NowNum: 0, NewNum: 0} // 先填0，后续有需求再添加

		overviewGlobalItem := model.Overview{NowCases: nowGlobalCasesItem, Cases: casesGlobalItem, Deaths: deathGlobalItem, Vaccine: vaccineGlobalItem, Recovered: recoveredGlobalItem}
		globalTableItem := model.GlobalOverviewAndDetailsWithDate{Date: curDate, Overview: overviewGlobalItem, Detailed: globalDetail}
		globalTable = append(globalTable, globalTableItem)
	}
	return globalTable
}
