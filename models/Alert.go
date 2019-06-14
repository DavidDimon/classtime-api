package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Alert struct {
	gorm.Model
	GridId uint       `json:"gridId"`
	Date   *time.Time `json:"date"`
}
