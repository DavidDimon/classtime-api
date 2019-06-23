package models

import (
	u "classtime/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

/*
	roles:
	0 - student
	1 - teacher
	2 - secretary
	3 - admin
*/

//a struct to rep user account
type User struct {
	gorm.Model
	Email       string        `json:"email"`
	Password    string        `json:"-"`
	Name        string        `json:"name"`
	Token       string        `json:"token" sql:"-"`
	Disciplines []*Discipline `gorm:"many2many:user_disciplines;"`
	Role        uint          `json:"role"`
	StudentID   string        `json:"studentId"`
}

//UserCreate is struct to create a new user
type UserCreate struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	StudentID string `json:"studentId"`
}

func ValidateEmail(user *User) (map[string]interface{}, bool) {
	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}
	return u.Message(false, "success"), true
}
