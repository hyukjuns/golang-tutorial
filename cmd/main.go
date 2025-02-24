package main

import (
	"dev/pkg/functest"
	"dev/pkg/structtest"
	"fmt"
)

func funcSample() {
	fmt.Println(functest.Multiply(1, 2, 3))
	fmt.Println(functest.PrtWord2("h", "e", "l", "l", "o"))

	test := []int{1, 2, 3, 4, 5}
	result := functest.Multiply(test...)
	fmt.Println(result)

	res1, res2 := functest.SumAndMulResult(10, 10, functest.Sum, functest.Mul)
	fmt.Println("sum:", res1, "mul:", res2)

	fmt.Println("---변수에 함수 할당---")
	fmt.Println("슬라이스에 함수 할당")

	f := []func(int, int) int{functest.Mul1, functest.Sum1}
	a := f[0](10, 10)
	b := f[1](10, 10)

	fmt.Println(a, b)

	fmt.Println("변수에 함수 할당")
	var f1 func(int, int) int = functest.Mul1
	f2 := functest.Sum1
	fmt.Println(f1(10, 10), f2(10, 10))

	fmt.Println("Map에 함수 할당")
	f3 := map[string]func(int, int) int{
		"Mul_func": functest.Mul1,
		"Sum_func": functest.Sum1,
	}
	fmt.Println(f3["Mul_func"](10, 10), f3["Sum_func"](10, 10))

	fmt.Println("재귀함수")
	fmt.Println(functest.Factorial(5))
	functest.PrtHello(5)

	fmt.Println("익명 함수")
	functest.AnonymousFunc()
}
func funcSampleDefer() {
	functest.Defer1()
	functest.Defer3("Defer Tests")
	functest.Stack()
	functest.DupleFunc()
}
func closureSample() {
	functest.ClosureTest2()
}
func main() {
	structtest.UserType4()
}
