package main

import "fmt"

func main() {
	i := 100

	if i >= 120 {
		fmt.Println("over 120")
	} else if i >= 100 && i < 120 {
		fmt.Println("over 100 and under 120")
	} else if i < 100 && i > 50 {
		fmt.Println("under 100 and over 50")
	} else {
		fmt.Println("under 50")
	}
}
