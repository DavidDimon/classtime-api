package models

import (
	u "classtime/utils"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

type Discipline struct {
	gorm.Model
	Name          string   `json:"name"`
	Term          string   `json:"term"` // term = semester
	Hashcode      string   `json:"hashcode"`
	Users         []*User  `gorm:"many2many:user_disciplines;"`
	Grid          Grid     `gorm:"foreignkey:DisciplineId"`
	WeekDays      string   `json:"-"`
	WeekDaysArray []string `gorm:"-" json:"weekDays"`
	Classroom     string   `json:"classroom"`
}

/*
DisciplineJSON struct used by update discipline with users ids
*/
type DisciplineJSON struct {
	Name        string
	Term        string
	Users       []uint
	UsersRemove []uint
	WeekDays    []string
	GridID      uint
	Classroom   string
}

/*
FindDay by string day
*/
func FindDay(search string) (string, int) {
	for index, value := range days {
		if strings.ToLower(value) == strings.ToLower(search) {
			return value, index
		}
	}
	return "Invalid day", 0
}

/*
ParseWeekDays weekdays = days to save
parse into 1,2,3
example: days = ["Sunday","Monday"]
result = "0,1"
*/
func ParseWeekDays(weekDays []string) string {
	var result []int
	for _, value := range weekDays {
		day, index := FindDay(value)
		if day != "Invalid day" {
			result = append(result, index)
		}
	}
	return u.ArrayToString(result, ",")
}

/*
GetDays all days by weekdays from db
weekDays = "0,5"
result = ["Sunday", "Friday"]
*/
func GetDays(weekDays string) []string {
	var result []string
	weekDaysArray := strings.Split(weekDays, ",")
	for _, value := range weekDaysArray {
		index, _ := strconv.Atoi(value)
		day, _ := FindDay(days[index])
		if day != "Invalid day" {
			result = append(result, day)
		}
	}
	return result
}
