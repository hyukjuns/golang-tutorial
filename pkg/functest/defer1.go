package functest

import "fmt"

func Defer1() {
	// Defer 함수: 실행을 지연시키는 기능
	// Defer로 호출한 함수는 메인 쓰레드가 종료되기 직전에 호출된다
	// 타 언어의 finally 문과 비슷
	// 주로 리소스 반환에 사용 -> 열린 파일 닫기, Mutex 잠금 해제
	fmt.Println("Start...")
	defer Defer2()
	fmt.Println("end!")
}

func Defer2() {
	fmt.Println("Defer2 Called!")
}

func Defer3(msg string) {
	// 함수가 종료되기 전에 실행
	defer func() {
		fmt.Println(msg)
	}()

	// 먼저실행
	func() {
		fmt.Println("Hello")
	}()
}

func Stack() {
	// 10이 가장 마지막에 Defer로 처리되었기 때문에, 10부터 출력됨
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
}

// 중첩함수느 먼저 실행됨
func DupleFunc() {
	defer end(start("hey"))
	fmt.Println("You are in DupleFunc")
}
func start(s string) string {
	fmt.Println(s, "I'm Start Function")
	return s

}
func end(s string) {
	fmt.Println(s, "I'm End Function")
}
