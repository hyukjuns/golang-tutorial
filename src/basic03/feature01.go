package main

import "fmt"

func main() {
	// 후치 연산자 허용 i++, 전치 연산자 불허 ++i
	// 증감연산자 반환값 없음
	// 포인터(Pointer 허용, 연산 비허용)
	// 주석 //, /**/

	sum, i := 0, 0

	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)
}
