package controller

import (
	"net/http"
	"parking/auth"
	"parking/helper"
	"parking/model"

	"github.com/go-playground/validator/v10"
)

var packageData model.Package
var packages []model.Package
var paymentData model.Subscription
var Subscription []model.Subscription
var adminData model.Admin
var userData model.User

var response = helper.GetResponse

var validateData = validator.New()

type RequestContext struct {
	W    http.ResponseWriter
	R    *http.Request
	Data map[string]any
}

func AuthAdmin(r *http.Request) *auth.AdminClaims {
	r.Context().Value("admin").(*auth.AdminClaims).Admin.Password = "no password"

	return r.Context().Value("admin").(*auth.AdminClaims)
}
func AuthUser(r *http.Request) *auth.Claims {
	r.Context().Value("user").(*auth.Claims).User.Password = "no password"
	return r.Context().Value("user").(*auth.Claims)
}
