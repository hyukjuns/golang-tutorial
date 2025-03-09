package goroutine

import "fmt"

func Channel5() {
	// Close: 채널 닫기
	// Range: 채널이 닫히기 전까지 채널안에서 값을 꺼낸다
	// 채널을 닫지 않고 열려있으면 계속 대기함

	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch) // 5회 채널에 값 전송 후 채널 닫기
	}()

	// 채널이 닫힐때까지 값을 수신
	for i := range ch {
		fmt.Println(i)
	}

}
