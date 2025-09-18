package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"parking/auth"
	"parking/config"
	"parking/controller"
	"parking/db"
	"parking/dbQuery"
	"parking/helper"
	"parking/model"

	"github.com/gorilla/mux"
)

type Response struct {
	Message   string `json:"message"`
	InputData string `json:"InputData"`
	Status    int    `json:"status"`
}
type UserResponse struct {
	Status bool         `json:"status"`
	Result []model.User `json:"result"`
}

func GetEtags(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("ETag", etag)
	match := r.Header.Get("If-None-Match")
	w.Header().Set("ETag", match)
	w.Header().Set("Content-Type", "application/json")
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	myData, ok := data["inputAction"].(string)
	if !ok { //to avoid Errors
		http.Error(w, "inputAction must be a string", http.StatusBadRequest)
		return
	}
	message := "Hello, " + myData + helper.GetPanic()
	resp := Response{
		Message:   helper.GetHelp(myData),
		InputData: message,

		Status: 200,
	}

	json.NewEncoder(w).Encode(resp)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newUser := model.User{Name: data["name"].(string), Email: data["email"].(string)}
	if err := dbQuery.CreateUser(newUser); err != nil {
		fmt.Println("Create Error:", err)
	}
	resp := UserResponse{
		Status: true,
	}
	json.NewEncoder(w).Encode(resp)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, _ := dbQuery.GetUsers()
	//fmt.Println("Users:", users)
	resp := UserResponse{
		Status: true,
		Result: users,
	}
	json.NewEncoder(w).Encode(resp)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	myData, ok := data["inputAction"].(string)
	if !ok { //to avoid Errors
		http.Error(w, "inputAction must be a string", http.StatusBadRequest)
		return
	}
	message := "Hello, " + myData + helper.GetPanic()
	resp := Response{
		Message:   helper.GetHelp(myData),
		InputData: message,

		Status: 200,
	}
	json.NewEncoder(w).Encode(resp)
}
func CreateItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	myData := data["inputAction"]
	json.NewEncoder(w).Encode(myData)
}
func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)
	r := mux.NewRouter()

	//Admin
	r.HandleFunc("/api/AdminLogin", auth.AdminLoginHandler).Methods("POST")
	r.HandleFunc("/api/AdminProtected", auth.AdminAuthMiddleware(controller.AdminProtected)).Methods("POST")

	//Users

	r.HandleFunc("/api/login", auth.LoginHandler).Methods("POST") //this come from another file
	r.HandleFunc("/api/protected", auth.AuthMiddleware(auth.ProtectedHandler)).Methods("POST")
	r.HandleFunc("/api/protectData", auth.AuthMiddleware(auth.ProtectedRoute)).Methods("POST")
	r.HandleFunc("/api/GetTags", GetEtags).Methods("GET")
	r.HandleFunc("/api/items", getItems).Methods("GET")

	//r.HandleFunc("/api/items", get).Methods("GET")
	r.HandleFunc("/api/Create", CreateItems).Methods("POST")
	r.HandleFunc("/api/GetUsers", GetUsers).Methods("GET")
	r.HandleFunc("/api/CreateUser", CreateUser).Methods("POST")

	http.ListenAndServe(":8080", r)
}
