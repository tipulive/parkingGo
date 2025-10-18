package controller

import (
	//"database/sql"
	"encoding/json"
	"fmt"

	//"log"
	"parking/db"
	"parking/helper"

	//"parking/auth"
	"parking/model"

	"github.com/go-playground/validator/v10"

	"net/http"
)

var validate = validator.New()

func Parking(w http.ResponseWriter, r *http.Request) {
	//EditParking(w, r)
	// Define actions

}

func Test() {

}
func EditParking(w http.ResponseWriter, r *http.Request) { //editing in case there is errors

	/*var users []model.Admin
	  err := db.DB.Select(&users, "SELECT id, name, email FROM users where id")
	  if err != nil {
	      log.Fatal(err)
	  }
	  fmt.Printf("User: %+v\n", users[0].Name)*/

	/*var admin = model.Admin{
	      Name: "eric",
	  }

	  if err := validate.Struct(admin); err != nil {
	      if err != nil {
	          fmt.Printf("validation failed: %v", err)
	      }

	      fmt.Printf("")
	  }*/
	var admin model.Admin
	response := helper.GetResponse

	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		/*http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		  return*/
		resp := fmt.Sprintf("Invalid JSON: "+err.Error(), http.StatusBadRequest)

		json.NewEncoder(w).Encode(response(false, resp))
		return

	}

	// Validate input
	if err := validate.Struct(admin); err != nil {
		errs := err.(validator.ValidationErrors)
		w.WriteHeader(http.StatusBadRequest)
		for _, e := range errs {
			//fmt.Fprintf(w, "❌ Field '%s' failed on '%s' rule\n", e.Field(), e.Tag())
			resp := fmt.Sprintf("❌ Field '%s' failed on '%s' rule", e.Field(), e.Tag())

			json.NewEncoder(w).Encode(response(false, resp))
			return
		}

	}
	var users []model.Admin
	err := db.DB.Select(&users, "SELECT id, name, email FROM users")

	if err != nil {
		// This means actual DB error (connection, query issue, etc.)
		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}

	// At this point err == nil
	if len(users) == 0 {
		// No users found
		json.NewEncoder(w).Encode(response(false, "User not found"))
		return
	}
	json.NewEncoder(w).Encode(response(true, users))
	return

	/*var users model.Admin
	  err := db.DB.Get(&users, "SELECT id, name, email FROM users WHERE id=$1", 8)
	  if err != nil {
	      if err == sql.ErrNoRows {
	          log.Println("User not found")
	      } else {

	          response := helper.GetResponse(false, err)

	          json.NewEncoder(w).Encode(response)
	      }
	  }*/

	//test := helper.GetResponse(true, users)

	//json.NewEncoder(w).Encode(helper.GetResponse(false, users))
	//fmt.Println(time.Now().Local())

}
func DelParking() { //delete function

}
func ViewParking() {

}
