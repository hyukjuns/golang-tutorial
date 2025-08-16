package goroutine

import (
	"fmt"
)

func Channel3() {
	ch := make(chan bool)
	cnt := 6

	// 수신과 송신은 동기식으로 진행됨
	// 쓰레드에서 대기 발생
	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go:", i)
			// time.Sleep(1 * time.Second)
		}
	}()
	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main:", i)
	}

}
