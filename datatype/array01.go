package main

import "fmt"

func main() {

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
