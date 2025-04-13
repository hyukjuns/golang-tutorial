package variables

import "fmt"

func Enum3() {

	const (
		_ = iota
		A
		_
		C
	)
	fmt.Println(A, C)
	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		_
		PLATINUM
	)
	fmt.Println(DEFAULT, SILVER, PLATINUM)
}
