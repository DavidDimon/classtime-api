package controllers

import (
	db "classtime/db"
	"classtime/models"
	u "classtime/utils"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.UserCreate{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := db.CreateUser(user) //Create user
	u.Respond(w, resp)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
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
	if resp["status"] == false {
		w.WriteHeader(http.StatusUnauthorized)
	}

	u.Respond(w, resp)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if HasPermission(w, r, 2) == false {
		return
	}
	params := mux.Vars(r)
	id := params["id"]
	data := db.GetUser(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if HasPermission(w, r, 2) == false {
		return
	}

	data := db.GetUsers()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
