package datatype

import "fmt"

func Datatype04() {
	// Map
	// Hashtable, Dictionary(Python)
	// Key-Value로 자료 저장
	// 레퍼런스 타입 (참조 값 전달)
	// 비교 연산자 사용 불가능 (참조 타입 이므로)
	// 특징: 참조타입(Key)으로 사용 불가능, 값(Value)로 모든 타입 사용 가능
	// make 함수 및 축약(리터럴)로 초기화 가능
	// 순서 없음

	// 기본 선언
	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int)                // 자료형 생략
	map3 := make(map[string]int)                   // 축약

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	// 선언 및 초기화
	map4 := map[string]int{}
	map4["apple"] = 25
	map4["banana"] = 100

	map5 := map[string]int{
		"apple":  15,
		"banana": 200,
	}
	fmt.Println(map4, map5)

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 100
	fmt.Println(map6)

}
