package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-2369")

	if buffer.Len() == 0 {
		t.Errorf("No QRCode generated")
	}

	_, err := png.Decode(buffer)
	if err != nil {
		t.Errorf("Generated QR Code is not a png: %s", err)
	}
}
