package controller

import (
	"encoding/json"
	"net/http"
	"parking/helper"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	cfg := helper.UploadConfig{
		UploadDir:    "uploads",
		MaxFileSize:  5 * 1024 * 1024, // 5 MB
		AllowedTypes: []string{".jpg", ".png", ".pdf"},
		BaseURL:      "http://localhost:8080",
		Name:         "file",
	}

	result := helper.ImgUpload(w, r, cfg)
	if !result.Success {
		json.NewEncoder(w).Encode(result)
		return
	}
	json.NewEncoder(w).Encode(result)
}
func Qrcode(w http.ResponseWriter, r *http.Request) {
	// Automatically creates ./qrcodes/ if missing
	ok := helper.GenerateQrCode("https://example.com", "qrcodes/example.png", 256)
	if !ok {
		json.NewEncoder(w).Encode(response(false, "QR generation failed ❌"))
		return
	}

	json.NewEncoder(w).Encode(response(true, "QR generation successful ✅"))

}
