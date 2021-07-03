package v1

import (
	"net/http"
	"strconv"

	"github.com/TualatinX/durian-go/model"
	"github.com/TualatinX/durian-go/service"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "gcp"})
}

// Register doc
// @description 注册
// @Tags user
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "{"success": true, "message": "用户创建成功"}"
// @Failure 400 {string} string "{"success": false, "message": "用户已存在"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	user := model.User{Username: username, Password: password}
	_, notFound := service.QueryAUserByUsername(username)
	if notFound {
		service.CreateAUser(&user)
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "用户创建成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "用户已存在"})
	}
}

// Register doc
// @description 登录
// @Tags user
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
// @description 修改用户信息
// @Tags user
// @Param user_id formData string true "用户ID"
// @Param username formData string true "用户名"
// @Param password1 formData string true "原密码"
// @Param password2 formData string true "新密码"
// @Param email formData string true "邮箱"
// @Param info formData string true "个人信息"
// @Success 200 {string} string "{"success": true, "message": "登录成功", "data": "model.User的所有信息"}"
// @Router /user/modify [post]
func ModifyUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	username := c.Request.FormValue("username")
	password1 := c.Request.FormValue("password1")
	password2 := c.Request.FormValue("password2")
	email := c.Request.FormValue("email")
	info := c.Request.FormValue("info")
	user, notFoundUserByID := service.QueryAUserByID(userID)
	if notFoundUserByID {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户ID不存在",
		})
		return
	}
	if password1 != user.Password {
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
	_, notFoundUserByEmail := service.QueryAUserByEmail(email)
	if !notFoundUserByEmail && email != user.Email {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "邮箱已被占用",
		})
		return
	}
	err := service.UpdateAUser(&user, username, password2, email, info)
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
		"message": "用户信息修改成功",
		"data":    user,
	})
	return
}

// TellUserInfo doc
// @description 查看用户个人信息
// @Tags user
// @Param user_id formData string true "用户ID"
// @Success 200 {string} string "{"success": true, "message": "查看用户信息成功", "data": "model.User的所有信息"}"
// @Router /user/info [post]
func TellUserInfo(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user, _ := service.QueryAUserByID(userID)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "查看用户信息成功",
		"data":    user,
	})
	return
}

// AddToFavorites doc
// @description 添加收藏
// @Tags user
// @Param user_id formData string true "用户ID"
// @Param paper_id formData string true "文献ID"
// @Param paper_info formData string true "文献描述"
// @Success 200 {string} string "{"success":true, "message":"收藏成功"}"
// @Router /user/favorite/add [post]
func AddToFavorites(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	paperID := c.Request.FormValue("paper_id")
	paperInfo := c.Request.FormValue("paper_info")

	if userID == 0 {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "请先登录"})
		return
	}

	if err := service.CreateAFavorite(userID, paperID, paperInfo); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "收藏成功"})
	}
}

// ListAllFavorites doc
// @description 获取收藏列表
// @Tags user
// @Param user_id formData string true "用户ID"
// @Success 200 {string} string "{"success":true, "message":"查询成功","data":"user的所有收藏"}"
// @Router /user/favorite/list [post]
func ListAllFavorites(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	favorites := service.QueryAllFavorites(userID)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "查询成功", "data": favorites})
	return
}

// RemoveFavorite doc
// @description 移除收藏
// @Tags user
// @Param favor_id formData string true "收藏ID"
// @Success 200 {string} string "{"success":true, "message":"删除成功"}"
// @Router /user/favorite/remove [post]
func RemoveFavorite(c *gin.Context) {
	favorID, _ := strconv.ParseUint(c.Request.FormValue("favor_id"), 0, 64)
	if err := service.DeleteAFavorite(favorID); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "删除成功"})
	}
	return
}
