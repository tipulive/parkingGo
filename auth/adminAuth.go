package auth

import (
	"encoding/json"
	"net/http"

	"parking/dbQuery"
	"parking/helper"

	"fmt"
	"parking/model"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AdminCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminClaims struct {
	Admin model.Admin `json:"admin"`
	Token string      `json:"token,omitempty"`
	jwt.RegisteredClaims
}
type AdminUserData struct {
	Username string `json:"username"`
	Tel      string `json:"Tel"`
}

var claims *AdminClaims

// Fake user check
var response = helper.GetResponse

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)
	admin := model.Admin{Email: creds.Username, Name: creds.Password}
	success, result := dbQuery.AdminLogin(admin)

	if !success {
		//http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		message := fmt.Sprintf("Invalid credentials", http.StatusUnauthorized)

		json.NewEncoder(w).Encode(response(false, message))
		return
	}
	//500 means 500 minutes
	expirationTime := time.Now().Add(500 * time.Minute) //minute to be expired
	claims = &AdminClaims{

		Admin: result[0],

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	claims.Admin.Password = "no Password"
	/*token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(AdminJwtKey)**/
	tokenStr, err := TokenSigned(AdminJwtKey, claims)

	claims.Token = tokenStr
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}
	Tokens := map[string]string{
		"Token": tokenStr,
	}
	json.NewEncoder(w).Encode(response(true, Tokens))
	//fmt.Println("Generated Token:", tokenStr)
	//json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
}

// AdminAuthMiddleware for Chi
func AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing auth header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		//claims = &AdminClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return AdminJwtKey, nil
		})

		if err != nil || !token.Valid {
			message := "Invalid credentials"
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response(false, message))
			return
		}

		// Save claims in context

		ctx := ClaimSet(r.Context(), AdminClaimsKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// middreware
/*func AdminAuthMiddleware(next http.HandlerFunc) http.HandlerFunc { //User Authentication
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
			//http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			message := fmt.Sprintf("Invalid credentials", http.StatusUnauthorized)

			json.NewEncoder(w).Encode(response(false, message))
			return
		}
		// Save claims in context
		ctx := context.WithValue(r.Context(), "admin", adminCla)
		next.ServeHTTP(w, r.WithContext(ctx))
		//next.ServeHTTP(w, r)
	}
}*/
