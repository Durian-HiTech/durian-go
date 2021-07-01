package router

import (
	v1 "github.com/TualatinX/durian-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("api/v1/user")
	{
		UserRouter.POST("/register", v1.Register)

	}
}
