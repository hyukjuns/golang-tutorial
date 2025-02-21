package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("ex1: ", i)
	}

	// error case 1
	/*
		for i := 0; i<5; i++
		{
			fmt.Println("ex1: ", i)
		}
	*/

	// error case 2
	/*
		for i := 0; i < 5; i ++
			fmt.Println("ex1: ", i)
	*/

	// infinite
	/*
		for {
			fmt.Println("hello,")
			fmt.Println("world!")
		}
	*/

	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println(index, name)
	}
	for _, name := range loc {
		fmt.Println(name)
	}

}
