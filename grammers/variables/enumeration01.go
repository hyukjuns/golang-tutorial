package variables

import "fmt"

func Enum1() {
	const (
		Jan = 1
		Feb = 2
		Mar = 3
		Apr = 4
		May = 5
		Jun = 6
	)

	fmt.Println(Jan, Feb, Mar, Apr, May, Jun)

}
