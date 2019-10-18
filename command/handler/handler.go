package handler

import (
	"fmt"
	"time"
)

func SayHello(args ...interface{}) {
	fmt.Println(args)

	time.Sleep(30 * time.Second)
}
