package main

import "fmt"

func main() {
	// 반드시 boolean, 1,0 사용불가
	// 소괄호 사용 x

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("over 15")
	}

	if b >= 25 {
		fmt.Println("over 25")
	}

	// error case 1
	/*
		if b > 25
		{

		}
	*/

	// error case 2
	/*
		if b >= 25
			fmt.Println("over 25")
	*/

	// error case 3
	/*
		if c := 1; c {
			fmt.Println("over 25")
		}
	*/

	// error case 4
	/*
		if c := 40; c >= 35 {
			fmt.Println("over 35")
		}
		c = 10
	*/
}
