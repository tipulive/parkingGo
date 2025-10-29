package controller

import (
	"encoding/json"
	"net/http"
	"parking/auth"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

	authData := r.Context().Value(auth.UserClaimsKey).(*auth.Claims)

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

	json.NewEncoder(w).Encode(authData)
	//json.NewEncoder(w).Encode(auth.User)

	//w.Write([]byte("This is %s a protected endpoint. You're authenticated!",auth.Name))
}
