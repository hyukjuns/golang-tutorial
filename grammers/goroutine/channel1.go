package goroutine

import (
	"fmt"
	"time"
)

func work1(v chan int) {
	fmt.Println("Start work1 --->", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("End work1 --->", time.Now())
	v <- 1 // 데이터 송신
}

func work2(v chan int) {
	fmt.Println("Start work2 --->", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("End work2 --->", time.Now())
	v <- 2 // 데이터 송신
}

func Channel1() {

	fmt.Println("Start Main Routine", time.Now())
	// var c chan int
	// c = make(chan int)
	v := make(chan int)
	go work1(v)
	go work2(v)

	// 데이터 수신
	<-v
	<-v
	fmt.Println("End Main Routine", time.Now())

}
