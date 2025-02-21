package functest

func Sum(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func SumAndMulResult(a, b int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(a, b), f2(a, b)
}
