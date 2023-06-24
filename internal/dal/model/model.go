package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string `gorm:"type:varchar(20);not null;default:'';comment:'昵称'" json:"nick_name"`
	Username string `gorm:"type:varchar(32);not null;unique;comment:'用户名'" json:"username"`
	Password string `gorm:"type:varchar(128);not null;comment:'密码'" json:"password"`
}

type Todo struct {
	gorm.Model
	Date    string `gorm:"type:varchar(20);not null;comment:'日期'" json:"date"`
	Content string `gorm:"type:MEDIUMTEXT;not null;comment:'内容'" json:"content"`
	Done    bool   `gorm:"type:tinyint(1);not null;default:0;comment:'是否完成'" json:"done"`
	UserId  uint   `gorm:"type:int;not null;comment:'用户id'" json:"user_id"`
	User    User   `gorm:"foreignKey:UserId" json:"user"`
}
