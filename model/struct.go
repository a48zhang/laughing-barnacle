package model

import "time"

type User struct {
	Email          string
	PasswordDigest string     //密码加密后的密文
	Gender         string     //性别
	Birthday       *time.Time //生日
	Signature      string     //个性签名
}
