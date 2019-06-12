package models

import (
	"github.com/jinzhu/gorm"
)

type Alert struct {
	gorm.Model
	GridId uint
}
