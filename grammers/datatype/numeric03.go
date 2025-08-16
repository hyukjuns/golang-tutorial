package datatype

import "fmt"

func Datatype13() {
	// 실수(부동소수점)
	// float32(7자리), float64(15자리)

	var num1 float32 = 0.14
	var num2 float32 = .75647
	var num3 float32 = 442.0378621123
	var num4 float32 = 10.0

	// 지수 표기법
	var num5 float32 = 14e6
	var num6 float64 = .156875e+3
	var num7 float64 = 4.32421e-10

	fmt.Println("ex: ", num1)
	fmt.Println("ex: ", num2)
	fmt.Println("ex: ", num3)
	fmt.Println("ex: ", num4)
	fmt.Println("ex: ", num4-0.1)
	fmt.Println("ex: ", float32(num4-0.1))
	fmt.Println("ex: ", float64(num4-0.1))
	fmt.Println("ex: ", num5)
	fmt.Println("ex: ", num6)
	fmt.Println("ex: ", num7)

}
