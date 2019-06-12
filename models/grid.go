package models

import (
	"github.com/jinzhu/gorm"
)

type Grid struct {
	gorm.Model
	DisciplineId uint
	Alerts       []Alert `gorm:"foreignkey:GridId"`
}
