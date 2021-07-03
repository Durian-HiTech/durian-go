package router

import (
	v1 "github.com/TualatinX/durian-go/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(Router *gin.RouterGroup) {
	Router.GET("/", v1.Index)
	UserRouter := Router.Group("/user")
	{
		UserRouter.POST("/register", v1.Register)
		UserRouter.POST("/login", v1.Login)
		UserRouter.POST("/modify", v1.ModifyUser)
		UserRouter.POST("/info", v1.TellUserInfo)
		UserRouter.POST("/subscribe", v1.Subscribe)
		UserRouter.POST("/list_all_subscriptions", v1.ListAllSubscriptions)
		UserRouter.POST("/remove", v1.RemoveSubscription)
	}
}
