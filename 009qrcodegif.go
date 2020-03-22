package main

import (
	"github.com/skip2/go-qrcode"
	"image"
	"image/gif"
	"os"
)

func main() {
	channels := []string{"weixin", "toutiao", "qq", "baidu"}
	qrcodeSize := 256
	qrGif := gif.GIF{}
	for _, channel := range channels {
		q, err := qrcode.New("http://baidu.com?channel="+channel, qrcode.Medium)
		if err == nil {
			qrImage := q.Image(qrcodeSize)
			qrGif.Image = append(qrGif.Image, qrImage.(*image.Paletted))
			qrGif.Delay = append(qrGif.Delay, 10)
		}
	}
	qrFile, _ := os.OpenFile("qr_code.gif", os.O_WRONLY|os.O_CREATE, 0644)
	defer qrFile.Close()
	gif.EncodeAll(qrFile, &qrGif)
}
