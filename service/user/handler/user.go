package handler

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"user/model"
	"user/proto/user"
	"user/utils"
)

type User struct {

}


func (h *User) SendSms(ctx context.Context, in *user.Request, out *user.Response) error{
	resp := make(map[string]string)
	// 获取短信验证码

	phone := in.Phone
	imgCode := in.ImgCode
	uuid := in.Uuid

	result := model.CheckImgCode(uuid, imgCode)
	if result {
		//发送短信
		tr := &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		//随机生成验证码 播种随机数种子

		rand.Seed(time.Now().Unix())
		smsCode :=fmt.Sprintf("%06d",rand.Int31n(1000000))


		response, err := client.Get(fmt.Sprintf("http://api.sms.cn/sms/?ac=send&uid=j1102945510&pwd=4acc173568c5917a84f9c9da0ddbb02c&template=100006&mobile=%s&content={\"code\":%s}",phone,smsCode))

		if err != nil {
			fmt.Println("验证码发送:err:",err)
			return err
		}
		// body,err := ioutil.ReadAll(resp.Body)
		formDta := make(map[string]interface{})
		json.NewDecoder(response.Body).Decode(&formDta)
		for key,value := range formDta{
			fmt.Println("key:",key," => value :",value)
		}

		if formDta["stat"] == "100"{

			resp["error"] = utils.RECODE_OK
			resp["msg"] = utils.RecodeText(utils.RECODE_OK)
			// 将短信验证码存入redis
			err := model.SaveSmsCode(phone, smsCode)
			fmt.Println("smsCode is:",smsCode)
			if err != nil {
				fmt.Println("存储短信验证码失败:",err)
				resp["error"] =utils.RECODE_DBERR
				resp["msg"] = utils.RecodeText(utils.RECODE_DBERR)
			}

		}else {
			resp["error"] =utils.RECODE_SMSERR
			resp["msg"] = utils.RecodeText(utils.RECODE_SMSERR)
		}

	}else {
		// 图片验证码校验失败
		resp["error"] =utils.RECODE_DATAERR
		resp["msg"] = utils.RecodeText(utils.RECODE_DATAERR)
		//r
	}
	out.Errno = resp["error"]
	out.Errmsg = resp["msg"]
	return nil
}

func (h *User) Register(ctx context.Context, in *user.RegReq ,out *user.Response) error {
	err := model.RegisUser(in.Mobile, in.Passwd, in.SmsCode)
	if err!= nil{
		fmt.Println("注册失败:",err)
		out.Errno = utils.RECODE_DBERR
		out.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	out.Errno = utils.RECODE_OK
	out.Errmsg = utils.RecodeText(utils.RECODE_OK)
	return nil
}