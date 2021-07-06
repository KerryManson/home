package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var RedisPool redis.Pool
var SMSCODESERVIVE int32 = 300

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
