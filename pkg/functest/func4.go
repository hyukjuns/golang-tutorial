package functest

import "fmt"

func Factorial(n int) int {
	// 재귀함수
	// 프로그램이 보기 쉽고, 코드 간결, 오류 수정 용이
	// 코드 이해 어렵, 메모리 많이 사용, 무한루프 가능성
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func PrtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("(", n, ")")
	PrtHello((n - 1))
}
