package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"parking/dbQuery"
	"parking/model"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/*
	type Claims struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Tel      string `json:"tel"`
		jwt.RegisteredClaims
	}
*/
type Claims struct {
	User model.User `json:"user"`
	jwt.RegisteredClaims
}
type UserData struct {
	Username string `json:"username"`
	Tel      string `json:"Tel"`
}

// Fake user check
func AuthenticateUser(username, password string) bool {

	return username == "admin" && password == "1234"

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)
	user := model.User{Email: creds.Username, Name: creds.Password}
	success, result := dbQuery.UserLogin(user)

	/*if success {
		json.NewEncoder(w).Encode(result)
	} else {
		print(result)
	}*/
	if !success {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	//json.NewEncoder(w).Encode(result)

	expirationTime := time.Now().Add(5 * time.Minute) //minute to be expired

	claims := &Claims{
		/*User: model.User{
			ID:          result[0].ID,
			PhoneNumber: result[0].PhoneNumber,
			Name:        result[0].Name,
			Country:     result[0].Country,
			Email:       result[0].Email,
			CreatedAt:   result[0].CreatedAt,
		},*/
		User: result[0],
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	/*claims = &Claims{
		Username: "kevine",
		Name:     "murumba", // or creds.Name if that's what you meant

	}

	claims, ok := r.Context().Value("user").(*Claims)

	if ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}*/
	// Inject username into context
	//context.WithValue(r.Context(), userContextKey, claims.Username)

	// Pass context down

	json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
}

// middreware
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc { //User Authentication
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing auth header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		// Save claims in context
		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
		//next.ServeHTTP(w, r)
	}
}
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	auth := r.Context().Value("user").(*Claims)

	/*resp := Claims{
		User: model.User{
			ID:          auth.User.ID,
			PhoneNumber: auth.User.PhoneNumber,
			Name:        auth.User.Name,
			Country:     auth.User.Country,
			Email:       auth.User.Email,
			CreatedAt:   auth.User.CreatedAt,
		},
	}*/

	resp := UserData{
		Username: auth.User.Name,
		Tel:      auth.User.PhoneNumber,
	}
	json.NewEncoder(w).Encode(resp)
	//json.NewEncoder(w).Encode(auth.User)

	//w.Write([]byte("This is %s a protected endpoint. You're authenticated!",auth.Name))
}
func ProtectedRoute(w http.ResponseWriter, r *http.Request) {

	/*resp := Claims{
		Username: "hello",
		Name:     "bibi",
	}*/
	json.NewEncoder(w).Encode("hello")

	//w.Write([]byte("This is %s a protected endpoint. You're authenticated!",auth.Name))
}
