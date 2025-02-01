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

loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue loop2
			}
			fmt.Println("ex2: ", i, j)
		}
	}

here2:
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue here2
		}
		fmt.Println("ex3: ", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex3: ", i)
	}
}
