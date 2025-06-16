package basic01

import (
	"fmt"
	"math/rand"
	"time"
)

func Switch01() {
	// 제어문(조건문) - switch
	// switch 뒤 표현식(expression) 생략 가능
	// case 뒤 expression 사용 가능

	a := 7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	switch c := "go"; c {
	case "go":
		fmt.Println("go!")
	case "java":
		fmt.Println("java!")
	default:
		fmt.Println("No Match!")
	}

	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("golang!")
	case "java":
		fmt.Println("java!")
	default:
		fmt.Println("None!")
	}

	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i < j!")
	case i == j:
		fmt.Println("i == j!")
	case i > j:
		fmt.Println("i > j!")
	}

}

func Switch02() {
	// case 1
	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, "over 50 and under 100")
	case i > 25 && i < 50:
		fmt.Println("i -> ", i, "over 25 and under 50")
	default:
		fmt.Println("i -> ", i, "default")
	}
}

func Switch03() {
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		if a >= 20 {
			fmt.Println("over 20")
		}
		fmt.Println("a -> ", a, "는 홀수")
	}

	switch e := "go"; e {
	case "java":
		fmt.Println("Java!")
	case "go":
		fmt.Println("Go!")
		fallthrough
	case "python":
		fmt.Println("Python!")
	case "ruby":
		fmt.Println("Ruby!")
	case "c":
		fmt.Println("C!")
	}
}
