package packagetest

import (
	"fmt"
)

// main 함수보다 먼저 실행
// 패키지 로드시에 가장 먼저 실행
// 가장 먼저 초기화 되는 작업 작성시 유용
func init() {
	fmt.Println("init start")
}

func Packagetest03() {
	fmt.Println("main start")
}
