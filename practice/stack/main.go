package main

import "fmt"

type stack []int

func (s *stack) Push(value int) {
	*s = append(*s, value)
}
func (s *stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return -1, false
	}
	idx := len(*s) - 1
	val := (*s)[idx]
	*s = (*s)[:idx]
	return val, true
}

func main() {
	mystack := stack{}

	for i := 0; i <= 10; i++ {
		mystack.Push(i)
	}

	fmt.Println(mystack)

	for {
		val, ok := mystack.Pop()
		if !ok {
			fmt.Println("Stack is Empty!")
			break
		}
		fmt.Println(val)
	}
	fmt.Println(mystack)

}
