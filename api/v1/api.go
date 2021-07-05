package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/TualatinX/durian-go/model"
	"github.com/TualatinX/durian-go/service"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "gcp"})
}

// Register doc
// @description 注册
// @Tags 用户管理
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param user_type formData string true "用户类型（0: 普通用户，1: 认证机构用户）"
// @Param affiliation formData string false "认证机构名"
// @Success 200 {string} string "{"success": true, "message": "用户创建成功"}"
// @Failure 400 {string} string "{"success": false, "message": "用户已存在"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	userType, _ := strconv.ParseUint(c.Request.FormValue("user_type"), 0, 64)
	affiliation := c.Request.FormValue("affiliation")
	user := model.User{Username: username, Password: password, UserType: userType, Affiliation: affiliation}
	_, notFound := service.QueryAUserByUsername(username)
	if notFound {
		service.CreateAUser(&user)
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "用户创建成功"})
	} else {
		c.JSON(400, gin.H{"success": false, "message": "用户已存在"})
	}
}

// Login doc
// @description 登录
// @Tags 用户管理
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "{"success": true, "message": "登录成功", "detail": user的信息}"
// @Failure 400 {string} string "{"success": false, "message": "密码错误"}"
// @Failure 404 {string} string "{"success": false, "message": "没有该用户"}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	user, notFound := service.QueryAUserByUsername(username)
	if notFound {
		c.JSON(404, gin.H{"success": false, "message": "没有该用户"})
	} else {
		if user.Password != password {
			c.JSON(400, gin.H{"success": false, "message": "密码错误"})
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
// @Failure 400 {string} string "{"success": false, "message": "原密码输入错误"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /user/modify [post]
func ModifyUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	username := c.Request.FormValue("username")
	password_old := c.Request.FormValue("password_old")
	password_new := c.Request.FormValue("password_new")
	user, notFoundUserByID := service.QueryAUserByID(userID)
	if notFoundUserByID {
		c.JSON(404, gin.H{
			"success": false,
			"message": "用户ID不存在",
		})
		return
	}
	if password_old != user.Password {
		c.JSON(400, gin.H{
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
	err := service.UpdateAUser(&user, username, password_new)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
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
// @Router /user/info [post]
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
// @Failure 404 {string} string "{"success": false, "message": "已经订阅过这个城市的疫情信息"}"
// @Failure 401 {string} string "{"success": false, "message": "数据库error, 一些其他错误"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /sub/subscribe [post]
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
		c.JSON(400, gin.H{
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
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"user的所有订阅"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /sub/list_all_subs [post]
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
// @Param subscription_id formData string true "订阅ID"
// @Success 200 {string} string "{"success":true, "message":"删除成功"}"
// @Failure 401 {string} string "{"success": false, "message": "数据库error, 一些其他错误"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /sub/del_sub [post]
func RemoveSubscription(c *gin.Context) {
	subscriptionID, _ := strconv.ParseUint(c.Request.FormValue("subscription_id"), 0, 64)

	_, notFoundSubscription := service.QueryASubscriptionByID(subscriptionID)
	if notFoundSubscription {
		c.JSON(404, gin.H{
			"success": false,
			"message": "这一订阅ID不存在",
		})
		return
	}

	if err := service.DeleteASubscription(subscriptionID); err != nil {
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
// @Success 200 {string} string "{"success": true, "message": "用户提问成功", "detail": 提问的全部信息}"
// @Failure 401 {string} string "{"success": false, "message": "数据库error, 一些其他错误"}"
// @Failure 404 {string} string "{"success": false, "message": "用户ID不存在"}"
// @Router /notice/create_question [post]
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

	question_title := c.Request.FormValue("question_title")
	question_content := c.Request.FormValue("question_content")

	question := model.Question{UserID: userID, QuestionTitle: question_title, QuestionContent: question_content, QuestionTime: time.Now()}
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

// ListAllComments doc
// @description 列出某个问题的全部评论
// @Tags 防控知识板块
// @Param question_id formData string true "问题ID"
// @Success 200 {string} string "{"success": true, "message": "查看成功", "data": "某问题的所有评论"}"
// @Failure 404 {string} string "{"success": false, "message": "问题ID不存在"}"
// @Router /notice/list_all_comments [post]
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
// @Router /notice/create_comment [post]
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
	comment_content := c.Request.FormValue("comment_content")
	comment := model.Comment{UserID: userID, QuestionID: questionID, CommentContent: comment_content, CommentTime: time.Now(), UserType: userType}
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
