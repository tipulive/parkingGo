package controller

import (
	"net/http"
	"parking/dbQuery"
)

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
