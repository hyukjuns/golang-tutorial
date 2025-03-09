package goroutine

import "fmt"

func rangeSum(rg int, c chan int) {
	sum := 0
	for i := 0; i <= rg; i++ {
		sum += i
	}
	c <- sum
}
func Channel2() {

	c := make(chan int)

	go rangeSum(10, c)
	go rangeSum(5, c)
	go rangeSum(2, c)

	// 순서대로 값 수신
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

}
