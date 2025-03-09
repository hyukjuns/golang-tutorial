package goroutine

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe1 Start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 End", time.Now())
}
func exe2() {
	fmt.Println("exe2 Start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 End", time.Now())
}
func exe3() {
	fmt.Println("exe3 Start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 End", time.Now())
}

func loopExe(name string) {
	fmt.Println(name, " Start", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>", i)
	}
	fmt.Println(name, " End", time.Now())
}

func GoroutineTest1() {

	// 가장 먼저 실행
	exe1()

	// 메인 함수 종료시 고루틴도 함께 종료됨
	fmt.Println("Main Routine Start", time.Now())
	go exe2()
	go exe3()
	time.Sleep(1 * time.Second)
	// fmt.Scanln()
	fmt.Println("Main Routine End", time.Now())

	loopExe("Test1")
	fmt.Println("Main Routine Start", time.Now())
	go loopExe("Test2")
	go loopExe("Test3")
	time.Sleep(1 * time.Second)
	// fmt.Scanln()
	fmt.Println("Main Routine End", time.Now())

}
