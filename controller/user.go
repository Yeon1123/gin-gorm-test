package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yujy/gin-gorm-rest/config"
	"github.com/yujy/gin-gorm-rest/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ShowAccount godoc
// @Summary      Show all users
// @Description  Get user list
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {array} models.MemberGo
// @Router       / [get]
func GetUsers(c *gin.Context) {
	users := []models.MemberGo{}
	config.DB.Find(&users)
	c.JSON(200, users)
}

func CreateUser(createdUser models.MemberGo) *gorm.DB {
	var user = createdUser
	err := config.DB.Create(&user)
	return err
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

// ShowAccount godoc
// @Summary      Create User
// @Description  Create a new user account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param 		 request body models.MemberInfo true "request member"
// @Success      201 {object} models.MemberInfo
// @Router       /signup [post]
func Signup(c *gin.Context) {
	var newUserInfo models.MemberInfo
	var memberGo models.MemberGo

	if err := c.BindJSON(&newUserInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	//이메일 중복검사
	result := config.DB.Find(&memberGo, "member_email=?", newUserInfo.MemberEmail)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	//비밀번호 암호화
	hashpw, err := hashPassword(newUserInfo.MemberPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
		return
	}

	memberGo.MemberInfo = newUserInfo
	memberGo.MemberInfo.MemberPassword = hashpw

	if err := CreateUser(memberGo); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred creating user"})
		return
	}

	c.JSON(http.StatusCreated, memberGo.MemberInfo)
}

// 암호화 로직
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(hashVal, userPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))
	if err != nil {
		return false
	} else {
		return true
	}
}

func SignIn(c *gin.Context) {
	requestUser := new(models.MemberInfo)
	dbUser := new(models.MemberGo)

	if err := c.BindJSON(&requestUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad Request"})
		return
	}

	result := config.DB.Find(&dbUser, "member_email=?", requestUser.MemberEmail)
	if result.RowsAffected <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email not found"})
		return
	}

	if !checkPasswordHash(dbUser.MemberInfo.MemberPassword, requestUser.MemberPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password does not match"})
		return
	}

	c.JSON(http.StatusOK, dbUser.MemberInfo)
}
