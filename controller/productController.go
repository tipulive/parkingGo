package controller

import (
	"encoding/json"

	"net/http"
	"parking/db"
	"parking/helper"
	//"strings"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if ok := helper.ValidateAndDecodeJSON(w, r, &productData, validateData); !ok {
		return // helper already wrote response
	}

	res, err := db.DB.Exec(`insert into productDatas
	(productCode, price, qty) values(?,?,?)`, productData.ProductCode, productData.Price, productData.Qty)
	if err != nil {
		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}
	rows, _ := res.RowsAffected()
	id, _ := res.LastInsertId()
	if rows == 0 {
		json.NewEncoder(w).Encode(response(false, "Something Wrong Please try again"))
		return

	}

	json.NewEncoder(w).Encode(response(true, id))
	//return

}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if ok := helper.ValidateAndDecodeJSON(w, r, &productData, validateData); !ok {
		return // helper already wrote response
	}
	res, err := db.DB.Exec(`delete from productDatas where productCode=?`, productData.ProductCode)
	if err != nil {
		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}
	rows, _ := res.RowsAffected()

	if rows == 0 {
		json.NewEncoder(w).Encode(response(false, "Something Wrong Please try again"))
		return

	}

	json.NewEncoder(w).Encode(response(true, rows))
}

func LoadProduct(w http.ResponseWriter, r *http.Request) {
	if ok := helper.ValidateAndDecodeJSON(w, r, &productData, validateData); !ok {
		return // helper already wrote response
	}
	err := db.DB.Select(&products, "Select *from productDatas ")

	if err != nil {
		// This means actual DB error (connection, query issue, etc.)
		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}
	json.NewEncoder(w).Encode(response(true, products))
}
func SearchProduct(w http.ResponseWriter, r *http.Request) {
	if ok := helper.ValidateAndDecodeJSON(w, r, &productData, validateData); !ok {
		return // helper already wrote response
	}
	err := db.DB.Select(&products, "Select *from productDatas where productCode=?", productData.ProductCode)

	if err != nil {

		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}
	json.NewEncoder(w).Encode(response(true, products))
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if ok := helper.ValidateAndDecodeJSON(w, r, &productData, validateData); !ok {
		return // helper already wrote response
	}
	res, err := db.DB.Exec(`update productDatas set qty=?,price=? where productCode=?`, productData.Qty, productData.Price, productData.ProductCode)
	if err != nil {
		json.NewEncoder(w).Encode(response(false, err.Error()))
		return
	}
	rows, _ := res.RowsAffected()

	if rows == 0 {
		json.NewEncoder(w).Encode(response(false, "productCode not found or Something Wrong Please try again"))
		return

	}

	json.NewEncoder(w).Encode(response(true, "update Successful"))
}
