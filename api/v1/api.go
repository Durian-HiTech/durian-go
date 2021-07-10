package v1

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/TualatinX/durian-go/model"
	"github.com/TualatinX/durian-go/service"
	"github.com/gin-gonic/gin"
)

// Index doc
// @description 测试 Index 页
// @Tags 测试
// @Success 200 {string} string "{"success": true, "message": "gcp"}"
// @Router / [GET]
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "gcp"})
}

// Register doc
// @description 注册
// @Tags 用户管理
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param info formData string true "用户个人信息"
// @Param user_type formData string true "用户类型（0: 普通用户，1: 认证机构用户）"
// @Param affiliation formData string false "认证机构名"
// @Success 200 {string} string "{"success": true, "message": "用户创建成功"}"
// @Failure 200 {string} string "{"success": false, "message": "用户已存在"}"
// @Router /user/register [POST]
func Register(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	info := c.Request.FormValue("info")
	userType, _ := strconv.ParseUint(c.Request.FormValue("user_type"), 0, 64)
	affiliation := c.Request.FormValue("affiliation")
	user := model.User{Username: username, Password: password, UserType: userType, Affiliation: affiliation}
	_, notFound := service.QueryAUserByUsername(username)
	if notFound {
		service.CreateAUser(&user)
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "用户创建成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "用户已存在"})
	}
}

// Login doc
// @description 登录
// @Tags 用户管理
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param info formData string true "用户个人信息"
// @Success 200 {string} string "{"success": true, "message": "登录成功", "detail": user的信息}"
// @Failure 200 {string} string "{"success": false, "message": "密码错误"}"
// @Failure 200 {string} string "{"success": false, "message": "没有该用户"}"
// @Router /user/login [POST]
func Login(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	user, notFound := service.QueryAUserByUsername(username)
	if notFound {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "没有该用户"})
	} else {
		if user.Password != password {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": "密码错误"})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "登录成功", "detail": user})
		}
	}
}

// ModifyUser doc
// @description 修改用户信息（支持修改用户名和密码）
// @Tags 用户管理
// @Param user_id formData string true "用户ID"
// @Param username formData string true "用户名"
// @Param password_old formData string true "原密码"
// @Param password_new formData string true "新密码"
// @Success 200 {string} string "{"success": true, "message": "修改成功", "data": "model.User的所有信息"}"
// @Failure 200 {string} string "{"success": false, "message": "原密码输入错误"}"
// @Failure 200 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Failure 400 {string} string "{"success": false, "message": "数据库操作时的其他错误"}"
// @Router /user/modify [POST]
func ModifyUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	username := c.Request.FormValue("username")
	info := c.Request.FormValue("info")
	passwordOld := c.Request.FormValue("password_old")
	passwordNew := c.Request.FormValue("password_new")
	user, notFoundUserByID := service.QueryAUserByID(userID)
	if notFoundUserByID {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户ID不存在",
		})
		return
	}
	if passwordOld != user.Password {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "原密码输入错误",
		})
		return
	}
	_, notFoundUserByName := service.QueryAUserByUsername(username)
	if !notFoundUserByName && username != user.Username {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户名已被占用",
		})
		return
	}
	err := service.UpdateAUser(&user, username, passwordNew, info)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	//data, _ := jsoniter.Marshal(&user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
		"data":    user,
	})
}

// TellUserInfo doc
// @description 查看用户个人信息
// @Tags 用户管理
// @Param user_id formData string true "用户ID"
// @Success 200 {string} string "{"success": true, "message": "查看用户信息成功", "data": "model.User的所有信息"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /user/info [POST]
func TellUserInfo(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user, notFoundUserByID := service.QueryAUserByID(userID)
	if notFoundUserByID {
		c.JSON(404, gin.H{
			"success": false,
			"message": "用户ID不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "查看用户信息成功",
		"data":    user,
	})
}

// Subscribe doc
// @description 订阅城市疫情信息
// @Tags 订阅城市
// @Param user_id formData string true "用户ID"
// @Param city_name formData string true "城市名字"
// @Success 200 {string} string "{"success":true, "message":"订阅成功"}"
// @Failure 200 {string} string "{"success": false, "message": "已经订阅过这个城市的疫情信息"}"
// @Failure 401 {string} string "{"success": false, "message": "数据库error, 一些其他错误"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /sub/subscribe [POST]
func Subscribe(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	cityName := c.Request.FormValue("city_name")

	_, notFoundUserByID := service.QueryAUserByID(userID)
	if notFoundUserByID {
		c.JSON(404, gin.H{
			"success": false,
			"message": "用户ID不存在",
		})
		return
	}

	_, notFoundSubscriptionByUserIDAndCityName := service.QueryASubscriptionByUserIDAndCityName(userID, cityName)

	if !notFoundSubscriptionByUserIDAndCityName {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "已经订阅过这个城市的疫情信息",
		})
		return
	}

	if err := service.CreateASubscription(userID, cityName); err != nil {
		c.JSON(401, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "订阅成功"})
	}
}

// ListAllSubscriptions doc
// @description 获取订阅列表
// @Tags 订阅城市
// @Param user_id formData string true "用户ID"
// @Param city_name formData string true "城市名字"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"user的所有订阅"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /sub/list_all_subs [POST]
func ListAllSubscriptions(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	_, notFoundUserByID := service.QueryAUserByID(userID)
	if notFoundUserByID {
		c.JSON(404, gin.H{
			"success": false,
			"message": "用户ID不存在",
		})
		return
	}
	subscriptions := service.QueryAllSubscriptions(userID)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": subscriptions})
}

// RemoveSubscription doc
// @description 删除订阅
// @Tags 订阅城市
// @Param user_id formData string true "用户ID"
// @Param subscription_id formData string true "订阅ID"
// @Success 200 {string} string "{"success":true, "message":"删除成功"}"
// @Failure 401 {string} string "{"success": false, "message": "数据库error, 一些其他错误"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /sub/del_sub [POST]
func RemoveSubscription(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	cityName := c.Request.FormValue("city_name")
	if err := service.DeleteASubscription(userID, cityName); err != nil {
		c.JSON(401, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "删除成功"})
	}
}

// CreateAQuestion doc
// @description 创建一个问题
// @Tags 防控知识板块
// @Param user_id formData string true "用户ID"
// @Param question_title formData string true "提问标题"
// @Param question_content formData string true "提问内容"
// @Success 200 {string} string "{"success": true, "message": "用户提问成功", "detail": "提问的全部信息"}"
// @Failure 401 {string} string "{"success": false, "message": "数据库error, 一些其他错误"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /notice/create_question [POST]
func CreateAQuestion(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)

	_, notFoundUserByID := service.QueryAUserByID(userID)
	if notFoundUserByID {
		c.JSON(404, gin.H{
			"success": false,
			"message": "用户ID不存在",
		})
		return
	}

	questionTitle := c.Request.FormValue("question_title")
	questionContent := c.Request.FormValue("question_content")

	question := model.Question{UserID: userID, QuestionTitle: questionTitle, QuestionContent: questionContent, QuestionTime: time.Now()}
	err := service.CreateAQuestion(&question)
	if err != nil {
		c.JSON(401, gin.H{
			"success": false,
			"message": "提问失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "用户提问成功", "detail": question})
	}
}

// ListAllQuestions doc
// @description 列出全部问题
// @Tags 防控知识板块
// @Success 200 {string} string "{"success": true, "message": "查看成功", "data": "全部问题"}"
// @Router /notice/list_all_questions [GET]
func ListAllQuestions(c *gin.Context) {
	questionList := service.QueryAllQuestions()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": questionList})
}

// ListAQuestion doc
// @description 列出某个问题的详情
// @Tags 防控知识板块
// @Param question_id formData string true "问题ID"
// @Success 200 {string} string "{"success": true, "message": "查看成功", "data": "某问题的所有信息"}"
// @Failure 404 {string} string "{"success": false, "message": "问题ID不存在"}"
// @Router /notice/question_detail [POST]
func ListAQuestion(c *gin.Context) {
	questionID, _ := strconv.ParseUint(c.Request.FormValue("question_id"), 0, 64)
	question, notFoundQuestionByID := service.QueryAQuestionByID(questionID)
	if notFoundQuestionByID {
		c.JSON(404, gin.H{
			"success": false,
			"message": "问题ID不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查看成功", "data": question})
}

// ListAllComments doc
// @description 列出某个问题的全部评论
// @Tags 防控知识板块
// @Param question_id formData string true "问题ID"
// @Success 200 {string} string "{"success": true, "message": "查看成功", "data": "某问题的所有评论"}"
// @Failure 404 {string} string "{"success": false, "message": "问题ID不存在"}"
// @Router /notice/list_all_comments [POST]
func ListAllComments(c *gin.Context) {
	questionID, _ := strconv.ParseUint(c.Request.FormValue("question_id"), 0, 64)
	_, notFoundQuestionByID := service.QueryAQuestionByID(questionID)
	if notFoundQuestionByID {
		c.JSON(404, gin.H{
			"success": false,
			"message": "问题ID不存在",
		})
		return
	}
	comments := service.QueryAllComments(questionID)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查看成功", "data": comments})
}

// CreateAComment doc
// @description 创建一条评论
// @Tags 防控知识板块
// @Param user_id formData string true "用户ID"
// @Param user_type formData string true "用户类型"
// @Param question_id formData string true "问题ID"
// @Param comment_content formData string true "评论内容"
// @Success 200 {string} string "{"success": true, "message": "用户评论成功"}"
// @Failure 400 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Failure 401 {string} string "{"success": false, "message": "问题ID不存在"}"
// @Failure 402 {string} string "{"success": false, "message": "评论失败"}"
// @Router /notice/create_comment [POST]
func CreateAComment(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	_, notFoundUserByID := service.QueryAUserByID(userID)
	if notFoundUserByID {
		c.JSON(400, gin.H{
			"success": false,
			"message": "用户ID不存在",
		})
		return
	}

	questionID, _ := strconv.ParseUint(c.Request.FormValue("question_id"), 0, 64)
	_, notFoundQuestionByID := service.QueryAQuestionByID(questionID)
	if notFoundQuestionByID {
		c.JSON(401, gin.H{
			"success": false,
			"message": "问题ID不存在",
		})
		return
	}

	userType, _ := strconv.ParseUint(c.Request.FormValue("user_type"), 0, 64)
	commentContent := c.Request.FormValue("comment_content")
	comment := model.Comment{UserID: userID, QuestionID: questionID, CommentContent: commentContent, CommentTime: time.Now(), UserType: userType}
	err := service.CreateAComment(&comment)
	if err != nil {
		c.JSON(402, gin.H{
			"success": false,
			"message": "评论失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "用户评论成功"})
}

// ListAllNotice doc
// @description 获取所有公告，返回列表
// @Tags 公告
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有公告""}"
// @Router /notice/list_all_notice [GET]
func ListAllNotice(c *gin.Context) {
	noticeList := service.QueryAllNotice()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": noticeList})
}

// ViewNewsDetail doc
// @description 查看单条公告
// @Tags 公告
// @Param notice_id formData string true "公告ID"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"该条公告的详细信息"}"
// @Failure 404 {string} string "{"success":true, "message":"查询失败，公告ID不存在"}"
// @Router /notice/notice_detail [POST]
func ViewNoticeDetail(c *gin.Context) {
	noticeID, _ := strconv.ParseUint(c.Request.FormValue("notice_id"), 0, 64)
	notice, notFound := service.QueryANoticeByID(noticeID)
	if !notFound {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": notice})
	} else {
		c.JSON(404, gin.H{"success": false, "message": "查询失败，公告ID不存在"})
	}
}

// ListAllRumor doc
// @description 获取所有辟谣，返回列表
// @Tags 公告
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有辟谣""}"
// @Router /notice/list_all_rumor [GET]
func ListAllRumor(c *gin.Context) {
	rumorList := service.QueryAllRumor()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": rumorList})
}

// ViewRumorDetail doc
// @description 查看单条辟谣
// @Tags 公告
// @Param rumor_id formData string true "辟谣ID"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"该条辟谣的详细信息"}"
// @Failure 404 {string} string "{"success":true, "message":"查询失败，辟谣ID不存在"}"
// @Router /notice/rumor_detail [POST]
func ViewRumorDetail(c *gin.Context) {
	rumorID, _ := strconv.ParseUint(c.Request.FormValue("rumor_id"), 0, 64)
	rumor, notFound := service.QueryARumorByID(rumorID)
	if !notFound {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": rumor})
	} else {
		c.JSON(404, gin.H{"success": false, "message": "查询失败，辟谣ID不存在"})
	}
}

// ListAllNews doc
// @description 获取所有新闻，返回列表
// @Tags 新闻
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有新闻""}"
// @Router /news/list_all_news [GET]
func ListAllNews(c *gin.Context) {
	newsList := service.QueryAllNews()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": newsList})
}

// ViewNewsDetail doc
// @description 查看单条新闻
// @Tags 新闻
// @Param news_id formData string true "新闻ID"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"该条新闻的详细信息"}"
// @Failure 404 {string} string "{"success":true, "message":"查询失败，新闻ID不存在"}"
// @Router /news/detail [POST]
func ViewNewsDetail(c *gin.Context) {
	newsID, _ := strconv.ParseUint(c.Request.FormValue("news_id"), 0, 64)
	news, notFound := service.QueryANewsByID(newsID)
	if !notFound {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": news})
	} else {
		c.JSON(404, gin.H{"success": false, "message": "查询失败，新闻ID不存在"})
	}
}

// ListHighRiskAreas doc
// @description 获取所有中高风险地区，返回列表
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有中高风险地区"}"
// @Router /data/list_all_high_risk_areas [GET]
func ListHighRiskAreas(c *gin.Context) {
	areas := service.QueryAllHighRiskAreas()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": areas})
}

// FetchRequiredData doc
// @description 获取在数据库中直接存的 Json File
// @Tags 数据
// @Param name formData string true "数据文件名"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"Json data"}"
// @Failure 200 {string} string "{"success":true, "message":"查询失败，无所需数据"}"
// @Router /data/query_data [POST]
func FetchRequiredData(c *gin.Context) {
	name := c.Request.FormValue("name")
	// directData, notFound := service.QueryDataByName(name)
	// if notFound {
	// 	c.JSON(http.StatusOK, gin.H{"success": false, "message": "查询失败，无所需数据"})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": directData})
	// }
	fin, err := os.Open("./data/" + name + ".json")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "查询失败，无所需数据"})
	} else {
		cin, _ := ioutil.ReadAll(fin)
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": string(cin)})
	}
}

// ListAllCovidCases doc
// @description 获取所有地区的新冠感染人数，返回列表
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠感染人数"}"
// @Router /data/list_all_covid_cases [GET]
func ListAllCovidCases(c *gin.Context) {
	covidList := service.QueryAllCovidCases()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": covidList})
}

// ListAllCovidCasesResponse doc
// @description 获取所有地区的新冠感染人数，返回列表 [根据时间分组]
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠感染人数"}"
// @Router /data/list_all_covid_cases_response [GET]
func ListAllCovidCasesResponse(c *gin.Context) {
	covidList := service.QueryAllCovidCasesResponse()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": covidList})
}

// ListAllCovidDeaths doc
// @description 获取所有地区的新冠死亡人数，返回列表
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠死亡人数"}"
// @Router /data/list_all_covid_deaths [GET]
func ListAllCovidDeaths(c *gin.Context) {
	covidList := service.QueryAllCovidDeaths()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": covidList})
}

// ListAllCovidCasesResponse doc
// @description 获取所有地区的新冠死亡人数，返回列表 [根据时间分组]
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠死亡人数"}"
// @Router /data/list_all_covid_deaths_response [GET]
func ListAllCovidDeathsResponse(c *gin.Context) {
	covidList := service.QueryAllCovidDeathsResponse()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": covidList})
}

// ListAllCovidRecovereds doc
// @description 获取所有地区的新冠治愈人数，返回列表
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠治愈人数""}"
// @Router /data/list_all_covid_recovereds [GET]
func ListAllCovidRecovereds(c *gin.Context) {
	covidList := service.QueryAllCovidRecovereds()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": covidList})
}

// ListAllCovidCasesResponse doc
// @description 获取所有地区的新冠治愈人数，返回列表 [根据时间分组]
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠治愈人数"}"
// @Router /data/list_all_covid_recovereds_response [GET]
func ListAllCovidRecoveredsResponse(c *gin.Context) {
	covidList := service.QueryAllCovidRecoveredsResponse()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": covidList})
}

// ListAllCovidVaccines doc
// @description 获取所有地区的新冠疫苗接种人数，返回列表
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠疫苗接种人数""}"
// @Router /data/list_all_covid_vaccines [GET]
func ListAllCovidVaccines(c *gin.Context) {
	covidList := service.QueryAllCovidVaccines()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": covidList})
}

// ListAllCovidVaccinesResponse doc
// @description 获取所有地区的新冠疫苗接种人数，返回列表 [根据时间分组]
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠治愈人数"}"
// @Router /data/list_all_covid_vaccines_response [GET]
func ListAllCovidVaccinesResponse(c *gin.Context) {
	covidList := service.QueryAllCovidVaccinesResponse()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": covidList})
}

// ListAllCovidCDRV doc
// @description 获取所有地区的新冠感染/死亡/治愈/疫苗接种人数【信息综合】，返回列表
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠感染/死亡/治愈/疫苗接种人数""}"
// @Router /data/list_all_covid_cdrv [GET]
func ListAllCovidCDRV(c *gin.Context) {
	covidListC := service.QueryAllCovidCases()
	covidListD := service.QueryAllCovidDeaths()
	covidListR := service.QueryAllCovidRecovereds()
	covidListV := service.QueryAllCovidVaccines()
	result := model.CovidCDRV{Case: covidListC, Deaths: covidListD, Recovered: covidListR, Vaccine: covidListV}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": result})
}

// ListAllCovidCDRV doc
// @description 获取所有地区的新冠感染/死亡/治愈/疫苗接种人数【信息综合】，返回列表 [根据时间分组]
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠感染/死亡/治愈/疫苗接种人数""}"
// @Router /data/list_all_covid_cdrv_response [GET]
func ListAllCovidCDRVResponse(c *gin.Context) {
	covidListC := service.QueryAllCovidCasesResponse()
	covidListD := service.QueryAllCovidDeathsResponse()
	covidListR := service.QueryAllCovidRecoveredsResponse()
	covidListV := service.QueryAllCovidVaccinesResponse()
	result := model.CovidCDRVResponse{Case: covidListC, Deaths: covidListD, Recovered: covidListR, Vaccine: covidListV}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": result})
}

// ListAllCovidCasesResponseProvince doc
// @description 获取所有地区的新冠感染人数，返回列表 [根据时间分组] [Province]
// @Tags 数据
// @Param province formData string true "区域名"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠感染人数 [Province]"}"
// @Router /data/list_all_covid_cases_response_province [POST]
func ListAllCovidCasesResponseProvince(c *gin.Context) {
	provinceName := c.Request.FormValue("province")
	covidList, _ := service.QueryAllCovidCasesResponseProvince(provinceName)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": covidList})
}

// ListAllCovidDeathsResponseProvince doc
// @description 获取所有地区的新冠死亡人数，返回列表 [根据时间分组] [Province]
// @Tags 数据
// @Param province formData string true "区域名"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠死亡人数 [Province]"}"
// @Router /data/list_all_covid_deaths_response_province [POST]
func ListAllCovidDeathsResponseProvince(c *gin.Context) {
	provinceName := c.Request.FormValue("province")
	covidList, _ := service.QueryAllCovidDeathsResponseProvince(provinceName)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": covidList})
}

// ListAllCovidCasesResponseProvince doc
// @description 获取所有地区的新冠治愈人数，返回列表 [根据时间分组] [Province]
// @Tags 数据
// @Param province formData string true "区域名"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠治愈人数 [Province]"}"
// @Router /data/list_all_covid_recovereds_response_province [POST]
func ListAllCovidRecoveredsResponseProvince(c *gin.Context) {
	provinceName := c.Request.FormValue("province")
	covidList, _ := service.QueryAllCovidRecoveredsResponseProvince(provinceName)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": covidList})
}

// ListAllCovidCDRVProvince doc
// @description 获取所有地区的新冠感染/死亡/治愈【信息综合】，返回列表 [根据时间分组] [Province]
// @Tags 数据
// @Param province formData string true "区域名"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有地区的新冠感染/死亡/治愈/疫苗接种人数 [Province]"}"
// @Router /data/list_all_covid_cdrv_response_province [POST]
func ListAllCovidCDRVResponseProvince(c *gin.Context) {
	provinceName := c.Request.FormValue("province")
	covidListC, _ := service.QueryAllCovidCasesResponseProvince(provinceName)
	covidListD, _ := service.QueryAllCovidDeathsResponseProvince(provinceName)
	covidListR, _ := service.QueryAllCovidRecoveredsResponseProvince(provinceName)
	covidListV := covidListR
	for _, v := range covidListV {
		v.Info = "{}"
	}
	result := model.CovidCDRVResponseProvince{Case: covidListC, Deaths: covidListD, Recovered: covidListR, Vaccine: covidListV}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": result})
}

// ListOverviewData doc
// @description 获取世界或中国的现存确诊、新增确诊、累积确诊、累计及新增新冠感染/死亡/治愈【信息综合】，返回列表 [根据时间分组] [Province]
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","nowcases":{"nownum": 123, "newnum": 123}等数据}"
// @Router /data/list_overview [GET]
func ListOverviewData(c *gin.Context) {
	accumulativeDeaths, newDeaths, numDeaths, _ := service.QueryDeathOverview()
	accumulativeRecovered, newRecovered, numRecovered, _ := service.QueryRecoveredOverview()
	accumulativeVaccine, newVaccine, numVaccine, _ := service.QueryVaccineOverview()
	accumulativeCases, newCases, numCases, _ := service.QueryCasesOverview()

	// Global的overview部分
	globalOverviewCases = numCases[0]                                     // 全球总确诊
	globalOverviewNowCases = numCases[0] - numDeaths[0] - numRecovered[0] // 全球现存确诊=全球总确诊-全球总死亡-全球总治愈
	globalOverviewNewCases = numCases[1]                                  // 全球新增确诊

	globalOverviewDeathsNownum = numDeaths[0]       // 全球累计死亡数
	globalOverviewDeathsNewnum = numDeaths[1]       // 全球新增死亡数
	globalOverviewRecoveredNownum = numRecovered[0] // 全球累计治愈数
	globalOverviewRecoveredNewnum = numRecovered[1] // 全球新增治愈数
	globalOverviewVaccineNownum = numVaccine[0]     // 全球累计疫苗接种数
	globalOverviewVaccineNewnum = numVaccine[1]     // 全球新增疫苗接种数

	// Global的detail部分
	i := 0
	var detail []map[string]string

	for ; i < 197; i++ {
		m := make(map[string]string)
		m["name"] = accumulativeCases[i].CountryName

		m["cases"] = accumulativeCases[i].Info
		m["nowcases"] = accumulativeCases[i].Info - accumulativeDeaths[i].Info - accumulativeRecovered[i] // 现存确诊=累计确诊-累计死亡-累计治愈
		m["newcases"] = newCases[i].Info

		m["vaccine"] = accumulativeVaccine[i].Info
		m["newvaccine"] = newVaccine[i].Info
		m["recovered"] = accumulativeRecovered[i].Info
		m["newrecovered"] = newRecovered[i].Info
		m["deaths"] = accumulativeDeaths[i].Info
		m["newdeaths"] = newDeaths[i].Info

		detail = append(detail, m)
	}
	// 中国部分

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", 
		"data": {
			"Global": {
				"overview": {
					"nowcases": {
						"nownum": globalOverviewNowCases,
						"newnum": globalOverviewNewCases,
					},
					"cases": {
						"nownum": globalOverviewCases,
					},
					"deaths": {
						"nownum": globalOverviewDeathsNownum,
						"newnum": globalOverviewDeathsNewnum,
					},
					"recovered": {
						"nownum": globalOverviewRecoveredNownum,
						"newnum": globalOverviewRecoveredNewnum,
					},
					"vaccine": {
						"nownum": globalOverviewVaccineNownum,
						"newnum": globalOverviewVaccineNewnum,
					},
				},
				"detailed": detail,
			},
			"China": {
				"overview": {
					"nowcases": {
						"nownum": numCases[2],
						"newnum": numCases[2] - numDeaths[2] - numRecovered[2],
					},
					"cases": {
						"nownum": numCases[3],
					},
					"deaths": {
						"nownum": numDeaths[2],
						"newnum": numDeaths[3],
					},
					"recovered": {
						"nownum": numRecovered[2],
						"newnum": numRecovered[3],
					},
					"vaccine": {
						"nownum": numVaccine[2],
						"newnum": numVaccine[3],
					},
				},
			}
		}
	})
}
