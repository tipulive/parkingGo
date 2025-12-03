package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"parking/db"
	"parking/helper"
	"strings"
)

func ViewSales(w http.ResponseWriter, r *http.Request) {
	//authData := AuthAdmin(r)
	if ok := helper.ValidateAndDecodeJSON(w, r, &paymentData, validateData); !ok {
		return // helper already wrote response
	}
	mapQuery := map[string]any{
		"today": map[string]any{
			"query": "WHERE id=? AND name=?",
			"param": []any{paymentData.ID, paymentData.PackName}, // slice of parameters
		},
		"thisWeek": map[string]any{
			"query": paymentData.PackName,
			"param": []any{2, "week basic"}, // should also be slice
		},
		"thisMonth": map[string]any{
			"query": paymentData.PackName,
			"param": []any{3, "month basic"}, // slice
		},
		"packages": packages,
	}

	// Correct SQL query
	Query := fmt.Sprintf("SELECT * FROM packages %v limit 1", mapQuery[paymentData.InputAction].(map[string]any)["query"].(string))

	// Extract parameters
	//params := mapQuery["today"].(map[string]any)["param"].([]any)
	params := mapQuery[paymentData.InputAction].(map[string]any)["param"].([]any)
	// Execute query
	err := db.DB.Select(&packages, Query, params...)
	if err != nil {
		// This means actual DB error (connection, query issue, etc.)
		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}

	if len(packages) == 0 {
		// pack found
		json.NewEncoder(w).Encode(response(false, "hello"))
		return
	}

	json.NewEncoder(w).Encode(response(true, packages))

}

func CheckPayment(w http.ResponseWriter, r *http.Request) { //this will show me how much user will pay for any hours he will park
	/*-- Most practical approach for 100 hours
	  SET @user_hours = 24;

	  -- Find the package that gives the best value for long durations
	  SELECT
	      pp.id as package_id,
	      pp.package_name,
	      pp.hours_included,
	      pp.price,
	      -- Calculate cost per hour for this package
	      ROUND(pp.price / pp.hours_included, 2) as cost_per_hour,
	      -- Calculate how many of these packages needed
	      CEIL(@user_hours / pp.hours_included) as packages_needed,
	      -- Calculate total cost
	      CEIL(@user_hours / pp.hours_included) * pp.price as total_cost,
	      -- Calculate effective hourly rate
	      ROUND((CEIL(@user_hours / pp.hours_included) * pp.price) / @user_hours, 2) as effective_hourly_rate,
	      -- Calculate total hours purchased (might be more than needed)
	      CEIL(@user_hours / pp.hours_included) * pp.hours_included as total_hours_purchased
	  FROM parking_packages pp
	  WHERE pp.is_active = TRUE
	  ORDER BY effective_hourly_rate ASC, total_cost ASC
	  LIMIT 1;*/
	id := 100
	query := `
    SELECT
        pp.id as uid,
        pp.name,
        pp.packValid,
        pp.price,
        -- Calculate cost per hour for this package
        ROUND(pp.price / pp.packValid, 2) as cost_per_hour,
        -- Calculate how many of these packages needed
        CEIL(? / pp.packValid) as packages_needed,
        -- Calculate total cost
        CEIL(? / pp.packValid) * pp.price as total_cost,
        -- Calculate effective hourly rate
        ROUND((CEIL(? / pp.packValid) * pp.price) / ?, 2) as effective_hourly_rate,
        -- Calculate total hours purchased (might be more than needed)
        CEIL(? / pp.packValid) * pp.packValid as total_hours_purchased
    FROM packages pp
    ORDER BY effective_hourly_rate ASC, total_cost ASC
    LIMIT 1;
`

	err := db.DB.Select(&packages, query, id, id, id, id, id)

	if err != nil {
		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}
	if len(packages) == 0 {
		json.NewEncoder(w).Encode(response(false, "Package not Available?"))
		return
	}
	json.NewEncoder(w).Encode(response(true, packages))

}
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
