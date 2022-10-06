package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yujy/gin-gorm-rest/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
	router.POST("/signup", controller.Signup)
}
