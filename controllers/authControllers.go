package controllers

import (
	db "classtime/db"
	"classtime/models"
	u "classtime/utils"
	"encoding/json"
	"net/http"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := db.Create(user) //Create user
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := db.Login(user.Email, user.Password)
	u.Respond(w, resp)
}

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := db.GetUser(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetUsers = func(w http.ResponseWriter, r *http.Request) {
	data := db.GetUsers()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
