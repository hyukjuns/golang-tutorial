package main

import "fmt"

func main() {
	// 상수
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	// const d = getHeight()
	const e = 35.6
	const f = false
	/*
		에러 발생
		const g string
		g = "Test3"
	*/

	fmt.Println("a,b,c,e,f:", a, b, c, e, f)
}
