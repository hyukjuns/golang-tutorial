package main

import "fmt"

func plus(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}
func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}
func main() {
	var a int
	var b int
	fmt.Print("Enter Two Number")
	fmt.Scan(&a, &b)
	fmt.Println("Result:")
	fmt.Printf("Plus: %d\n", plus(a, b))
	fmt.Printf("Minus: %d\n", minus(a, b))
	fmt.Printf("Multiply: %d\n", multiply(a, b))
	fmt.Printf("Divide: %d\n", divide(a, b))
}
