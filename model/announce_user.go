package model

import "gorm.io/gorm"

type AnnounceUser struct {
	gorm.Model
	Aid    uint `gorm:"not null"`
	Uid    uint `gorm:"not null"`
	Status int  `gorm:"default:0"` //已读状态
}
