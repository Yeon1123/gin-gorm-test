package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yujy/gin-gorm-rest/config"
	"github.com/yujy/gin-gorm-rest/routes"
)

func main() {
	router := gin.New()
	// call router every start
	config.Connect()
	routes.UserRoute(router)
	router.Run(":80")
}
