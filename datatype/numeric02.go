package main

import "fmt"

func main() {
	// 정수형 문자

	// byte == uint8
	var char1 byte = 72
	var char2 byte = 0110
	var char3 byte = 0x48

	// rune: 유니코드(한글)
	var char4 rune = '가'
	var char5 rune = 0142574
	var char6 rune = 0xC57C

	// %c, %d, %o, %x
	fmt.Printf("%c %c %c \n", char1, char2, char3)
	fmt.Printf("%d %d %d \n", char1, char2, char3)
	fmt.Printf("%d %o %x \n", char1, char2, char3)

	fmt.Printf("%c %c %c \n", char4, char5, char6)
	fmt.Printf("%d %d %d \n", char4, char5, char6)
	fmt.Printf("%d %o %x \n", char4, char5, char6)

}
