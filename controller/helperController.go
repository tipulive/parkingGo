package controller

import (
	"net/http"
	"parking/auth"
	"parking/helper"
	"parking/model"

	"github.com/go-playground/validator/v10"
)

var packageData model.Package         //form
var packages []model.Package          //table
var paymentData model.Subscription    //form
var Subscription []model.Subscription //table
var adminData model.Admin
var userData model.User
var productData model.ProductData
var products []model.ProductData

//var claims *auth.AdminClaims

var response = helper.GetResponse

var validateData = validator.New()

type RequestContext struct {
	W    http.ResponseWriter
	R    *http.Request
	Data map[string]any
}

func AuthAdmin(r *http.Request) *auth.AdminClaims {
	//r.Context().Value(auth.AdminClaimsKey).(*auth.AdminClaims).Admin.Password = "no password"

	return r.Context().Value(auth.AdminClaimsKey).(*auth.AdminClaims)
}
func AuthUser(r *http.Request) *auth.Claims {
	r.Context().Value(auth.UserClaimsKey).(*auth.Claims).User.Password = "no password"
	return r.Context().Value(auth.UserClaimsKey).(*auth.Claims)
}
