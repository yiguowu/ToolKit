package lib

import (
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"net/url"
)

func GetSecretFromImageString(data string) string {
	img, _ := GetImageFromString(data)
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)
	u, _ := url.Parse(result.String())
	m, _ := url.ParseQuery(u.RawQuery)
	return m.Get("secret")
}
