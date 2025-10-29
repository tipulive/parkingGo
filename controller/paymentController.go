package controller

import (
	"encoding/json"
	"net/http"
	"parking/db"
	"parking/helper"
	"strings"
)

func Payment(w http.ResponseWriter, r *http.Request) {

	//authData := AuthAdmin(r)
	if ok := helper.ValidateAndDecodeJSON(w, r, &paymentData, validateData); !ok {
		return // helper already wrote response
	}

	err := db.DB.Select(&packages, "SELECT * FROM packages WHERE uid=? and status='on' limit 1", paymentData.PackUID)

	if err != nil {
		// This means actual DB error (connection, query issue, etc.)
		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}

	UidToday := helper.GenerateInfos(strings.ToLower("Pid"))
	if len(packages) == 0 {
		// pack found
		json.NewEncoder(w).Encode(response(false, "Package not Available?"))
		return
	}
	//json.NewEncoder(w).Encode(response(true, packages))

	query := `
    INSERT INTO payment_histories (
        uid, uidUser, uidCreator, subscriber,
        packName, packUid, packPrice, packValid, packType,
        activated, userStatus, paidStatus, created_at, updated_at
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, DATE_ADD(?, INTERVAL ? DAY))
`

	res, err := db.DB.Exec(query,
		UidToday.ID,
		paymentData.UIDUser,
		AuthAdmin(r).Admin.UID,
		AuthAdmin(r).Admin.Subscriber,
		packages[0].Name,
		packages[0].UID,
		packages[0].Price,
		packages[0].PackValid,
		packages[0].PackType,
		"off",
		"member",
		"membership",
		UidToday.CreatedAt, // this is the `created_at` base for DATE_ADD
		UidToday.CreatedAt, // this is the `updated_at`
		packages[0].PackValid,
	)

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
	/*ctx := &RequestContext{
		W: w,
		R: r,

		Data: make(map[string]any),
	}
	//AuthAdmin(r).Admin.Password = "no password"
	ctx.Data["user"] = AuthAdmin(r)
	ctx.Data["payType"] = "membership"

	//delete(ctx.Data["user"]["password"], "password")
	ProcessPayment(ctx)*/
}
func ProcessPayment(ctx *RequestContext) {

	//name := ctx.Data["user"].(*auth.AdminClaims).Admin.Name
	json.NewEncoder(ctx.W).Encode(map[string]any{
		"message": ctx.Data["payType"],
		"result":  ctx.Data["user"],
		"user":    packageData.Name,
	})
}
func membership() {

}
func cash() {

}

func search() {

}
