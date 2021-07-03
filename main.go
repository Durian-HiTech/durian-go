package main

import (
	"github.com/TualatinX/durian-go/docs"
	"github.com/TualatinX/durian-go/initialize"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Durian Covid-19 Golang Backend
// @version 1.0
// @description Durian HiTech
// @schemes http https
func main() {
	docs.SwaggerInfo.Title = "Durian HiTech"
	docs.SwaggerInfo.Description = "This is Durian's Covid-19 Golang Backend"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	initialize.InitMySQL()

	r := initialize.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
