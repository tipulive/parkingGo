package helper

import (
	"fmt"
	"io"

	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type UploadResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	URL     string `json:"url,omitempty"`
}

type UploadConfig struct {
	UploadDir    string
	MaxFileSize  int64
	AllowedTypes []string
	BaseURL      string
	Name         string
}

// HandleFileUpload uploads a single file, validates it, and returns a JSON response.
// You can call this from any route handler or controller.
func ImgUpload(w http.ResponseWriter, r *http.Request, cfg UploadConfig) UploadResponse {
	// Set max upload size
	r.Body = http.MaxBytesReader(w, r.Body, cfg.MaxFileSize)

	err := r.ParseMultipartForm(cfg.MaxFileSize)
	if err != nil {

		return UploadResponse{false, "File too large or invalid", ""}
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		return UploadResponse{false, "No file uploaded", ""}
	}
	defer file.Close()

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(handler.Filename))
	if !isAllowedType(ext, cfg.AllowedTypes) {
		return UploadResponse{false, "File type not allowed", ""}
	}

	// Ensure upload directory exists
	if err := os.MkdirAll(cfg.UploadDir, os.ModePerm); err != nil {
		return UploadResponse{false, "Failed to create upload directory", ""}
	}

	// Generate unique file name
	//newFileName := uuid.New().String() + ext
	newFileName := cfg.Name + ext
	savePath := filepath.Join(cfg.UploadDir, newFileName)

	// Save file
	out, err := os.Create(savePath)
	if err != nil {
		//json.NewEncoder(w).Encode(GetResponse(false,  "Error saving file"))
		return UploadResponse{false, "Error saving file", ""}
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return UploadResponse{false, "Failed to write file", ""}
	}

	// Return full URL (e.g. http://localhost:8080/uploads/uuid.png)
	fileURL := fmt.Sprintf("%s/%s/%s", cfg.BaseURL, cfg.UploadDir, newFileName)

	return UploadResponse{
		Success: true,
		Message: "File uploaded successfully",
		URL:     fileURL,
	}
}

func isAllowedType(ext string, allowed []string) bool {
	for _, a := range allowed {
		if strings.EqualFold(a, ext) {
			return true
		}
	}
	return false
}
