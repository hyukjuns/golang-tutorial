package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func loopExeCpu(name int) {
	fmt.Println(name, " Start", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>", i)
	}
	fmt.Println(name, " End", time.Now())
}

func GoroutineTest2() {

	// 현재 시스템의 CPU 코어 개수 반환 및 설정
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Current System CPU: ", runtime.GOMAXPROCS(0))

	fmt.Println("Main Routine Start", time.Now())
	for i := 0; i < 10; i++ {
		go loopExeCpu(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Main Routine End", time.Now())

	s := "Goroutine Closure"
	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i)
	}
	time.Sleep(1 * time.Second)

}
