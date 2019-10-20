package byn


import (
	"gin/model"
)


type Student struct {
	model.BaseModel
	Username string `json:"username" gorm:"column:username;not null;" validate:"required" `
	Password string `json:"password" gorm:"column:password;not null;" validate:"required" `
}



func (student *Student) TableName() string {
	return "students"
}



