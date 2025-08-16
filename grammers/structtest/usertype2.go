package structtest

import "fmt"

type cnt int

func testConvertInt(i int) {
	fmt.Println(i)
}

func testConvertCnt(i cnt) {
	fmt.Println(i)
}

func UserType2() {
	a := cnt(5)
	fmt.Println(a)

	var b cnt = 15
	fmt.Println(b)

	// 정해진 타입으로 매개변수 넣어줘야함
	// testConvertInt(b)
	testConvertInt(int(b))
	testConvertCnt(b)
	testConvertCnt(cnt(10))

}
