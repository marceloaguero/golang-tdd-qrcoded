package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello QR Code")

	qrcode := GenerateQRCode("555-2358")
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

// GenerateQRCode returns the QR code corresponding to the string received
func GenerateQRCode(code string) []byte {
	return nil
}
