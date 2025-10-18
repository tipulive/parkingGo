package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"parking/auth"
	"parking/controller"

	//"parking/controller"

	"parking/dbQuery"

	"parking/helper"
	"parking/model"

	//"github.com/gorilla/mux"
	"github.com/go-chi/chi/v5"
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
type ErrResponse struct {
	Status bool        `json:"status"`
	Result interface{} `json:"result"`
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
func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // e.g., 400, 401, 405, 500
	json.NewEncoder(w).Encode(ErrResponse{
		Status: false,
		Result: message,
	})
}
func ApiRoute() {
	r := chi.NewRouter()

	//Admin
	//r.HandleFunc(fmt.Sprintf("/api/v%s/AdminLogin", os.Getenv("API_VERS")), auth.AdminLoginHandler).Methods("POST")
	//r.HandleFunc(fmt.Sprintf("/api/v%s/AdminLogin", os.Getenv("API_VERS")), auth.AdminLoginHandler).Methods("POST")

	port := ":" + os.Getenv("APP_PORT")
	apiVer := os.Getenv("API_VERS")
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		writeJSONError(w, http.StatusNotFound, "Route not found")
	})

	r.Route(fmt.Sprintf("/api/%s", apiVer), func(r chi.Router) {
		// Public routes

		r.Post("/AdminLogin", auth.AdminLoginHandler)
		r.Post("/login", auth.LoginHandler)

		// Protected routes (using Admin middleware)
		r.Group(func(r chi.Router) {
			r.Use(auth.AdminAuthMiddleware)

			r.Post("/AdminProtected", controller.AdminProtected)

			r.Route("/package", func(r chi.Router) {
				r.Get("/", controller.Package)           // POST /parking
				r.Post("/add", controller.AddPackage)    // POST /parking/add
				r.Post("/edit", controller.EditPackage)  // POST /parking/update
				r.Post("/delete", controller.DelPackage) // POST /parking/delete
			})
			r.Route("/payment", func(r chi.Router) {
				r.Get("/", controller.Payment)                // POST /parking
				r.Post("/cash", controller.AddPackage)        // POST /parking/add
				r.Post("/membership", controller.EditPackage) // POST /parking/update
				r.Post("/delete", controller.DelPackage)      // POST /parking/delete
			})

			r.Route("/parking", func(r chi.Router) {
				r.Get("/", controller.Parking) // POST /parking
				//r.Post("/add", controller.AddPackage)    // POST /parking/add
				//r.Post("/edit", controller.EditPackage)  // POST /parking/update
				//r.Post("/delete", controller.DelPackage) // POST /parking/delete
			})

		})

		// Protected routes (using Users middleware)
		r.Group(func(r chi.Router) {
			r.Use(auth.AuthMiddleware)

			r.Post("/protectData", auth.ProtectedHandler)
			//r.Post("/addpackage", controller.AddPackage)

		})
	})
	http.ListenAndServe(port, r)

}
