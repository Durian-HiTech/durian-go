package router

import (
	v1 "github.com/TualatinX/durian-go/api/v1"

	"github.com/gin-gonic/gin"
)

// 初始化路由
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
		NoticeRouter.GET("/list_all_notice", v1.ListAllNotice)
		NoticeRouter.POST("/detail", v1.ViewNoticeDetail)
		NoticeRouter.POST("/create_question", v1.CreateAQuestion)
		NoticeRouter.POST("/list_a_question", v1.ListAQuestion)
		NoticeRouter.POST("/create_comment", v1.CreateAComment)
		NoticeRouter.POST("/list_all_comments", v1.ListAllComments)
		NoticeRouter.GET("/list_all_questions", v1.ListAllQuestions)
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
	DataRouter := Router.Group("/data")
	{
		DataRouter.GET("/list_all_high_risk_areas", v1.ListHighRiskAreas)
		DataRouter.POST("/query_data", v1.FetchRequiredData)
		DataRouter.GET("/list_all_covid_cases", v1.ListAllCovidCases)
		DataRouter.GET("/list_all_covid_deaths", v1.ListAllCovidDeaths)
		DataRouter.GET("/list_all_covid_recovereds", v1.ListAllCovidRecovereds)
		DataRouter.GET("/list_all_covid_vaccines", v1.ListAllCovidVaccines)
		DataRouter.GET("/list_all_covid_cdrv", v1.ListAllCovidCDRV)
	}
}
