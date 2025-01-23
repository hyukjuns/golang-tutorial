package main

import "fmt"

func calculator(a int, b int) string {

	myString := "hello" + "world"
	fmt.Println(myString)
	return fmt.Sprintf("Minus: %d", a-b)

}

func main() {
	fmt.Println(calculator(1, 2))
}
