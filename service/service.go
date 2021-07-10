package service

import (
	"errors"

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
func UpdateAUser(user *model.User, username string, password string, info string) error {
	user.Username = username
	user.Password = password
	user.Info = info
	err := global.DB.Save(user).Error
	return err
}

// 创建用户订阅城市
func CreateASubscription(userID uint64, cityName string) (err error) {
	subscription := model.Subscription{UserID: userID, CityName: cityName}
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
	err := global.DB.Where("user_id = ? AND city_name = ?", userID, cityName).First(&subscription).Error
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
	err = global.DB.Where("user_id = ? AND city_name = ?", userID, cityName).First(&subscription).Error
	_ = global.DB.Delete(&subscription).Error
	return err
}

// 查询某用户的所有城市
func QueryAllSubscriptions(userID uint64) (subscriptions []model.Subscription) {
	global.DB.Where("user_id = ?", userID).Find(&subscriptions)
	return subscriptions
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

// 创建一个知识版块下的问题
func CreateAQuestion(question *model.Question) (err error) {
	if err = global.DB.Create(&question).Error; err != nil {
		return err
	}
	return
}

// 根据问题 ID 查询一个问题
func QueryAQuestionByID(questionID uint64) (questionWithUsername model.QuestionWithUsername, notFound bool) {
	var question model.Question
	err := global.DB.Where("question_id = ?", questionID).First(&question).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return questionWithUsername, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		user, _ := QueryAUserByID(question.UserID)
		questionWithUsername = model.QuestionWithUsername{Question: question, Username: user.Username}
		return questionWithUsername, false
	}
}

// 查询所有问题
func QueryAllQuestions() (res []model.QuestionWithUsername) {
	var questions []model.Question
	global.DB.Order("question_time desc").Find(&questions)
	for _, e := range questions {
		user, _ := QueryAUserByID(e.UserID)
		res = append(res, model.QuestionWithUsername{Question: e, Username: user.Username})
	}
	return res
}

// 创建一个对问题的评论
func CreateAComment(comment *model.Comment) (err error) {
	if err = global.DB.Create(&comment).Error; err != nil {
		return err
	}
	return
}

// 列出某个问题的所有评论
func QueryAllComments(questionID uint64) (resWithUsername []model.CommentWithUsername) {
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
		resWithUsername = append(resWithUsername, model.CommentWithUsername{Comment: e, Username: user.Username})
	}
	return resWithUsername
}

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

// 查询中国或世界的累计死亡数和新增死亡数
func QueryDeathOverview() (accumulativeDeaths []model.CovidDeathsNoDate, newDeaths []model.CovidDeathsNoDate, nums []int64, notFound bool) {
	var deaths []model.CovidDeaths
	err := global.DB.Order("date desc, country_name asc").Find(&deaths).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return accumulativeDeaths, newDeaths, nums, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		var globalDeathToday int64
		var globalDeathYesterday int64
		var chinaDeathToday int64
		var chinaDeathYesterday int64

		length := len(deaths)
		curDate := deaths[0].Date
		i := 0
		flag := 0

		globalDeathToday = 0
		globalDeathYesterday = 0
		chinaDeathToday = 0
		chinaDeathYesterday = 0

		for i = 0; i < length; i++ {
			if deaths[i].Date == curDate {
				if flag == 0 && deaths[i].CountryName == "Global" {
					globalDeathToday = int64(deaths[i].Info)
					flag += 1
					continue
				}
				accumulativeDeaths = append(accumulativeDeaths, model.CovidDeathsNoDate{CountryName: deaths[i].CountryName, Info: deaths[i].Info})
				if deaths[i].CountryName == "China" {
					chinaDeathToday = int64(deaths[i].Info)
				}

			} else {
				curDate = deaths[i].Date
				break
			}
		}
		lenAcc := len(accumulativeDeaths)
		for ; i < length; i++ {
			if deaths[i].Date == curDate {
				if flag == 1 && deaths[i].CountryName == "Global" {
					globalDeathYesterday = int64(deaths[i].Info)
					flag += 1
					continue
				}
				j := 0
				for j = 0; j < lenAcc; j++ {
					if accumulativeDeaths[j].CountryName == deaths[i].CountryName {
						newDeaths = append(newDeaths, model.CovidDeathsNoDate{CountryName: deaths[i].CountryName, Info: accumulativeDeaths[j].Info - deaths[i].Info})
						break
					}
				}
				if j == lenAcc { // 说明没找到这个国家在最新一天的累计死亡数，用0x3f3f3f3f标记
					newDeaths = append(newDeaths, model.CovidDeathsNoDate{CountryName: deaths[i].CountryName, Info: 0x3f3f3f3f})
				}
				if deaths[i].CountryName == "China" {
					chinaDeathYesterday = int64(deaths[i].Info)
				}
			} else {
				break
			}
		}
		nums = append(nums, globalDeathToday)
		nums = append(nums, globalDeathToday-globalDeathYesterday)
		nums = append(nums, chinaDeathToday)
		nums = append(nums, chinaDeathToday-chinaDeathYesterday)
		return accumulativeDeaths, newDeaths, nums, false
	}
}

// 查询中国或世界的累计治愈数和新增治愈数
func QueryRecoveredOverview() (accumulativeRecovered []model.CovidRecoveredNoDate, newRecovered []model.CovidRecoveredNoDate, nums []int64, notFound bool) {
	var recovered []model.CovidRecovered
	err := global.DB.Order("date desc, country_name asc").Find(&recovered).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return accumulativeRecovered, newRecovered, nums, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		var globalRecoveredToday int64
		var globalRecoveredYesterday int64
		var chinaRecoveredToday int64
		var chinaRecoveredYesterday int64

		length := len(recovered)
		curDate := recovered[0].Date
		i := 0
		flag := 0

		globalRecoveredToday = 0
		globalRecoveredYesterday = 0
		chinaRecoveredToday = 0
		chinaRecoveredYesterday = 0

		for i = 0; i < length; i++ {
			if recovered[i].Date == curDate {
				if flag == 0 && recovered[i].CountryName == "Global" {
					globalRecoveredToday = int64(recovered[i].Info)
					flag += 1
					continue
				}
				accumulativeRecovered = append(accumulativeRecovered, model.CovidRecoveredNoDate{CountryName: recovered[i].CountryName, Info: recovered[i].Info})
				if recovered[i].CountryName == "China" {
					chinaRecoveredToday = int64(recovered[i].Info)
				}

			} else {
				curDate = recovered[i].Date
				break
			}
		}
		lenAcc := len(accumulativeRecovered)
		for ; i < length; i++ {
			if recovered[i].Date == curDate {
				if flag == 1 && recovered[i].CountryName == "Global" {
					globalRecoveredYesterday = int64(recovered[i].Info)
					flag += 1
					continue
				}
				j := 0
				for j = 0; j < lenAcc; j++ {
					if accumulativeRecovered[j].CountryName == recovered[i].CountryName {
						newRecovered = append(newRecovered, model.CovidRecoveredNoDate{CountryName: recovered[i].CountryName, Info: accumulativeRecovered[j].Info - recovered[i].Info})
						break
					}
				}
				if j == lenAcc { // 说明没找到这个国家在最新一天的累计治愈数，用0x3f3f3f3f标记
					newRecovered = append(newRecovered, model.CovidRecoveredNoDate{CountryName: recovered[i].CountryName, Info: 0x3f3f3f3f})
				}
				if recovered[i].CountryName == "China" {
					chinaRecoveredYesterday = int64(recovered[i].Info)
				}
			} else {
				break
			}
		}
		nums = append(nums, globalRecoveredToday)
		nums = append(nums, globalRecoveredToday-globalRecoveredYesterday)
		nums = append(nums, chinaRecoveredToday)
		nums = append(nums, chinaRecoveredToday-chinaRecoveredYesterday)
		return accumulativeRecovered, newRecovered, nums, false
	}
}

// 查询中国或世界的累计接种数和新增接种数
func QueryVaccineOverview() (accumulativeVaccine []model.CovidVaccineNoDate, newVaccine []model.CovidVaccineNoDate, nums []int64, notFound bool) {
	var vaccine []model.CovidVaccine
	err := global.DB.Order("date desc, country_name asc").Find(&vaccine).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return accumulativeVaccine, newVaccine, nums, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		var globalVaccineToday int64
		var globalVaccineYesterday int64
		var chinaVaccineToday int64
		var chinaVaccineYesterday int64

		length := len(vaccine)
		curDate := vaccine[0].Date
		i := 0

		globalVaccineToday = 0
		globalVaccineYesterday = 0
		chinaVaccineToday = 0
		chinaVaccineYesterday = 0

		for i = 0; i < length; i++ {
			if vaccine[i].Date == curDate {
				if vaccine[i].CountryName == "Global" {
					globalVaccineToday = int64(vaccine[i].Info)
					continue
				}
				accumulativeVaccine = append(accumulativeVaccine, model.CovidVaccineNoDate{CountryName: vaccine[i].CountryName, Info: vaccine[i].Info})
				if vaccine[i].CountryName == "China" {
					chinaVaccineToday = int64(vaccine[i].Info)
				}

			} else {
				curDate = vaccine[i].Date
				break
			}
		}
		lenAcc := len(accumulativeVaccine)
		for ; i < length; i++ {
			if vaccine[i].Date == curDate {
				if vaccine[i].CountryName == "Global" {
					globalVaccineYesterday = int64(vaccine[i].Info)
					continue
				}
				j := 0
				for j = 0; j < lenAcc; j++ {
					if accumulativeVaccine[j].CountryName == vaccine[i].CountryName {
						newVaccine = append(newVaccine, model.CovidVaccineNoDate{CountryName: vaccine[i].CountryName, Info: accumulativeVaccine[j].Info - vaccine[i].Info})
						break
					}
				}
				if j == lenAcc { // 说明没找到这个国家在最新一天的累计接种数，用0x3f3f3f3f标记
					newVaccine = append(newVaccine, model.CovidVaccineNoDate{CountryName: vaccine[i].CountryName, Info: 0x3f3f3f3f})
				}
				if vaccine[i].CountryName == "China" {
					chinaVaccineYesterday = int64(vaccine[i].Info)
				}
			} else {
				break
			}
		}
		nums = append(nums, globalVaccineToday)
		nums = append(nums, globalVaccineToday-globalVaccineYesterday)
		nums = append(nums, chinaVaccineToday)
		nums = append(nums, chinaVaccineToday-chinaVaccineYesterday)
		return accumulativeVaccine, newVaccine, nums, false
	}
}

// 查询中国或世界的累计确诊数和新增确诊数
func QueryCasesOverview() (accumulativeCases []model.CovidCasesNoDate, newCases []model.CovidCasesNoDate, nums []int64, notFound bool) {
	var cases []model.CovidCases
	err := global.DB.Order("date desc, country_name asc").Find(&cases).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return accumulativeCases, newCases, nums, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		var globalCasesToday int64
		var globalCasesYesterday int64
		var chinaCasesToday int64
		var chinaCasesYesterday int64

		length := len(cases)
		curDate := cases[0].Date
		i := 0

		globalCasesToday = 0
		globalCasesYesterday = 0
		chinaCasesToday = 0
		chinaCasesYesterday = 0

		for i = 0; i < length; i++ {
			if cases[i].Date == curDate {
				if cases[i].CountryName == "Global" {
					globalCasesToday = int64(cases[i].Info)
					continue
				}
				accumulativeCases = append(accumulativeCases, model.CovidCasesNoDate{CountryName: cases[i].CountryName, Info: cases[i].Info})
				if cases[i].CountryName == "China" {
					chinaCasesToday = int64(cases[i].Info)
				}

			} else {
				curDate = cases[i].Date
				break
			}
		}
		lenAcc := len(accumulativeCases)
		for ; i < length; i++ {
			if cases[i].Date == curDate {
				if cases[i].CountryName == "Global" {
					globalCasesYesterday = int64(cases[i].Info)
					continue
				}
				j := 0
				for j = 0; j < lenAcc; j++ {
					if accumulativeCases[j].CountryName == cases[i].CountryName {
						newCases = append(newCases, model.CovidCasesNoDate{CountryName: cases[i].CountryName, Info: accumulativeCases[j].Info - cases[i].Info})
						break
					}
				}
				if j == lenAcc { // 说明没找到这个国家在最新一天的累计确诊数，用0x3f3f3f3f标记
					newCases = append(newCases, model.CovidCasesNoDate{CountryName: cases[i].CountryName, Info: 0x3f3f3f3f})
				}
				if cases[i].CountryName == "China" {
					chinaCasesYesterday = int64(cases[i].Info)
				}
			} else {
				break
			}
		}
		nums = append(nums, globalCasesToday)
		nums = append(nums, globalCasesToday-globalCasesYesterday)
		nums = append(nums, chinaCasesToday)
		nums = append(nums, chinaCasesToday-chinaCasesYesterday)
		return accumulativeCases, newCases, nums, false
	}
}

// 查询中国各省份的累计确诊数和新增确诊数
func QueryChinaProvinceDetailCases() (accumulativeCases []model.CovidChinaCasesNoDate, newCases []model.CovidChinaCasesNoDate, notFound bool) {
	var cases []model.CovidChinaCases
	err := global.DB.Order("date desc, province_name asc").Find(&cases).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return accumulativeCases, newCases, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		i := 0
		for i = 0; i < 34; i++ {
			accumulativeCases = append(accumulativeCases, model.CovidChinaCasesNoDate{ProvinceName: cases[i].ProvinceName, Info: cases[i].Info})
		}
		for ; i < 68; i++ {
			newCases = append(newCases, model.CovidChinaCasesNoDate{ProvinceName: cases[i].ProvinceName, Info: cases[i-34].Info - cases[i].Info})
		}
		return accumulativeCases, newCases, false
	}
}

// 查询中国各省份的累计死亡数和新增死亡数
func QueryChinaProvinceDetailDeaths() (accumulativeDeaths []model.CovidChinaDeathsNoDate, newDeaths []model.CovidChinaDeathsNoDate, notFound bool) {
	var deaths []model.CovidChinaDeaths
	err := global.DB.Order("date desc, province_name asc").Find(&deaths).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return accumulativeDeaths, newDeaths, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		i := 0
		for i = 0; i < 34; i++ {
			accumulativeDeaths = append(accumulativeDeaths, model.CovidChinaDeathsNoDate{ProvinceName: deaths[i].ProvinceName, Info: deaths[i].Info})
		}
		for ; i < 68; i++ {
			newDeaths = append(newDeaths, model.CovidChinaDeathsNoDate{ProvinceName: deaths[i].ProvinceName, Info: deaths[i-34].Info - deaths[i].Info})
		}
		return accumulativeDeaths, newDeaths, false
	}
}

// 查询中国各省份的累计治愈数和新增治愈数
func QueryChinaProvinceDetailRecovered() (accumulativeRecovered []model.CovidChinaRecoveredNoDate, newRecovered []model.CovidChinaRecoveredNoDate, notFound bool) {
	var recovered []model.CovidChinaRecovered
	err := global.DB.Order("date desc, province_name asc").Find(&recovered).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return accumulativeRecovered, newRecovered, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		i := 0
		for i = 0; i < 34; i++ {
			accumulativeRecovered = append(accumulativeRecovered, model.CovidChinaRecoveredNoDate{ProvinceName: recovered[i].ProvinceName, Info: recovered[i].Info})
		}
		for ; i < 68; i++ {
			newRecovered = append(newRecovered, model.CovidChinaRecoveredNoDate{ProvinceName: recovered[i].ProvinceName, Info: recovered[i-34].Info - recovered[i].Info})
		}
		return accumulativeRecovered, newRecovered, false
	}
}
