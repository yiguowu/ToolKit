package lib

import (
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/xlzd/gotp"
	"net/url"
)

func GetCurrentCode(secret string) string {
	totp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")
	return totp.Now()
}

func GetSecretFromImageString(data string) string {
	img, _ := GetImageFromString(data)
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)
	u, _ := url.Parse(result.String())
	m, _ := url.ParseQuery(u.RawQuery)
	return m.Get("secret")
}
