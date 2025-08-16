package packagetest

// import "fmt"
// import "os"

import (
	"fmt"
	"os"
)

func Packagetest06() {
	// package: 코드 구조화(모듈) 및 재사용
	// 느슨한 결합도, 패키지 단위의 독립적이고 작은 단위로 개발, 작은 패키지를 결합해서 프로그램을 작성할 것을 권고 (클린코드)
	// 패키지 이름 == 디렉토리 이름
	// 같은 패키지 내 소스파일들은 디렉토리명을 패키지 명으로 사용
	// 네이밍 규칙: 소문자 private. 대문자 public
	// main 패키지는 특별하게 인식, 컴파일러는 공유 라이브러리가 아닌 프로그램의 시작점(entrypoint)으로 생각함

	var name string
	fmt.Println("이름: ")
	fmt.Scanf("%s", &name)
	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)
}
