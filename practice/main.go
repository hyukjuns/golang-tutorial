package main

import "fmt"

func Pyramid() {
	rows := 10
	padding := 10
	for i := 0; i < rows; i++ {
		stars := 2*i + 1
		for j := 0; j < padding; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < stars; j++ {
			fmt.Print("*")
		}
		for j := 0; j < padding; j++ {
			fmt.Print(" ")
		}
		padding -= 1
		fmt.Println()
	}
}

func main() {
	Pyramid()
}
