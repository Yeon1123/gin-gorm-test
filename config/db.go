package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/yujy/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	enverr := godotenv.Load(".env")
	
	if enverr != nil {
		log.Fatal("Error Loading environment: ", enverr)
	}

	//환경설정 사용자 정보 가져오기
	dbname := os.Getenv("DBNAME")
	dbuser := os.Getenv("DBUSER")
	dbpassword := os.Getenv("DBPASSWORD")

	logConfig := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	connect := "host=localhost user=" + dbuser + " password=" + dbpassword + " dbname=" + dbname + " port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connect), &gorm.Config{
		Logger: logConfig,
	})
	
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.MemberGo{})
	DB = db
}
