package main

import "fmt"

type queue []int

func (q *queue) Push(val int) {
	*q = append(*q, val)
}
func (q *queue) Pop() (int, bool) {
	if len(*q) == 0 {
		return -1, false
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, true
}

func main() {
	myqueue := queue{}

	for i := 0; i <= 10; i++ {
		myqueue.Push(i)
	}
	fmt.Println(myqueue)

	for {
		val, ok := myqueue.Pop()
		if !ok {
			fmt.Println("Queue is Empty!")
			break
		}
		fmt.Println(val)
	}
}
