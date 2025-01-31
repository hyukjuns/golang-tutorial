package main

import "fmt"

func main() {
	// 문장 끝 세미콜론 주의
	// 자동으로 끝에 세미콜론 삽입됨

	for i := 0; i <= 10; i++ {
		// fmt.Println("ex1: ");fmt.Println(i)
		fmt.Print("ex1: ")
		fmt.Println(i)
	}

	for j := 10; j >= 0; j-- {
		fmt.Println("ex2: ", j)
	}
}
