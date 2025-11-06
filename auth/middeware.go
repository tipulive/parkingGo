package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"parking/helper"
)

// PermissionMiddleware checks if the authenticated admin has permission for a specific key.
func PermissionMiddleware(permissionKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			admin := r.Context().Value(AdminClaimsKey).(*AdminClaims)
			if admin == nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(response(false, "Unauthorized"))
				return
			}

			perm := helper.ConvertStringToJson(admin.Admin.PermissionJSON)
			if perm[permissionKey] != "true" { //nukugango nitazaba ingana na true
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(response(false, "You do not have permission to access this option. Please contact the system admin."))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func AppVersionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("App-Version", "5.0.0") // Example version
		appVersion := r.Header.Get("App-Version")
		appVer := os.Getenv("APP_VERS")
		if appVersion != appVer {
			w.WriteHeader(http.StatusBadRequest)
			//fmt.Fprint(w, "Missing App-Version header")
			//json.NewEncoder(w).Encode(response(false, appVersion))
			//downloadLink := os.Getenv("APP_DOWNLOAD_LINK")
			downloadLink := os.Getenv("APP_DOWNLOAD_LINK")
			responseMsg := map[string]string{
				"message":            "Missing App-Version header or Please Update Your App by Clicking Here",
				"newUrlink_download": downloadLink,
			}
			//json.NewEncoder(w).Encode(response(false, "Please Update Your App by Clicking Here: "+downloadLink))
			json.NewEncoder(w).Encode(response(false, responseMsg))
			//

			//json.NewEncoder(w).Encode(response(false, "Missing App-Version header or Please Update Your App by Clicking Here"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
func AuthAdmin(r *http.Request) any {
	panic("unimplemented")
}
