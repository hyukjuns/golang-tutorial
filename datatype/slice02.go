package main

import "fmt"

func main() {
	// 배열/슬라이스 복사/참조

	// 배열은 값 복사
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int
	arr2 = arr1
	arr2[0] = 7

	fmt.Println(arr1)
	fmt.Println(arr2)

	// 슬라이스는 메모리 주소 참조 (동일한 슬라이스)
	slice1 := []int{1, 2, 3}
	var slice2 []int
	slice2 = slice1
	slice2[0] = 7

	fmt.Println(slice1)
	fmt.Println(slice2)

	// 예외
	slice3 := make([]int, 50, 100)
	fmt.Println(slice3[4])
	// fmt.Println(slice3[50]) // index out of range [50] with length 50

	for i, v := range slice1 {
		fmt.Println(i, v)
	}
}
