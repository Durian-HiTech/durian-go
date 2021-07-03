package v1

import (
	"net/http"

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
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "用户已存在"})
	}
}

func Login(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	user, notFound := service.QueryAUserByUsername(username)
	if notFound {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "没有该用户"})
	} else {
		if user.Password != password {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "密码错误"})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "登录成功"})
		}
	}
}
