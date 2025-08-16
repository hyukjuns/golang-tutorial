package basic

import "fmt"

func Feature01() {
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

func Feature02() {
	// 문장 끝 세미콜론 주의
	// 자동으로 끝에 세미콜론 삽입됨

	for i := 0; i <= 10; i++ {
		// fmt.Println("ex1: ");fmt.Println(i)
		fmt.Print("ex1: ")
		fmt.Println(i)
	}

	for j := 10; j >= 0; j-- {
		fmt.Println("ex2: ", j)
	}
}

func Feature03() {
	// 코드 서식 지정
	// 한 사람이 코딩 한 것 같은 일관성 유지
	// 코드 스타일 유지
	// gofmt -h : Help
	// gofmt -w : 원본파일에 반영

	for i := 0; i <= 100; i++ {
		fmt.Println("ex1 :", i)
	}
}
