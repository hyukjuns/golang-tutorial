package goroutine

import "fmt"

func Channel6() {
	// Close: 채널 닫기
	// Range: 채널이 닫히기 전까지 채널안에서 값을 꺼낸다
	// 채널을 닫지 않고 열려있으면 계속 대기함

	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	val1, ok1 := <-ch
	fmt.Println(val1, ok1)
	val2, ok2 := <-ch
	fmt.Println(val2, ok2)
	val3, ok3 := <-ch
	fmt.Println(val3, ok3)
	close(ch)
	val4, ok4 := <-ch
	fmt.Println(val4, ok4)

}
