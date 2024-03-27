package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	Uid uint `gorm:"not null"`
	Fid uint `gorm:"not null"`
}
