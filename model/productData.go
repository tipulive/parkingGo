package model

type ProductData struct {
	ID          uint64 `json:"id" db:"id"`
	ProductCode string `json:"product_code" db:"productCode"`
	Price       string `json:"price" db:"price"`
	Qty         string `json:"qty" db:"qty"`
	ImgUrl      string `json:"img_url" db:"img_url"`
}
