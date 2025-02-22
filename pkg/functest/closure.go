package functest

import "fmt"

func ClosureTest() {
	// Closure
	// 익명함수를 사용할 경우 함수를 변수에 할당해서 사용 가능
	// 함수 안에서, 함수를 선언하고 정의 가능
	// 외부 함수에서 선언된 변수에 접근 가능한 함수
	// 함수가 선언되는 순간에 함수가 실행(익명함수) 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
	// 함수가 호출 될 때, 이전에 존재했던 값을 유지하기 위해서 사용
	// 비동기, 누적카운트 등에 사용됨, 무분별한 전역변수 남발을 방지
	// 캡쳐된 객체들이 메모리에 상주하므로, 메모리 부족현상이 발생할 수 잇음
	multiply := func(x, y int) int {
		return x * y
	}
	r1 := multiply(10, 10)
	fmt.Println(r1)

	m, n := 10, 5            // 변수가 캡쳐됨
	sum := func(c int) int { // 익명함수를 변수에 할당
		return m + n + c // 캡쳐되어 들어온 지역변수가 소멸되지 않는다 (함수 호출 시 마다 사용 가능)
	}
	r2 := sum(10)

	fmt.Println(r2)
}

func ClosureTest2() {
	cnt := increaseCnt()
	cnt2 := increaseCnt()

	for i := 0; i < 10; i++ {
		fmt.Println(cnt())
		fmt.Println(cnt2())
	}

}

func increaseCnt() func() int {
	n := 0 // 캡쳐된 지역변수 (메모리에 상주함)
	return func() int {
		n += 1
		return n
	}
}
