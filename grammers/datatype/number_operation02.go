package datatype

import (
	"fmt"
	"math"
)

func Datatype09() {
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println(n1 % n2)
	fmt.Println(n1 << 2)
	fmt.Println(n1 >> 2)
	fmt.Println(^n1)

	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	// 실수형으로 형변환 해야 소수점 살림
	fmt.Println(float32(n3) + n4)
	fmt.Println(n3 + int(n4))
	fmt.Println(math.MaxUint16)
	fmt.Println(math.MaxInt)
	fmt.Println(math.MaxInt64)
	fmt.Println(uint16(n6))
	fmt.Println(n3, n4, n5, n6)
	fmt.Println(math.MaxFloat64)

}
