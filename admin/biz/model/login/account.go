package login

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Id     int64  `json:"id" column:"id"`
	Name   string `json:"name" column:"name"`
	Passwd string `json:"passwd" column:"passwd"`
	Status int    `json:"status" column:"status"`
}

func (u *Account) TableName() string {
	return "account"
}
