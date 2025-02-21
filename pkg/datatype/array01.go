package main

import "fmt"

func main() {

	// Array vs Slice
	// 길이고정 vs 길이가변
	// 값 타입 vs 참조 타입
	// 복사 전달  vs 참조 값 전달
	// 비교연산자 허용 vs 비교연산자 불가
	// Array 는 값 복사
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1

	fmt.Println(arr1, &arr1)
	fmt.Println(arr2, &arr2)

	fmt.Printf("%v %p\n", arr1, &arr1)
	fmt.Printf("%v %p\n", arr2, &arr2)

	arr2[2] = 10
	fmt.Printf("%v %p\n", arr1, &arr1)
	fmt.Printf("%v %p\n", arr2, &arr2)

	var arr3 [5]int
	var arr4 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr5 = [5]int{1, 2, 3, 4, 5}
	arr6 := [5]int{1, 2, 3, 4, 5}
	arr7 := [5]int{1, 2, 3}
	arr8 := [...]int{1, 2, 3, 4, 5}
	arr9 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	fmt.Println(arr3, arr4, arr5, arr6, arr7, arr8, arr9)
	fmt.Println(len(arr3), len(arr9[1]))
}
