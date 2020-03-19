package main

import (
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	log.Println("Hello QR Code")

	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = GenerateQRCode(file, "555-2368", Version(1))
	if err != nil {
		log.Fatal(err)
	}
}

// GenerateQRCode returns a QR code representing the string code received
func GenerateQRCode(w io.Writer, code string, version Version) error {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	return png.Encode(w, img)
}

type Version int8
