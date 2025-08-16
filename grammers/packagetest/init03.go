package packagetest

import (
	"basic03/lib"
	"fmt"
)

var num int32

func init() {
	num = 30
	fmt.Println("main init start!")
}
func Packagetest05() {
	fmt.Println("num is over 10?: ", lib.CheckNum(num))
}
