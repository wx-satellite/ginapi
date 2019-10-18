package handler

import "fmt"

func SayHello(args ...interface{}) {
	fmt.Println(args)
}
