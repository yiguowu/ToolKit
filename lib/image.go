package lib

import (
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func GetImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	return img, err
}

func GetImageFromString(data string) (image.Image, error) {
	i := strings.Index(data, ",")
	if strings.Contains(data, "image/png") {
		dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[i+1:]))
		img, err := png.Decode(dec)
		return img, err
	} else if strings.Contains(data, "image/jpeg") {
		dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[i+1:]))
		img, err := jpeg.Decode(dec)
		return img, err
	} else {
		return nil, errors.New("UnsupportedFormat")
	}
}
