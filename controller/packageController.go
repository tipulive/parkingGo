package controller

import (
	//"database/sql"
	"encoding/json"
	"strings"

	//"log"

	"parking/db"
	"parking/helper"

	"net/http"
)

func Package(w http.ResponseWriter, r *http.Request) {

	err := db.DB.Select(&packages, "Select *from packages ")

	if err != nil {
		// This means actual DB error (connection, query issue, etc.)
		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}
	json.NewEncoder(w).Encode(response(true, packages))
}
func AddPackage(w http.ResponseWriter, r *http.Request) {
	// Decode JSON body
	//authData := r.Context().Value("admin").(*auth.AdminClaims)
	authData := AuthAdmin(r)
	if ok := helper.ValidateAndDecodeJSON(w, r, &packageData, validateData); !ok {
		return // helper already wrote response
	}
	Name := strings.ToLower(helper.CleanString(packageData.Name))

	err := db.DB.Select(&packages, "SELECT id, name, uid FROM packages WHERE subscriber=? and name=? limit 1",
		authData.Admin.Subscriber, Name)

	if err != nil {
		// This means actual DB error (connection, query issue, etc.)
		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}

	// At this point err == nil
	if len(packages) != 0 {
		// pack found
		json.NewEncoder(w).Encode(response(false, "Package exist please add new one"))
		return
	}
	UidToday := helper.GenerateInfos(strings.ToLower(packageData.Name))

	res, err := db.DB.Exec(`insert into packages
	(name, uid, packValid, packCover, 
	status, packEligible, packType, packDetail,
    price, uidCreator, subscriber, commentData, 
	created_at) values(?,?,?,?,?,?,?,?,?,?,?,?,?)`, Name, UidToday.ID, packageData.PackValid, packageData.PackCover,
		packageData.Status, packageData.PackEligible, packageData.PackType, packageData.PackDetail,
		packageData.Price, authData.Admin.UID, authData.Admin.Subscriber, packageData.CommentData,
		UidToday.CreatedAt)
	if err != nil {
		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}
	rows, _ := res.RowsAffected()
	id, _ := res.LastInsertId()
	if rows == 0 {
		json.NewEncoder(w).Encode(response(false, "Something Wrong Please try again"))
		return

	}

	json.NewEncoder(w).Encode(response(true, id))
	//return

}

func EditPackage(w http.ResponseWriter, r *http.Request) {

}
func DelPackage(w http.ResponseWriter, r *http.Request) {

}
