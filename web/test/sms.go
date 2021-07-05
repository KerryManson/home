package main


import (
"github.com/cloopen/go-sms-sdk/cloopen"
"log"
)
func main() {
	cfg := cloopen.DefaultConfig().
		// 开发者主账号,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIAccount("8a216da87a332d53017a75c0f0fc1c29").
		// 主账号令牌 TOKEN,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIToken("auth a0c9ef4cbe12418882f78a67df040a7e")
	sms := cloopen.NewJsonClient(cfg).SMS()
	// 下发包体参数
	input := &cloopen.SendRequest{
		// 应用的APPID
		AppId: "appId",
		// 手机号码
		To: "17666663433",
		// 模版ID
		TemplateId: "templateId",
		// 模版变量内容 非必填
		Datas: []string{"您的验证码是****"},
	}
	// 下发
	resp, err := sms.Send(input)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Response MsgId: %s \n", resp.TemplateSMS.SmsMessageSid)

}