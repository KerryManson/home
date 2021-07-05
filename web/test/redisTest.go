package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main()  {
	// 连接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect Redis Err",err)

	}
	defer conn.Close()

	reply, err := conn.Do("set", "hello2", "world")
	if err != nil {
		fmt.Println("set hello Err:",err)
	}
	s, err := redis.String(reply, err)

	fmt.Println("err:",err)
	fmt.Println(s)
	s2, err := redis.Strings(conn.Do("keys", "*"))
	fmt.Println(s2)
}
