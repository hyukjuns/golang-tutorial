package packagetest

import (
	mylib "basic03/lib" // alias "PACKAGE"
	_ "basic03/lib2"
	"fmt"
)

func Packagetest02() {
	// 패키지 접근제어
	// 별칭사용(alias)
	// 빈 식별자 사용

	fmt.Println("over 10?: ", mylib.CheckNum(11))

}
