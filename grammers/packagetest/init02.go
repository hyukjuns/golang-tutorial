package packagetest

import (
	"basic03/lib"
	"fmt"
)

func init() {
	fmt.Println("init 1")
}

func Packagetest04() {
	fmt.Println("main 1")
	fmt.Println("test lib", lib.CheckNum(100))
}
