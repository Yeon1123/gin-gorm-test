package models

import (
	"gorm.io/gorm"
)

// Spring Boot에서 생성한 json - 객체 간 변환도 요렇게 가능하다!
// 생성시 자동으로 snake_case 되도록 설정
type MemberGo struct {
	gorm.Model
	MemberName     string `json:"member_name"`
	MemberEmail    string `json:"member_email"`
	MemberPassword string `json:"member_password"`
}
