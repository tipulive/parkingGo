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
	/*perm := helper.ConvertStringToJson(AuthAdmin(r).Admin.PermissionJSON)
	if perm["test"] != "true" { //true means show,empty or false means hide
		json.NewEncoder(w).Encode(response(false, "you Do not have a permission to Access This Option Please Contact System Admin"))
		return
	}*/
	perm := helper.ConvertStringToJson(AuthAdmin(r).Admin.PermissionJSON)
	json.NewEncoder(w).Encode(response(true, perm["test"]))
	//
	//json.NewEncoder(w).Encode(response(true, AuthAdmin(r)))

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
	//_, err := db.DB.Exec(query, helper.GetFileName(), adminData.Email, 2) //Add Permission
	permissionkey := "app"
	_, err := db.DB.Exec(query, permissionkey, adminData.Email, 2) //Add Permission
	if err != nil {
		log.Fatal(err)
	}
	Claims := AuthAdmin(r)

	permMap := helper.ConvertStringToJson(AuthAdmin(r).Admin.PermissionJSON)
	permMap[permissionkey] = adminData.Email
	Claims.Admin.PermissionJSON = helper.ConvertJsonToString(permMap)
	//
	//Claims.Admin.PermissionJSON = "hello"
	tokeString, _ := auth.TokenSigned(auth.AdminJwtKey, Claims)

	json.NewEncoder(w).Encode(response(true, tokeString))

}
func EditPermission_Admin() {

}
