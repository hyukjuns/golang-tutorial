package datatype

import (
	"fmt"
	"math"
)

func Datatype08() {
	// 다른 데이터 타입 간에는 반드시 형 변환 후 연산 해야 함
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("n1: ", n1)
	fmt.Println("n2: ", n2)
	fmt.Println("n3: ", n3)
	fmt.Println("n4: ", n4)

	fmt.Println("int8: ", math.MaxInt8)
	fmt.Println("int16: ", math.MaxInt16)
	fmt.Println("int32: ", math.MaxInt32)
	fmt.Println("int64: ", math.MaxInt64)

	n5 := 100000 // int
	n6 := int16(10000)
	n7 := uint8(100)

	fmt.Println(n5, n6, n7)

	// 형 변환 없으면 다른 데이터타입 간에 연산 불가능
	fmt.Println(n5 + int(n6))

	// 비교 연산도 데이터타입이 같아야 함
	fmt.Println(n6 > int16(n7))

}
