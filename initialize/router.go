package initialize

import (
	v1 "github.com/TualatinX/durian-go/api/v1"
	"github.com/TualatinX/durian-go/middleware"
	"github.com/TualatinX/durian-go/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/", v1.Index)
	Group := r.Group("api/v1/")
	{
		router.InitRouter(Group)
	}
	return r
}
