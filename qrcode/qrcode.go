package qrcode

import (
	"github.com/skip2/go-qrcode"
)

// GenerateQRCode :
// Gera um QRCode a partir de um texto.
func GenerateQRCode(text string) []byte {
	var qrcodeSize = 512
	var borderSize = 10

	var qrcode, _ = qrcode.New(text, qrcode.Low)
	qrcode.DisableBorder = true
	var qrcodeBit, _ = qrcode.PNG(qrcodeSize - borderSize)

	return qrcodeBit
}
