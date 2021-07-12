package v1

import (
	"fmt"
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
// @Param user_info formData string true "用户个人信息"
// @Param user_type formData string true "用户类型（0: 普通用户，1: 认证机构用户）"
// @Param affiliation formData string false "认证机构名"
// @Success 200 {string} string "{"success": true, "message": "用户创建成功"}"
// @Failure 200 {string} string "{"success": false, "message": "用户已存在"}"
// @Router /user/register [POST]
func Register(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	userInfo := c.Request.FormValue("user_info")
	userType, _ := strconv.ParseUint(c.Request.FormValue("user_type"), 0, 64)
	affiliation := c.Request.FormValue("affiliation")
	user := model.User{Username: username, Password: password, UserInfo: userInfo, UserType: userType, Affiliation: affiliation}
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
			subs := service.QueryAllSubscriptions(user.UserID)
			showSub := false
			for _, sub := range subs {
				if sub.Name == "云南省" {
					showSub = true
					break
				}
			}
			if !showSub {
				c.JSON(http.StatusOK, gin.H{"success": true, "message": "登录成功", "detail": user, "show_sub": false})
			} else {
				c.JSON(http.StatusOK, gin.H{"success": true, "message": "登录成功", "detail": user, "show_sub": true})
			}
		}
	}
}

// ModifyUser doc
// @description 修改用户信息（支持修改用户名和密码）
// @Tags 用户管理
// @Param user_id formData string true "用户ID"
// @Param username formData string true "用户名"
// @Param user_info formData string true "用户个人信息"
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
	userInfo := c.Request.FormValue("user_info")
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
	err := service.UpdateAUser(&user, username, passwordNew, userInfo)
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
// @Param name formData string true "城市名字"
// @Success 200 {string} string "{"success":true, "message":"订阅成功"}"
// @Failure 200 {string} string "{"success": false, "message": "已经订阅过这个城市的疫情信息"}"
// @Failure 401 {string} string "{"success": false, "message": "数据库error, 一些其他错误"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /sub/subscribe [POST]
func Subscribe(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	cityName := c.Request.FormValue("name")

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
// @Param name formData string true "城市名字"
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

// ListSubscriptionsData doc
// @description 返回订阅信息
// @Tags 订阅城市
// @Param user_id formData string true "用户ID"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"user的所有订阅"}"
// @Failure 401 {string} string "{"success": false, "message": "错误！订阅信息列表为空"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /sub/list_subs_data [POST]
func ListSubscriptionsData(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	_, notFoundUserByID := service.QueryAUserByID(userID)
	if notFoundUserByID {
		c.JSON(404, gin.H{
			"success": false,
			"message": "用户ID不存在",
		})
		return
	}
	subscriptions := service.QueryAllSubscriptions(userID) // 用户订阅的省份、直辖市的列表
	subscriptionsData, length := service.QuerySubcriptionsData(subscriptions)
	if length == 0 {
		c.JSON(401, gin.H{
			"success": false,
			"message": "错误！订阅信息列表为空",
		})
		return
	}
	var information string
	for i := 0; i < length; i++ {
		if subscriptionsData[i].NewCases != 0 {
			information += fmt.Sprintf("%s有 %d 个新增确诊病例 ", subscriptionsData[i].ProvinceName, subscriptionsData[i].NewCases)
		}
	}
	informationLength := len(information)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": subscriptionsData, "information": information[0 : informationLength-1]})
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
	cityName := c.Request.FormValue("name")
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
// @Success 200 {string} string "{"success": true, "message": "查看成功", "data": "全部问题列表", "recommend_questions":"推荐问题列表"}"
// @Router /notice/list_all_questions [GET]
func ListAllQuestions(c *gin.Context) {
	questionList, recommendQuestions := service.QueryAllQuestions()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": questionList, "recommend_questions": recommendQuestions})
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

// ListAllNotice doc
// @description 获取所有防疫知识，返回列表
// @Tags 防疫知识
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有防疫知识""}"
// @Router /notice/list_all_knowledge [GET]
func ListAllKnowledge(c *gin.Context) {
	knowledgeList := service.QueryAllKnowledge()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": knowledgeList})
}

// ViewNewsDetail doc
// @description 查看单条防疫知识
// @Tags 防疫知识
// @Param notice_id formData string true "公告ID"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"该条防疫知识的详细信息"}"
// @Failure 404 {string} string "{"success":true, "message":"查询失败，防疫知识ID不存在"}"
// @Router /notice/knowledge_detail [POST]
func ViewKnowledgeDetail(c *gin.Context) {
	knowledgeID, _ := strconv.ParseUint(c.Request.FormValue("knowledge_id"), 0, 64)
	knowledge, notFound := service.QueryAKnowledgeByID(knowledgeID)
	if !notFound {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": knowledge})
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

// ListAllFlights doc
// @description 获取所有航班信息，返回列表
// @Tags 出行
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有航班信息""}"
// @Router /travel/list_all_flights [GET]
func ListAllFlights(c *gin.Context) {
	flightList := service.QueryAllFlights()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": flightList})
}

// FindSpecificFlightInfo doc
// @description 根据起始地查询航班信息，返回列表
// @Tags 出行
// @Param departure_city formData string true "出发地"
// @Param arrival_city formData string true "到达地"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"查询到的航班信息""}"
// @Failue 200 {string} string "{"success":true, "message":"查询成功","data":"没有找到该航班""}"
// @Router /travel/find_specific_flight_info [POST]
func FindSpecificFlightInfo(c *gin.Context) {
	departureCity := c.Request.FormValue("departure_city")
	arrivalCity := c.Request.FormValue("arrival_city")
	train, notFound := service.QuerySpecificFlightInfo(departureCity, arrivalCity)
	if notFound {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "查询失败", "data": "没有找到该航班"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": train})
	}
}

// ListAllTrains doc
// @description 获取所有火车信息，返回列表
// @Tags 出行
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有火车信息""}"
// @Router /travel/list_all_trains [GET]
func ListAllTrains(c *gin.Context) {
	trainList := service.QueryAllTrains()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": trainList})
}

// ListAllTrainInfo doc
// @description 获取所有火车信息，返回列表 [更新]
// @Tags 出行
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有火车信息""}"
// @Router /travel/list_all_train_info [GET]
func ListAllTrainInfo(c *gin.Context) {
	trainList := service.QueryAllTrainInfo()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": trainList})
}

// FindSpecificTrainInfo doc
// @description 根据起始地查询列车信息，返回列表
// @Tags 出行
// @Param departure_city formData string true "出发地"
// @Param arrival_city formData string true "到达地"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"查询到的火车信息""}"
// @Failue 200 {string} string "{"success":true, "message":"查询成功","data":"没有找到该列车""}"
// @Router /travel/find_specific_train_info [POST]
func FindSpecificTrainInfo(c *gin.Context) {
	departureCity := c.Request.FormValue("departure_city")
	arrivalCity := c.Request.FormValue("arrival_city")
	train, notFound := service.QuerySpecificTrainInfo(departureCity, arrivalCity)
	if notFound {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "查询失败", "data": "没有找到该列车"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": train})
	}
}

// ListAllCities doc
// @description 获取所有主要城市，返回列表
// @Tags 出行
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有主要城市列表""}"
// @Router /travel/list_all_main_cities [GET]
func ListAllCities(c *gin.Context) {
	cityList := service.QueryAllMainCity()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": cityList})
}

// ListAllCities doc
// @description 查询城市的中心位置坐标 [模糊搜索]
// @Tags 出行
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"城市坐标""}"
// @Failure 200 {string} string "{"success":false, "message":"查询失败","data":"没有找到该城市""}"
// @Router /travel/find_center_city_coordinate [POST]
func FindCenterCityCoordinate(c *gin.Context) {
	cityName := c.Request.FormValue("name")
	city, notFound := service.QueryCenterCityCoordinate(cityName)
	if notFound {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "查询失败", "data": "没有找到该城市"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": city})
	}

}

//----------------------------------------------------
//----------------------------------------------------
//------------------------新冠-------------------------
//----------------------------------------------------
//----------------------------------------------------

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
// @description 获取所有地区的新冠感染/死亡/治愈/疫苗接种人数 [信息综合]，返回列表
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
// @description 获取所有地区的新冠感染/死亡/治愈/疫苗接种人数 [信息综合]，返回列表 [根据时间分组]
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
// @description 获取所有地区的新冠感染/死亡/治愈 [信息综合]，返回列表 [根据时间分组] [Province]
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

// ListCountryOverviewData doc
// @description 获取中国或其它某个国家的各类整体数据，以及国家二级行政单位的现存确诊、新增确诊、累积确诊、累计及新增新冠感染/死亡/治愈 [信息综合]，返回列表 [根据日期-省份分组] [Province]
// @Tags 数据
// @Param country formData string true "国家名"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":{本部分格式见ChinaAnalysisSample.json}"
// @Router /data/list_country_overview [POST]
func ListCountryOverviewData(c *gin.Context) {
	countryName := c.Request.FormValue("country")
	var data []model.CountryOverviewAndDetails
	if countryName == "China" {
		data, _ = service.QueryChinaOverviewAndDetails()
		var dataRevert []model.CountryOverviewAndDetails
		length := len(data)
		for i := (length - 1); i >= 0; i-- {
			dataRevert = append(dataRevert, data[i])
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": dataRevert})
		return
	} else {
		data = service.QueryOtherCountryOverviewAndDetails(countryName)
		var dataRevert []model.CountryOverviewAndDetails
		length := len(data)
		for i := (length - 1); i >= 0; i-- {
			dataRevert = append(dataRevert, data[i])
		}
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": dataRevert})
		return
	}
}

// ListProvinceOverviewData doc
// @description 获取中国某个省份的各类整体数据，以及省份下各市的现存确诊、新增确诊、累积确诊、累计及新增新冠感染/死亡/治愈 [信息综合]，返回列表 [根据日期-市分组] [Province]
// @Tags 数据
// @Param name formData string true "省份或直辖市的拼音"
// @Param zhname formData string true "省份或直辖市的中文"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":{本部分格式见ChinaAnalysisSample.json}"
// @Router /data/list_province_overview [POST]
func ListProvinceOverviewData(c *gin.Context) {
	provinceNamePinYin := c.Request.FormValue("name")
	provinceName := c.Request.FormValue("zhname")
	data := service.QueryProvinceOverviewAndDetails(provinceName, provinceNamePinYin)
	var dataRevert []model.DistrictOverviewAndDetail
	length := len(data)
	for i := (length - 1); i >= 0; i-- {
		dataRevert = append(dataRevert, data[i])
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": dataRevert})

}

// ListOverviewData doc
// @description 获取世界或中国的现存确诊、新增确诊、累积确诊、累计及新增新冠感染/死亡/治愈 [信息综合]，返回列表 [根据国家分组] [Province]
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","nowcases":{"nownum": 123, "newnum": 123}等数据}"
// @Router /data/list_overview [GET]
func ListOverviewData(c *gin.Context) {
	globalTable, chinaTable := service.QueryOtherCountryOverviewAndDetailsForHomeData()
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "Global": globalTable, "China": chinaTable,
		"globalLen": len(globalTable.Detailed), "chinaLen": len(chinaTable.Detailed)})
}

// ListOverviewHistoryData doc
// @description 获取世界每一天的现存确诊、新增确诊、累积确诊、累计及新增新冠感染/死亡/治愈 [信息综合]，返回列表 [根据日期分组]
// @Tags 数据
// @Success 200 {string} string "{""success":true, "message":"查询成功","data":"结果列表""}"
// @Router /data/list_history_overview [GET]
func ListHistoryOverviewData(c *gin.Context) {
	var globalTableRevert []model.GlobalOverviewAndDetailsWithDate
	globalTable := service.QueryGlobalOverviewAndDetailsHistory()
	length := len(globalTable)
	for i := (length - 1); i >= 0; i-- {
		globalTableRevert = append(globalTableRevert, globalTable[i])
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "Global": globalTableRevert})
}

// ListDistrictOverviewData doc
// @description 获取三级行政单位的现存确诊、新增确诊、累积确诊、累计及新增新冠感染/死亡/治愈 [信息综合]，返回列表 [根据日期分组] [Province]
// @Param district formData string true "三级行政单位名 如长春、白城"
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","nowcases":{"nownum": 123, "newnum": 123}等数据}"
// @Router /data/list_district_overview [POST]
func ListDistrictOverviewData(c *gin.Context) {
	districtName := c.Request.FormValue("district")
	var dataRevert []model.DistrictOverview
	data := service.QueryDistrictOverview(districtName)
	length := len(data)
	for i := (length - 1); i >= 0; i-- {
		dataRevert = append(dataRevert, data[i])
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": dataRevert, "len": len(dataRevert)})
}

// ListVaccineOverviewData doc
// @description 获取世界每一天的疫苗接种数据，以及每一天的各国的疫苗接种数据[信息综合]，返回列表 [根据日期分组]
// @Tags 数据
// @Success 200 {string} string "{""success":true, "message":"查询成功","data":"结果列表""}"
// @Router /data/list_vaccine_overview [GET]
func ListVaccineOverviewData(c *gin.Context) {
	var globalTableRevert []model.GlobalVaccineOverviewAndDetailsWithDate
	globalTable := service.QueryVaccineOverviewData()
	length := len(globalTable)
	for i := (length - 1); i >= 0; i-- {
		globalTableRevert = append(globalTableRevert, globalTable[i])
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "Global": globalTableRevert})
}

// ListCurrentLocationData doc
// @description 获取当前位置的疫情信息：临时写一个支持北京香港和美国的，后续添加其他地区
// @Param country formData string true "国家名"
// @Param province formData string true "省份名"
// @Param city formData string true "城市名"
// @Param district formData string true "区名"
// @Tags 数据
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"所有信息""}"
// @Router /data/current_location_data [POST]
func ListCurrentLocationData(c *gin.Context) {
	countryName := c.Request.FormValue("country")
	provinceName := c.Request.FormValue("province")
	// cityName := c.Request.FormValue("city")
	districtName := c.Request.FormValue("district")

	var data model.CovidDetailCDRProvince
	if districtName != "" {
		data = service.QueryDistrictData("海淀") // covid_hangzhou_cases
	} else if countryName != "中国" {
		data = service.QueryCountryData("United States of America") // covid_cases
	} else if provinceName == "香港特别行政区" {
		data = service.QueryProvinceData("Hong Kong") // covid_china_cases
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": data})
}
