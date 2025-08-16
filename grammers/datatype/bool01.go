package datatype

import "fmt"

func Datatype02() {
	// 암묵적 형변환 x (0, nil -> false 변환 없음)
	var b1 bool = true
	var b2 bool = false
	b3 := true
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("b3 == b3:", b3 == b3)
	fmt.Println("b1 && b2:", b1 && b2)
	fmt.Println("b1 || b2:", b1 || b2)
	fmt.Println("!b1:", !b1)

}
