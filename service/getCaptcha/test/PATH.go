package main

import (
	"fmt"
	"getCaptcha/conf"
)

func main() {
	s := conf.GetCurrentAbS()
	fmt.Println(s)
}
