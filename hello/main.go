package main

import (
	"fmt"
	"testmod/pkg/hello"
	"testmod/pkg/world"
)

func main() {
	fmt.Print(hello.Hello())
	fmt.Print(" ")
	fmt.Println(world.World())
}
