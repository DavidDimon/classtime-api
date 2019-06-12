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

//a struct to rep user account
type User struct {
	gorm.Model
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	Token       string        `json:"token";sql:"-"`
	Disciplines []*Discipline `gorm:"many2many:user_disciplines;"`
	Role        string        `json:"string"`
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
