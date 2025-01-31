package main

import "fmt"

func main() {
	// short decalare variable
	// in function

	shortVar1 := 1
	shortVar2 := "test"
	shortVar3 := false

	shortVar1 = 10

	// shortVar3 := true // exception
	fmt.Println("shortVar1: ", shortVar1, "shortVar2: ", shortVar2, "shortVar3: ", shortVar3)

	if i := 10; i < 11 {
		fmt.Println("short variable test success")
	}
}
