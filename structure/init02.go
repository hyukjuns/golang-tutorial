package main

import (
	"basic03/lib"
	"fmt"
)

func init() {
	fmt.Println("init 1")
}

func main() {
	fmt.Println("main 1")
	fmt.Println("test lib", lib.CheckNum(100))
}
