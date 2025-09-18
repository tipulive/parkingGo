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

var AdminJwtKey = []byte("Admin_secret_key")

type AdminCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminClaims struct {
	Admin model.Admin `json:"admin"`
	jwt.RegisteredClaims
}
type AdminUserData struct {
	Username string `json:"username"`
	Tel      string `json:"Tel"`
}

// Fake user check

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)
	admin := model.Admin{Email: creds.Username, Name: creds.Password}
	success, result := dbQuery.AdminLogin(admin)

	if !success {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute) //minute to be expired

	claims := &AdminClaims{

		Admin: result[0],
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(AdminJwtKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
}

// middreware
func AdminAuthMiddleware(next http.HandlerFunc) http.HandlerFunc { //User Authentication
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing auth header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		adminCla := &AdminClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, adminCla, func(token *jwt.Token) (interface{}, error) {
			return AdminJwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		// Save claims in context
		ctx := context.WithValue(r.Context(), "admin", adminCla)
		next.ServeHTTP(w, r.WithContext(ctx))
		//next.ServeHTTP(w, r)
	}
}
