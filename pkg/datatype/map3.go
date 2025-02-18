package datatype

import "fmt"

func Map3() {
	map1 := map[string]string{
		"daum":   "https://daum.net",
		"naver":  "https://naver.com",
		"google": "https://google.com",
		"home1":  "http//test1.com",
	}
	fmt.Println(map1)

	// 추가
	map1["home2"] = "http://home2.com"
	fmt.Println(map1)

	// 변경
	map1["home2"] = "http://homesweethome.com"
	fmt.Println(map1)

	// 삭제
	delete(map1, "home1")
	delete(map1, "home2")
	fmt.Println(map1)

}
