package initialize

import (
	v1 "github.com/TualatinX/durian-go/api/v1"
	"github.com/TualatinX/durian-go/middleware"
	"github.com/TualatinX/durian-go/router"

	"github.com/gin-gonic/gin"
)

// 配置组路由
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/", v1.Index)
	r.GET("/test", v1.ListAllCovidCasesResponse)
	Group := r.Group("api/v1/")
	{
		router.InitRouter(Group)
	}
	return r
}
