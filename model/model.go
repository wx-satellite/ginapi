package model

import (
	"time"
)

type BaseModel struct {
	// "-"返回前端会被隐藏
	Id        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}


