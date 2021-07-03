package initialize

import (
	"github.com/TualatinX/durian-go/middleware"
	"github.com/TualatinX/durian-go/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	Group := r.Group("api/v1/")
	{
		router.InitRouter(Group)
	}
	return r
}
