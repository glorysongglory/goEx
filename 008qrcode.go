package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"
)

func main() {
	//文件输出
	qrcode.WriteFile("http://www.baidu.org/", qrcode.Medium, 256, "./qrcode.png")
	//byte数组输出
	byteArr, _ := qrcode.Encode("http://www.baidu.org/", qrcode.Medium, 256)
	fmt.Println(string(byteArr))
	// 自定义输出
	qr, err := qrcode.New("http://www.baidu.org", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{50, 205, 50, 255}
		qr.ForegroundColor = color.White
		qr.WriteFile(256, "./cusqrcode.png")
	}

}
