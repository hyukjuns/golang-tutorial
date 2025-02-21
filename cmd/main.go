package main

import (
	"dev/pkg/functest"
	"fmt"
)

func main() {

	fmt.Println(functest.Multiply(1, 2, 3))
	fmt.Println(functest.PrtWord("h", "e", "l", "l", "o"))

	test := []int{1, 2, 3, 4, 5}
	result := functest.Multiply(test...)
	fmt.Println(result)

	res1, res2 := functest.SumAndMulResult(10, 10, functest.Sum, functest.Mul)
	fmt.Println("sum:", res1, "mul:", res2)
}
