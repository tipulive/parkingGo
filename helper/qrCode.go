package helper

import (
	//"fmt"

	"os"
	"path/filepath"

	"github.com/skip2/go-qrcode"
)

func GenerateQrCode(content, filename string, size int) bool {
	// Ensure directory exists
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		//fmt.Println("❌ Failed to create folder:", err)
		return false
	}

	// Generate QR and save
	err := qrcode.WriteFile(content, qrcode.Medium, size, filename)
	if err != nil {
		//fmt.Println("❌ Failed to generate QR code:", err)
		return false
	}

	//fmt.Println("✅ QR code generated at:", filename)
	return true

	/*

		// Automatically creates ./qrcodes/ if missing
		ok := helper.GenerateQRCode("https://example.com", "qrcodes/example.png", 256)
		if !ok {
			fmt.Println("QR generation failed ❌")
			return
		}

		fmt.Println("QR generation successful ✅")*/

}
func MultGenerateQrCode() {
	panic("unimplemented")
}

func ScanQrCode() {
	panic("unimplemented")
}
