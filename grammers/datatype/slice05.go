package datatype

import "fmt"

func Datatype18() {
	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// make로 공간을 할당 후 복사 해야 한다.
	// 복사된 슬라이스 값 변경해도 원본에는 영향 없음

	// 복사
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5, 5)
	slice3 := []int{}

	copy(slice2, slice1) // slice2의 용량이 5이기때문에, 5까지만 복사함
	copy(slice3, slice1) //  복사 불가
	fmt.Println(slice2)
	fmt.Println(slice3)

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5, 5)
	copy(b, a) // 값 복사
	b[0] = 7
	b[4] = 77

	fmt.Println(a, b)

	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // 참조 복사 (배열을 부분 추출하면 슬라이스의 참조 복사 됨)

	d[1] = 7
	fmt.Println(c, d)

	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // 슬라이스 시작:끝:용량
	fmt.Println(len(f), cap(f), f)

}
