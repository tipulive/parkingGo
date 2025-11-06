package controller

import (
	"fmt"
	"log"
	"net/http"
	"parking/db"
	"parking/dbQuery"
)

func Inc(w http.ResponseWriter, r *http.Request) {

	newVersion, err := db.RDS.Incr(db.CTXRedis, db.RoleVersionKey).Uint64()
	if err != nil {
		log.Fatalf("failed to increment roles_version: %v", err)
	}
	db.RDS.Publish(db.CTXRedis, db.RoleVersionPubSub, fmt.Sprintf("%d", newVersion))
	fmt.Println("Admin updated roles → roles_version incremented to", newVersion)

}
func Create(w http.ResponseWriter, r *http.Request) {

	dbQuery.CreateRedis()

}
func Read(w http.ResponseWriter, r *http.Request) {

	dbQuery.ReadRedis()

}
func Delete(w http.ResponseWriter, r *http.Request) {

	dbQuery.DeleteRedis()

}
func Update(w http.ResponseWriter, r *http.Request) {

	dbQuery.UpdateRedis()

}
