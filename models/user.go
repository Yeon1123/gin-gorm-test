package models

import (
	"gorm.io/gorm"
)

// Spring Boot에서 생성한 json - 객체 간 변환도 요렇게 가능하다!
type User struct {
	gorm.Model
	Id       int    `json:"member_id" gorm:"primary_key"`
	Name     string `json:"member_name"`
	Email    string `json:"member_email"`
	Password string `json:"member_password"`
}
