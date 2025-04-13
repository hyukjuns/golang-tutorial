package variables

import "fmt"

func Enum2() {
	const (
		A = iota * 10
		B
		C
	)

	fmt.Println(A, B, C) // 0 10 20

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
	)
	fmt.Println(Jan, Feb, Mar, Apr, May, Jun) // 1 2 3 4 5 6
}
