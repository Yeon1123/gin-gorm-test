package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yujy/gin-gorm-rest/config"
	docs "github.com/yujy/gin-gorm-rest/docs"
	"github.com/yujy/gin-gorm-rest/routes"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return router
}

func main() {
	router := SetupRouter()
	setSwagger()
	// call router every start
	config.Connect()
	routes.UserRoute(router)
	ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":80")
}

func setSwagger() {
	docs.SwaggerInfo.Title = "Go Test Docs"
	docs.SwaggerInfo.Description = "This is a sample docs using go language."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
