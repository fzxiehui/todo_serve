package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NickName string `gorm:"type:varchar(20);not null;default:'';comment:'昵称'" json:"nick_name"`
	Username string `gorm:"type:varchar(20);not null;unique;comment:'用户名'" json:"username"`
	Password string `gorm:"type:varchar(20);not null;comment:'密码'" json:"password"`
}
