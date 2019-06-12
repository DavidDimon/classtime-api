package models

import (
	"github.com/jinzhu/gorm"
)

type Discipline struct {
	gorm.Model
	Name  string  `json:"string"`
	Term  string  `json:"string"` // term = semester
	Users []*User `gorm:"many2many:user_disciplines;"`
	Grid  Grid    `gorm:"foreignkey:DisciplineId"`
}
