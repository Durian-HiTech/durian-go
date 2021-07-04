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
	NoticeRouter := Router.Group("/notice")
	{
		NoticeRouter.POST("/create_question", v1.CreateAQuestion)
		NoticeRouter.POST("/create_comment", v1.CreateAComment)
		NoticeRouter.POST("/list_all_comments", v1.ListAllComments)
	}
	SubRouter := Router.Group("/sub")
	{
		SubRouter.POST("/subscribe", v1.Subscribe)
		SubRouter.POST("/list_all_subs", v1.ListAllSubscriptions)
		SubRouter.POST("/del_sub", v1.RemoveSubscription)
	}
	NewsRouter := Router.Group("/news")
	{
		NewsRouter.GET("/list_all_news", v1.ListAllNews)
		NewsRouter.POST("/detail", v1.ViewNewsDetail)
	}
}
