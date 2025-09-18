package controller

import (
	"encoding/json"
	"net/http"

	"parking/auth"
)

func AdminProtected(w http.ResponseWriter, r *http.Request) {
	auth := r.Context().Value("admin").(*auth.AdminClaims)

	if auth.Admin.UID == "" {
		http.Error(w, "Not an admin", http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode(auth.Admin)

	//w.Write([]byte("This is %s a protected endpoint. You're authenticated!",auth.Name))
}
