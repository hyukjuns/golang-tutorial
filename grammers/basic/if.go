package basic

import "fmt"

func If01() {
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

func If02() {
	var a int = 50
	b := 70

	if a >= 65 {
		fmt.Println("over 65")
	} else {
		fmt.Println("under 65")
	}
	if b >= 70 {
		fmt.Println("over 70")
	} else {
		fmt.Println("under 70")
	}
}

func If03() {
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
