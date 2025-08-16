package functest

import "fmt"

func AnonymousFunc() {
	// 익명함수, 즉시실행에 사용 (선언과 동시 사용)
	func() {
		fmt.Println("Hello Anonymous Function!")
	}()

	msg := "Hello World!"
	func(m string) {
		fmt.Println(m)
	}(msg)

	func(x, y int) {
		fmt.Println(x + y)
	}(10, 10)

	f := func(x, y int) int {
		return x * y
	}(10, 10)
	fmt.Println(f)
}
