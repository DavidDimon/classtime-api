package models

import (
	"github.com/jinzhu/gorm"
)

type Discipline struct {
	gorm.Model
	Name     string  `json:"name"`
	Term     string  `json:"term"` // term = semester
	Hashcode string  `json:"hashcode"`
	Users    []*User `gorm:"many2many:user_disciplines;"`
	Grid     Grid    `gorm:"foreignkey:DisciplineId"`
}

/*
DisciplineJSON struct used by update discipline with users ids
*/
type DisciplineJSON struct {
	Name        string
	Term        string
	Users       []uint
	UsersRemove []uint
	GridID      uint
}
