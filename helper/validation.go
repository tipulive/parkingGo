package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// ErrorResult represents a single result item in the response
type ErrorResult struct {
	Field     map[string]string `json:"Field,omitempty"`
	JSONError string            `json:"jsonError,omitempty"`
}

// Response is your standard API response
type Resp struct {
	Status bool          `json:"Status"`
	Result []ErrorResult `json:"result,omitempty"`
}

// GetResponse reuses your existing response helper
func GetResp(status bool, result interface{}) Resp {
	return Resp{
		Status: status,
		Result: result.([]ErrorResult),
	}
}

// ValidateAndDecodeJSON decodes JSON and validates all fields
// Returns false if any errors and sends structured response using GetResponse
func ValidateAndDecodeJSON(
	w http.ResponseWriter,
	r *http.Request,
	dst interface{},
	validate *validator.Validate,
) bool {

	result := ErrorResult{
		Field: make(map[string]string),
	}

	// Decode JSON
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		result.JSONError = fmt.Sprintf("%v", err)
	}

	// Validate struct fields
	if err := validate.Struct(dst); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				result.Field[e.Field()] = fmt.Sprintf("failed on '%s' rule", e.Tag())
			}
		} else {
			// fallback validation error
			result.Field["unknown"] = fmt.Sprintf("%v", err)
		}
	}

	// If any errors, respond
	if result.JSONError != "" || len(result.Field) > 0 {
		resp := GetResp(false, []ErrorResult{result})
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return false
	}

	return true
}
