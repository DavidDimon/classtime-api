package controllers

import (
	db "classtime/db"
	"classtime/models"
	u "classtime/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func AddAlert(w http.ResponseWriter, r *http.Request) {
	user := db.GetUserAuthenticated(r)
	alert := &models.Alert{}
	params := mux.Vars(r)
	id := params["id"]
	err := json.NewDecoder(r.Body).Decode(alert)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := db.AddAlert(id, user, alert)
	u.Respond(w, resp)
}

func RemoveAlert(w http.ResponseWriter, r *http.Request) {
	user := db.GetUserAuthenticated(r)
	alert := &models.Alert{}
	params := mux.Vars(r)
	id := params["id"]
	err := json.NewDecoder(r.Body).Decode(alert)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := db.RemoveAlert(id, user)
	if resp["status"] == false {
		w.WriteHeader(http.StatusForbidden)
	}
	u.Respond(w, resp)
}

func GetGrid(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data := db.GetGrid(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
