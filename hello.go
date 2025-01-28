package main

import (
	"fmt"
	"runtime"
)

var c, python, java bool = true, true, false
var i, j = 1, 2

const Pi = 3.14

func calculator(a, b int) string {
	myString := "hello" + "world"
	fmt.Println(myString)
	return fmt.Sprintf("Multiply: %d", a*b)
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(calculator(7, 7))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(i, j, c, python, java)

	k := 3
	fmt.Println(k)

	var i int
	var f float64
	var b2 bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b2, s)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	test := 30
	if test < 11 {
		fmt.Printf("%v \n", test)
	} else {
		fmt.Println("hellow")
	}
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}

	fmt.Println("couning")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
