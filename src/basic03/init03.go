package main

import (
	"basic03/lib"
	"fmt"
)

var num int32

func init() {
	num = 30
	fmt.Println("main init start!")
}
func main() {
	fmt.Println("num is over 10?: ", lib.CheckNum(num))
}
