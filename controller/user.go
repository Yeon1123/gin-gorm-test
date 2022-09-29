package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yujy/gin-gorm-rest/config"
	"github.com/yujy/gin-gorm-rest/models"
)

// get user
func GetUsers(c *gin.Context) {
	users := []models.MemberGo{}
	config.DB.Find(&users)
	c.JSON(200, users)
	// c.String(200, "Hello world!")
}

func CreateUser(c *gin.Context) {
	var user models.MemberGo
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(201, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.MemberGo
	config.DB.Where("id=?", c.Param("id")).Delete(&user)
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	var user models.MemberGo
	config.DB.Where("id=?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
