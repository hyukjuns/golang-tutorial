# golang-tutorial

```markdown
# 패키지 생성 및 임포트 방법
# 보통 pkg(외부), internal(내부) 디렉터리 아래 사용할 lib 배치 후 main에서 경로 지정하여 사용
# 퍼블릭 패키지는 함수명을 대문자로 시작 (프라이빗은 소문자로)
# pkg/DIRNAME/FILENAME.go
package DIRNAME
func UPPER_NAME () RETURN {
    return xxx
}
# main.go
# 프로젝트 디렉터리에서 go mod init MODULENAME 하고 나서, 해당 MODULENAME 을 루트디렉터리 경로로 잡아줘야 함
import (
    "MODULENAME/pkg/DIRNAME"
)

# Make Package(Binary)
go mod init
go install [PACKAGE]
```