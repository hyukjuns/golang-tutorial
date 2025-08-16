package datatype

import "fmt"

func Datatype07() {
	// 맵 조회 시 주의할 점
	map1 := map[string]int{
		"apple":     15,
		"google":    20,
		"microsoft": 100,
	}

	// Key 가 존재하면 값 1개 리턴 -> Value
	// Key 가 존재하지 않으면 값 2개 리턴 -> 0, false
	value1 := map1["apple"]
	value2 := map1["samsung"]
	value3, ok := map1["samsung"] // 두번째 리턴값으로 키 존재 유무 확인
	fmt.Println(value1, value2, value3, ok)

	if value, exist := map1["samsung"]; exist {
		fmt.Println(value)
	} else {
		fmt.Println("Key is not exist!")
	}

	if value, exist := map1["apple"]; exist {
		fmt.Println(value)
	} else {
		fmt.Println("Key is not exist!")
	}

	if _, exist := map1["apple"]; exist {
		fmt.Println(exist)
	} else {
		fmt.Println("Key is not exist!")
	}

}
