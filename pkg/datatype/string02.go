package main

import (
	"fmt"
	"strings"
)

func main() {
	// 문자열 연산
	// 추출 | 비교 | 조합(결합)

	var str1 string = "Golang"
	var str2 string = "World"

	// 배열 슬라이싱 -> 문자열 추출
	// 배열 인덱스 -> 유니코드
	fmt.Println(str1[0:2], str1[0])
	fmt.Println(str2[3:], str2[0])

	fmt.Println(str2[:4])
	fmt.Println(str1[1:3])

	// golang은 아스키 코드에 대한 사전식 비교를 함
	// 바이트 크기 비교가 아님
	str3 := "Golang" // 6 byte
	str4 := "World"  // 5 byte
	fmt.Println(str3 == str4)
	fmt.Println(str3 != str4)
	fmt.Println(str3 > str4)
	fmt.Println(str3 < str4)
	fmt.Println("abc" > "ABC")

	// 문자열 결합 -> Join
	// String은 Immutable 하기 때문에 메모리에서 수정이 불가, 따라서, 문자열 연산시 인스턴스가 계속 생성됨
	// 따라서 string buffer 와 같은 메소드를 사용해야 성능면에서 효율적임
	str5 := "Build simple, secure, scalable systems with Go" +
		"An open-source programming language supported by Google" +
		"Easy to learn and great for teams" +
		"Built-in concurrency and a robust standard library" +
		"Large ecosystem of partners, communities, and tools"

	str6 := "Organizations in every industry use Go to power their software and services "
	fmt.Println(str5 + str6)

	// 성능 효율적인 방법은 append와 join 사용
	strSet := []string{}
	strSet = append(strSet, str5)
	strSet = append(strSet, str6)
	// strSet에 있는 문자열을 어떻게 결합할건지 결정(구분자)
	fmt.Println(strings.Join(strSet, "||||||||||||||||"))

}
