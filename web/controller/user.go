package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"image/png"
	"net/http"
	p "web/proto"
	"web/utils"
)

// 获取 session信息
func GetSession(ctx *gin.Context)  {
	//  初始化错误返回的map
	resp := make(map[string]string)

	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	ctx.JSON(http.StatusOK, resp)
}
// 获取 图片信息
func GetImageCd(ctx *gin.Context)  {
	uuid := ctx.Param("uuid")

	// 指定服务发现
	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
		)
	// 初始化客户端
	microClient := p.NewGetCaptchaService("service.getCaptcha",consulService.Client())
	response, err := microClient.Call(context.TODO(), &p.Request{Uuid: uuid})
	if err != nil {
		fmt.Println("未找到远程服务")
		return
	}
	var img captcha.Image
	json.Unmarshal(response.Img, &img)
	png.Encode(ctx.Writer,img)

	//fmt.Println("str:",str)
	fmt.Println("uuid",uuid)


}
