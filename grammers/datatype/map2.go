package datatype

import "fmt"

func Datatype05() {
	// 맵 조회 및 순회
	map1 := map[string]string{
		"daum":   "https://daum.net",
		"naver":  "https://naver.com",
		"google": "https://google.com",
	}

	fmt.Println(map1["google"])

	// 순회: 순서 없으므로 랜덤
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	for _, v := range map1 {
		fmt.Println(v)
	}

}
