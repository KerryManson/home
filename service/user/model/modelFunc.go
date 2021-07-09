package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
)

var RedisPool redis.Pool
var SMSCODESERVIVE int32 = 30000000  // 验证码生存时间

func InitRedis()  {
	RedisPool = redis.Pool{MaxIdle: 20,
		MaxActive:       50,
		MaxConnLifetime: 60 * 5,
		IdleTimeout:     60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	fmt.Println("Init Redis suss")
}


func CheckImgCode(uuid, imgCode string) bool {
	//连接redis数据库
	conn := RedisPool.Get()
	defer conn.Close()

	redisCode, err := redis.String(conn.Do("get", uuid))
	if err != nil {
		fmt.Println("redis DO err:", err)
		return false
	}
	return redisCode == imgCode
}

func SaveSmsCode(phone, smsCode string) error {
	// redis 连接池
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("setex", phone+"_code", SMSCODESERVIVE, smsCode)
	return err

}

// 解密passwd
func PassWordVerify(s, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(s), []byte(hash))
	return err==nil
}

type SMSCodeErr struct {
}
type InsertDbErr struct {
}

func (s *SMSCodeErr) Error() string {
	return "验证码错误"
}
func (i InsertDbErr) Error()  string{
	return "插入数据库失败"
}

// 加密passwd
func PasswordHash(s string)(string,error){
	password, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	if err != nil {
		fmt.Println("加密失败:",err)
		return "", err
	}
	return string(password),err
}




// 校验短信验证码
func RegisUser(phone,password,smsCode string) error {
	RedisConn := RedisPool.Get()
	reply, err := redis.String(RedisConn.Do("get", phone+"_code"))
	if err != nil {
		fmt.Println("get phone sms Failed:",err)
		return err
	}
	h := md5.New()
	h.Write([]byte(password))
	pasawd_hash := hex.EncodeToString(h.Sum(nil))
	//hashPass,_ := PasswordHash(password)
	fmt.Println(reply)
	if reply == smsCode {
		user := User{Name: phone,
			Password_hash: pasawd_hash,
			Mobile: phone,
		}
		fmt.Println(user)
		err := GlobalDbConn.Create(&user).Error
		if err != nil {
			return &InsertDbErr{}
		}
		return nil

	}else {
		return &SMSCodeErr{}
	}

}