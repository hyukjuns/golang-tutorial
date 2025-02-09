package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// case 1
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, "over 50 and under 100")
	case i > 25 && i < 50:
		fmt.Println("i -> ", i, "over 25 and under 50")
	default:
		fmt.Println("i -> ", i, "default")
	}
}
