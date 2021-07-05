package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 储存id 到Redis数据库
func SaveImgCode(code,uuid string) error {

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect Redis Err",err)
		return err
	}
	defer conn.Close()

	_, err = conn.Do("setex", uuid, 60, code)
	return err
}
