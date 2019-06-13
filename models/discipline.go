package models

import (
	"github.com/jinzhu/gorm"
)

type Discipline struct {
	gorm.Model
	Name  string  `json:"name"`
	Term  string  `json:"term"` // term = semester
	Users []*User `gorm:"many2many:user_disciplines;"`
	Grid  Grid    `gorm:"foreignkey:DisciplineId"`
}
