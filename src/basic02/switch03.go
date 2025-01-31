package main

import "fmt"

func main() {
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		if a >= 20 {
			fmt.Println("over 20")
		}
		fmt.Println("a -> ", a, "는 홀수")
	}

	switch e := "go"; e {
	case "java":
		fmt.Println("Java!")
	case "go":
		fmt.Println("Go!")
		fallthrough
	case "python":
		fmt.Println("Python!")
	case "ruby":
		fmt.Println("Ruby!")
	case "c":
		fmt.Println("C!")
	}
}
