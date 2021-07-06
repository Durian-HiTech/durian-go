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
	global.DB.Find(&notice)
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
