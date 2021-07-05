package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"getCaptcha/conf"
	"getCaptcha/model"
	getCaptcha "getCaptcha/proto"
	"github.com/afocus/captcha"
	"image/color"
)

type GetCaptcha struct {

}


func (e *GetCaptcha) Call(ctx context.Context, req *getCaptcha.Request, rsp *getCaptcha.Response) error{
	cap := captcha.New()

	// 设置字体
	err := cap.SetFont(conf.GetCurrentAbS()+"/conf/comic.ttf")
	if err != nil {
		fmt.Println(err)
	}

	// 设置验证码大小
	cap.SetSize(128,64)
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{0,255,255,255})
	cap.SetBkgColor(color.RGBA{255,255,0,0})
	img,str := cap.Create(4,captcha.NUM)

	// 存储图片验证码到 redis中
	err = model.SaveImgCode(str, req.Uuid)
	if err != nil {
		return err
	}
	fmt.Println("save uuid suss:",req.Uuid)
	// 生成的图片序列化
	imgBuf , _ := json.Marshal(img)
	rsp.Img = imgBuf
	fmt.Println("已发送数据")
	return nil

}
