package models

import (
	"time"
)

// Spring Boot에서 생성한 json - 객체 간 변환도 요렇게 가능하다!
// 생성시 자동으로 snake_case 되도록 설정

// 임베디드 객체 사용
type MemberGo struct {
	ID         uint64     `json:"member_id" gorm:"primaryKey;autoIncrement:true"`
	MemberInfo MemberInfo `gorm:"embedded"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type MemberInfo struct {
	MemberName     string `json:"member_name"`
	MemberEmail    string `json:"member_email"`
	MemberPassword string `json:"member_password"`
}
