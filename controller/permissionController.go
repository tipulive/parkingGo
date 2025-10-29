package controller

import (
	//"fmt"

	"encoding/json"
	"log"
	"net/http"
	"parking/auth"
	"parking/db"
	"parking/helper"
	//"parking/db"
	//"parking/helper"
)

func LoadPermission(w http.ResponseWriter, r *http.Request) {
	/*
	   err := db.DB.Select(&packages, "SELECT id, name, uid FROM packages WHERE subscriber=? and name=? limit 1",

	   		authData.Admin.Subscriber, Name)

	   	if err != nil {
	   		// This means actual DB error (connection, query issue, etc.)
	   		json.NewEncoder(w).Encode(response(false, err.Error()))
	   		return
	   	}
	*/
	/*authData := AuthAdmin(r)
	if authData.Admin.PermissionJSON == "null" {
		json.NewEncoder(w).Encode(response(false, "nil"))
		return

	}*/
	jsonStr := AuthAdmin(r).Admin.PermissionJSON
	var result map[string]string
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		json.NewEncoder(w).Encode(response(true, err.Error()))
		return
	}
	json.NewEncoder(w).Encode(response(true, result[helper.GetFileName()]))

	//fmt.Println(authData.Admin.PermissionJSON)
}
func AddPermission(w http.ResponseWriter, r *http.Request) {
	if ok := helper.ValidateAndDecodeJSON(w, r, &adminData, validateData); !ok {
		return // helper already wrote response
	}
	query := `
	UPDATE admins
	SET permissionJson = JSON_MERGE_PATCH(
	  COALESCE(permissionJson, '{}'),
	  JSON_OBJECT(?, ?)
	)
	WHERE id = ?;
	`

	// 5️⃣ Save it back
	_, err := db.DB.Exec(query, helper.GetFileName(), adminData.Email, 2) //Add Permission
	if err != nil {
		log.Fatal(err)
	}
	ctx := auth.ClaimUpdate(r.Context(), auth.AdminClaimsKey, func(c *auth.AdminClaims) {
		c.Admin.PermissionJSON = adminData.Email
	})
	r = r.WithContext(ctx)
	// no next handler available in this HTTP handler; continue processing without calling next
	/*ctx := context.WithValue(r.Context(), "admin", adminCla)
	next.ServeHTTP(w, r.WithContext(ctx))*/
	// 1️⃣ create base context

	json.NewEncoder(w).Encode(map[string]any{
		"admin": AuthAdmin(r),
		"user":  adminData.Email,
	})
}
func EditPermission_Admin() {

}
