package main

import "fmt"

func main() {

here:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break here
			}
			fmt.Println("ex1: ", i, j)
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex2: ", i)
	}

loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue loop2
			}
			fmt.Println("ex3: ", i, j)
		}
	}
}
