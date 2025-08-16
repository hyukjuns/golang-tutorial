package goroutine

import (
	"fmt"
	"runtime"
)

func Channel4() {
	runtime.GOMAXPROCS(8)
	ch := make(chan bool, 4) // 버퍼사용
	cnt := 12

	// 버퍼사용
	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go:", i)
		}
	}()
	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main:", i)
	}

}
