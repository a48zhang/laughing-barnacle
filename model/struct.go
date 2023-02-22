package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email          string `gorm:"unique"`
	PasswordDigest string //密码加密后的密文
	//TODO
	Gender    string     //性别
	Birthday  *time.Time //生日
	Signature string     //个性签名
}

// Memo 备忘录内容
type Memo struct {
	gorm.Model
	uid     int
	Color   string
	Content string `gorm:"type:longtext"` //长字符串
}
