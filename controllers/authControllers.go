package controllers

import (
	db "classtime/db"
	"classtime/models"
	u "classtime/utils"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
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
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	resp := db.Login(strings.ToLower(pair[0]), pair[1])
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
