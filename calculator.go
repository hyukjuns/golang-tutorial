package main

import "fmt"

func calculator(a int, b int) string {

	myString := "hello" + "world"
	fmt.Println(myString)
	return fmt.Sprintf("Multiply: %d", a*b)

}

func main() {
	fmt.Println(calculator(7, 7))
}
