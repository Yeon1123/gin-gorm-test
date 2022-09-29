package config

import (
	"github.com/yujy/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connect db
// gorm docs
var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:inspace123@localhost:5432"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//user 이름의 DB가 생성된다! 이미 있는 경우 구조체와 필드를 비교해 새 컬럼을 추가한다. 신기한 기능
	db.AutoMigrate(&models.MemberGo{})
	DB = db

}
