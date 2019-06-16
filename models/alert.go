package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Alert struct {
	gorm.Model
	GridId   uint       `json:"gridId"`
	Message  string     `json:"message"`
	Username string     `json:"user"`
	UserId   uint       `json:"-"`
	Date     *time.Time `json:"date"`
}
