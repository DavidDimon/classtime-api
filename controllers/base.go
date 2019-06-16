package controllers

import (
	db "classtime/db"
	u "classtime/utils"
	"net/http"
)

func HasPermission(w http.ResponseWriter, r *http.Request, role uint) bool {
	user := db.GetUserAuthenticated(r)
	if user.Role < role {
		w.WriteHeader(http.StatusForbidden)
		u.Respond(w, u.Message(false, "Permission denied"))
		return false
	}
	return true
}
