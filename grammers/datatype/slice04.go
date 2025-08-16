package datatype

import (
	"fmt"
	"sort"
)

func Datatype17() {
	// 슬라이스 추출 및 정렬
	// slice[i:j] -> i 부터 j-1 까지 추출
	// slice[i:] -> i부터 마지막 까지 추출
	// slice[:j] -> 처음 부터 j-1 까지 추출
	// slice[:] -> 처음부터 마지막까지 추출

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice1[0:3])
	fmt.Println(slice1[0:])
	fmt.Println(slice1[:3])
	fmt.Println(slice1[0:len(slice1)])

	slice2 := []int{3, 5, 1, 4, 2}
	slice3 := []string{"b", "a", "d", "c", "e"}

	fmt.Println(sort.IntsAreSorted(slice2)) // true or false
	sort.Ints(slice2)
	fmt.Println(slice2)
	fmt.Println(sort.IntsAreSorted(slice2)) // true or false

	fmt.Println(sort.StringsAreSorted(slice3)) // true or false
	sort.Strings(slice3)
	fmt.Println(slice3)
	fmt.Println(sort.StringsAreSorted(slice3)) // true or false

}
