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
	}
	PortalRouter := Router.Group("/portal")
	{
		PortalRouter.POST("/sub", v1.Subscribe)
		PortalRouter.POST("/list_all_subs", v1.ListAllSubscriptions)
		PortalRouter.POST("/del_sub", v1.RemoveSubscription)
		PortalRouter.POST("/question", v1.CreateAQuestion)
		PortalRouter.POST("/comment", v1.CreateAComment)
	}
}
