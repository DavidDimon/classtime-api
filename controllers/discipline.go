package controllers

import (
	db "classtime/db"
	"classtime/models"
	u "classtime/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateDiscipline(w http.ResponseWriter, r *http.Request) {
	if HasPermission(w, r, 2) == false {
		return
	}

	discipline := &models.DisciplineJSON{}
	err := json.NewDecoder(r.Body).Decode(discipline)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := db.CreateDiscipline(discipline)
	u.Respond(w, resp)
}

func UpdateDiscipline(w http.ResponseWriter, r *http.Request) {
	// if HasPermission(w, r, 2) == false {
	// 	return
	// }
	discipline := &models.DisciplineJSON{}
	params := mux.Vars(r)
	id := params["id"]
	err := json.NewDecoder(r.Body).Decode(discipline)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := db.UpdateDiscipline(id, discipline)
	u.Respond(w, resp)
}

func GetDisciplines(w http.ResponseWriter, r *http.Request) {
	user := db.GetUserAuthenticated(r)
	data := db.GetDisciplines(user)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
