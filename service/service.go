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

// 更新用户的用户名、密码信息
func UpdateAUser(user *model.User, username string, password string) error {
	user.Username = username
	user.Password = password
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

// 根据用户名和其订阅城市名查询某个订阅情况
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

// 删除订阅城市
func DeleteASubscription(subscriptionID uint64) (err error) {
	var subscription model.Subscription
	err = global.DB.First(&subscription, subscriptionID).Error
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
func QueryAQuestionByID(questionID uint64) (question model.Question, notFound bool) {
	err := global.DB.Where("question_id = ?", questionID).First(&question).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return question, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		return question, false
	}
}

// 查询所有问题
func QueryAllQuestions() (questions []model.Question) {
	global.DB.Order("question_time desc").Find(&questions)
	return questions
}

// 创建一个对问题的评论
func CreateAComment(comment *model.Comment) (err error) {
	if err = global.DB.Create(&comment).Error; err != nil {
		return err
	}
	return
}

// 列出某个问题的所有评论
func QueryAllComments(questionID uint64) (res []model.Comment) {
	var comments []model.Comment
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
	return res
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
func QueryAllCovidCasesResponseProvince(province string) (response []model.CovidCasesProvince, notFound bool) {
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
func QueryAllCovidDeathsResponseProvince(province string) (response []model.CovidDeathsProvince, notFound bool) {
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
func QueryAllCovidRecoveredsResponseProvince(province string) (response []model.CovidRecoveredProvince, notFound bool) {
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
func QueryDeathOverview() (accumulative_deaths []model.CovidDeathsNoDate, new_deaths []model.CovidDeathsNoDate, nums []int64, notFound bool) {
	var deaths []model.CovidDeaths
	err := global.DB.Order("date desc, country_name asc").Find(&deaths).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return accumulative_deaths, new_deaths, nums, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		var global_death_today int64
		var global_death_yesterday int64
		var china_death_today int64
		var china_death_yesterday int64

		length := len(deaths)
		cur_date := deaths[0].Date
		i := 0
		flag := 0

		global_death_today = 0
		global_death_yesterday = 0
		china_death_today = 0
		china_death_yesterday = 0

		for i = 0; i < length; i++ {
			if deaths[i].Date == cur_date {
				if flag == 0 && deaths[i].CountryName == "Global" {
					global_death_today = int64(deaths[i].Info)
					flag += 1
					continue
				}
				accumulative_deaths = append(accumulative_deaths, model.CovidDeathsNoDate{CountryName: deaths[i].CountryName, Info: deaths[i].Info})
				if deaths[i].CountryName == "China" {
					china_death_today = int64(deaths[i].Info)
				}

			} else {
				cur_date = deaths[i].Date
				break
			}
		}
		len_acc := len(accumulative_deaths)
		for ; i < length; i++ {
			if deaths[i].Date == cur_date {
				if flag == 1 && deaths[i].CountryName == "Global" {
					global_death_yesterday = int64(deaths[i].Info)
					flag += 1
					continue
				}
				j := 0
				for j = 0; j < len_acc; j++ {
					if accumulative_deaths[j].CountryName == deaths[i].CountryName {
						new_deaths = append(new_deaths, model.CovidDeathsNoDate{CountryName: deaths[i].CountryName, Info: accumulative_deaths[j].Info - deaths[i].Info})
						break
					}
				}
				if j == len_acc { // 说明没找到这个国家在最新一天的累计死亡数，用0x3f3f3f3f标记
					new_deaths = append(new_deaths, model.CovidDeathsNoDate{CountryName: deaths[i].CountryName, Info: 0x3f3f3f3f})
				}
				if deaths[i].CountryName == "China" {
					china_death_yesterday = int64(deaths[i].Info)
				}
			} else {
				break
			}
		}
		nums = append(nums, global_death_today)
		nums = append(nums, global_death_today-global_death_yesterday)
		nums = append(nums, china_death_today)
		nums = append(nums, china_death_today-china_death_yesterday)
		return accumulative_deaths, new_deaths, nums, false
	}
}

// 查询中国或世界的累计治愈数和新增治愈数
func QueryRecoveredOverview() (accumulative_recovered []model.CovidRecoveredNoDate, new_recovered []model.CovidRecoveredNoDate, nums []int64, notFound bool) {
	var recovered []model.CovidRecovered
	err := global.DB.Order("date desc, country_name asc").Find(&recovered).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return accumulative_recovered, new_recovered, nums, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		var global_recovered_today int64
		var global_recovered_yesterday int64
		var china_recovered_today int64
		var china_recovered_yesterday int64

		length := len(recovered)
		cur_date := recovered[0].Date
		i := 0
		flag := 0

		global_recovered_today = 0
		global_recovered_yesterday = 0
		china_recovered_today = 0
		china_recovered_yesterday = 0

		for i = 0; i < length; i++ {
			if recovered[i].Date == cur_date {
				if flag == 0 && recovered[i].CountryName == "Global" {
					global_recovered_today = int64(recovered[i].Info)
					flag += 1
					continue
				}
				accumulative_recovered = append(accumulative_recovered, model.CovidRecoveredNoDate{CountryName: recovered[i].CountryName, Info: recovered[i].Info})
				if recovered[i].CountryName == "China" {
					china_recovered_today = int64(recovered[i].Info)
				}

			} else {
				cur_date = recovered[i].Date
				break
			}
		}
		for ; i < length; i++ {
			if recovered[i].Date == cur_date {
				if flag == 1 && recovered[i].CountryName == "Global" {
					global_recovered_yesterday = int64(recovered[i].Info)
					flag += 1
					continue
				}
				if i > 196 && recovered[i-197].CountryName == recovered[i].CountryName {
					new_recovered = append(new_recovered, model.CovidRecoveredNoDate{CountryName: recovered[i].CountryName, Info: recovered[i-197].Info - recovered[i].Info})
				} else { // 说明没找到这个国家在最新一天的累计治愈数，用0x3f3f3f3f标记
					new_recovered = append(new_recovered, model.CovidRecoveredNoDate{CountryName: recovered[i].CountryName, Info: 0x3f3f3f3f})
				}
				if recovered[i].CountryName == "China" {
					china_recovered_yesterday = int64(recovered[i].Info)
				}
			} else {
				break
			}
		}
		nums = append(nums, global_recovered_today)
		nums = append(nums, global_recovered_today-global_recovered_yesterday)
		nums = append(nums, china_recovered_today)
		nums = append(nums, china_recovered_today-china_recovered_yesterday)
		return accumulative_recovered, new_recovered, nums, false
	}
}

// 查询中国或世界的累计接种数和新增接种数
func QueryVaccineOverview() (accumulative_vaccine []model.CovidVaccineNoDate, new_vaccine []model.CovidVaccineNoDate, nums []int64, notFound bool) {
	var vaccine []model.CovidVaccine
	err := global.DB.Order("date desc, country_name asc").Find(&vaccine).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return accumulative_vaccine, new_vaccine, nums, true
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	} else {
		var global_vaccine_today int64
		var global_vaccine_yesterday int64
		var china_vaccine_today int64
		var china_vaccine_yesterday int64

		length := len(vaccine)
		cur_date := vaccine[0].Date
		i := 0

		global_vaccine_today = 0
		global_vaccine_yesterday = 0
		china_vaccine_today = 0
		china_vaccine_yesterday = 0

		for i = 0; i < length; i++ {
			if vaccine[i].Date == cur_date {
				if vaccine[i].CountryName == "Global" {
					global_vaccine_today = int64(vaccine[i].Info)
					continue
				}
				accumulative_vaccine = append(accumulative_vaccine, model.CovidVaccineNoDate{CountryName: vaccine[i].CountryName, Info: vaccine[i].Info})
				if vaccine[i].CountryName == "China" {
					china_vaccine_today = int64(vaccine[i].Info)
				}

			} else {
				cur_date = vaccine[i].Date
				break
			}
		}
		for ; i < length; i++ {
			if vaccine[i].Date == cur_date {
				if vaccine[i].CountryName == "Global" {
					global_vaccine_yesterday = int64(vaccine[i].Info)
					continue
				}
				if i > 196 && vaccine[i-197].CountryName == vaccine[i].CountryName {
					new_vaccine = append(new_vaccine, model.CovidVaccineNoDate{CountryName: vaccine[i].CountryName, Info: vaccine[i-197].Info - vaccine[i].Info})
				} else { // 说明没找到这个国家在最新一天的累计治愈数，用0x3f3f3f3f标记
					new_vaccine = append(new_vaccine, model.CovidVaccineNoDate{CountryName: vaccine[i].CountryName, Info: 0x3f3f3f3f})
				}
				if vaccine[i].CountryName == "China" {
					china_vaccine_yesterday = int64(vaccine[i].Info)
				}
			} else {
				break
			}
		}
		nums = append(nums, global_vaccine_today)
		nums = append(nums, global_vaccine_today-global_vaccine_yesterday)
		nums = append(nums, china_vaccine_today)
		nums = append(nums, china_vaccine_today-china_vaccine_yesterday)
		return accumulative_vaccine, new_vaccine, nums, false
	}
}
