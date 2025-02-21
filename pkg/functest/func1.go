package functest

func Multiply(n ...int) int { // 가변 매개변수
	tot := 1

	for _, value := range n {
		tot *= value
	}

	return tot
}

func PrtWord(msg ...string) string {
	result := ""
	for _, message := range msg {
		result += message
	}
	return result
}
