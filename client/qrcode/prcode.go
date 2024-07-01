package qrcode

import (
	"log"

	"github.com/skip2/go-qrcode"
)

func CreateQrcode(id string) []byte {
	bytes, err := qrcode.Encode(id, qrcode.Medium, 256)
	if err != nil {
		log.Fatal("Failed to generate qrcode:", err)
	}
	return bytes
}
