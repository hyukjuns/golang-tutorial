package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 문자열
	// ""(큰따음표) | `` (백스퀏)
	// char 타입 제공하지 않음
	// rune(int32(4byte))로 문자 코드값 표현
	// 문자는 '' (작은따음표)로 작성 가능
	// \ Escape
	// \b 백스페이스
	// \f 쪽바꿈
	// \n 개행
	// \r 복귀
	// \t 탭

	var str1 string = "c:\\gopath\\src\\"
	var str2 string = `c:\gopath\src\`
	fmt.Println(str1)
	fmt.Println(str2)

	var str3 string = "Hello, World!"
	var str4 string = "안녕하세요"
	var str5 string = "\ud55c\uae00"
	fmt.Println(str3, str4, str5)

	// len 은 bytes의 수
	fmt.Println(len(str3)) // 영단어 1개는 1 byte
	fmt.Println(len(str4)) // 한글 1개는 3 byte

	// 실제 문자열 길이는 utf8 패키지 사용
	fmt.Println(utf8.RuneCountInString(str3))
	fmt.Println(utf8.RuneCountInString(str4))
	fmt.Println(len([]rune(str4)))

	// go는 기본적으로 utf-8 인코딩 (유니코드 문자 집합)
	// 한글 1단어 == 3 byte
	// 영어 1단어 == 1 byte

	var str12 string = "Golang"
	var str22 string = "World"
	var str32 string = "고프로그래밍"

	fmt.Printf("%s!!!!!!!!!!!!!\n", str12)
	fmt.Println(str12)
	// 유니코드 출력
	fmt.Println(str12[0], str12[1], str12[2], str12[3], str12[4], str12[5])
	fmt.Println(str22[0], str22[1], str22[2], str22[3], str22[4])
	fmt.Println(str32[0], str32[1], str32[2], str32[3], str32[4], str32[5])

	// 문자 출력
	fmt.Printf("%c %c %c %c %c %c\n", str12[0], str12[1], str12[2], str12[3], str12[4], str12[5])
	fmt.Printf("%c %c %c %c %c\n", str22[0], str22[1], str22[2], str22[3], str22[4])
	fmt.Printf("%c %c %c %c %c %c\n", str32[0], str32[1], str32[2], str32[3], str32[4], str32[5])

	// 한글 깨질 시 rune 형 사용
	conStr := []rune(str32)
	fmt.Printf("%c %c %c %c %c %c\n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])

	for i, char := range str12 {
		fmt.Printf("%c(%d)\t", char, i)
	}
	fmt.Println()
	for i, char := range str22 {
		fmt.Printf("%c(%d)\t", char, i)
	}
}
