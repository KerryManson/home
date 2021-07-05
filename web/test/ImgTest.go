package main

import (
	"fmt"
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"net/http"
)

func main()  {
	cap := captcha.New()

	// 设置字体
	err := cap.SetFont("/Users/apple/Desktop/iHome/web/test/comic.ttf")
	if err != nil {
		fmt.Println(err)
	}

	// 设置验证码大小
	cap.SetSize(128,64)
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{0,255,255,255})
	cap.SetBkgColor(color.RGBA{255,255,0,0})

	http.HandleFunc("/r", func(writer http.ResponseWriter, request *http.Request) {
		img,str := cap.Create(6,captcha.ALL)
		png.Encode(writer,img)
		print(str)

	})

	fmt.Println("启动服务")
	err = http.ListenAndServe(":8061",nil)
	fmt.Println(err)
}
