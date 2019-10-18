package byn


import (
	"gin/model"
)

type Student struct {
	model.BaseModel
	Username string `json:"username" gorm:"column:username;not null;" binding:"required" validate:"min=1,max=128"`
	Password string `json:"password" gorm:"column:password;not null;" binding:"required" validate:"min=6,max=32"`
}



func (student *Student) TableName() string {
	return "student"
}



