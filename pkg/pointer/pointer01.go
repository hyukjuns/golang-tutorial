package pointer

import "fmt"

func Pointer01() {
	i := 7

	a := &i
	b := &i
	fmt.Println(a, b)
	fmt.Println(*a, *b)
	fmt.Println(&a, &b)

	i = 8
	fmt.Println(a, b)
	fmt.Println(*a, *b)
	fmt.Println(&a, &b)

	var p1 *int            // nil 로 초기화
	var p2 *int = new(int) // 메모리 주소값 할당

	fmt.Println(p1, p2)

	j := 10
	pj := &j
	fmt.Println(j, *pj, &j, pj)

	*pj++
	fmt.Println(j, *pj, &j, pj)

	*pj = 20
	fmt.Println(j, *pj, &j, pj)

	var test_value int = 1000
	var test_value_2 int = 1000

	rptc(&test_value)
	vptc(test_value_2)

	fmt.Println(test_value, test_value_2)
}

// 함수 매개변수는 기본적으로 지역 변수로 값 복사
// 원본 변수값 변경을 위해서 포인터로 전달
// 크기가 큰 배열의 경우 값 복사시 리소스 부하 -> 포인터 전달을 통해 참조형식으로 리소스 부하 해결
func rptc(n *int) {
	*n = 10
}

func vptc(n int) {
	n = 10
}
