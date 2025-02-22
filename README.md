# golang-tutorial

# Structure
/cmd: main 함수 배치
/pkg: 외부 호출될 소스 배치 사용

# Function
- Closure: 외부함수의 변수를 참조하는 익명함수를 리턴하는 함수 호출시 클로져가 생성됨, 클로저는 외부함수의 지역변수를 캡쳐하여 객체로서 메모리에 상주시키고, 객체화된 외부 지역변수를 계속해서 사용할 수 있음 (비동기작업, 누적함수에 사용됨, 메모리 누수 주의)
- Defer: 함수 호출시 실행 지연기능, Defer로 호출한 함수는 메인 쓰레드가 종료되기 직전에 호출된다(finally와 유사)
- 함수를 함수의 인자로 받아서 사용 가능
- 변수(슬라이스,맵,변수)에 함수 할당 가능
- 재귀호출, 즉시실행함수(익명함수) 사용 가능
- 선언방식
```go
- func Test() {}
- func Test() RETURN_TYPE {}
- func Test{VAR_NAME VAR_TYPE} {}
- func Test(VAR_NAME VAR_TYPE) RETURN_TYPE {}
- func Test(VAR_NAME VAR_TYPE) (RETURN_TYPE,RETURN_TYPE) {} // 멀티 리턴
- func Test(VAR_NAME VAR_TYPE) (RETURN_VAR_NAME_1 RETURN_TYPE,- RETURN_VAR_NAME_2 RETURN_TYPE) {} // 리턴값 이름 지정
- func Test(VAR_NAME VAR_TYPE, ARG_FUNC) (RETURN_TYPE,RETURN_TYPE) {} // 함수를 인자로 받음
- func Test(VAR_NAME ...VAR_TYPE) RETURN_TYPE {} // 가변 매개변수
```


# Pointer
- 변수의 지역성, 연속된 메모리 참조, 힙, 스택 등에 사용
- 주소의 값은 직접 변경 불가능(코딩 실수로 인한 버그 방지)
- * 표시로 포인터임을 명시
- nil로 초기화
- 함수 매개변수는 기본적으로 지역 변수로 값 복사
- 원본 변수값 변경을 위해서 포인터로 전달
- 크기가 큰 배열의 경우 값 복사시 리소스 부하 -> 포인터 전달을 통해 참조형식으로 리소스 부하 해결

# Pacakge

1. pkg/DIR/FILE.go
    
    ```go
    package DIR
    func Funcname() // 함수 첫 단어는 대문자(퍼블릭 함수)
    ```

2. go mod init MODULE

3. cmd/main.go

    ```go
    import MODLUE/pkg/DIR
    DIR.FUNCNAME()
    ```

# Array vs Slice
- 길이고정 vs 길이가변
- 값 타입 vs 참조 타입
- 대입연산자: 복사 전달  vs 참조 값 전달
- 비교연산자 허용 vs 비교연산자 불가
- len vs len, cap

- Array: 
    - arr1 := [5]int{1,2,3,4,5}

- Slice: 
    - slice1 := []int{1,2,3,4,5}
    - slice1 = append(slice1, 6, 7)
    - 용량 초과시 추가 용량은 초기 용량의 X2
    - 정렬확인: sort.IntsAreSorted(SLICE) // bool
    - 정렬: sort.Ints(SLICE) | sort.Strings(SLICE)
    - COPY 전제조건: make로 만든 슬라이스만 카피 가능 (용량이 지정되어 있는 슬라이스만 카피 가능)
    - COPY 슬라이스 값복사 방법: copy(slice2, slice1) // slice2의 용량이 5이기때문에, 5까지만 복사함
    - COPY2 참조복사 방법 2: SLICE2 = SLICE1[0:3] //  참조복사, 배열 혹은 슬라이스를 부분 추출해서 카피하면 참조 복사가 됨
    - COPY3 용량지정: SLICE2 = SLICE1[start:end:capacity] // 용량을 명시적으로 카피할 수 있음

# Map:
Hashtable, Dictionary(Python) 형식, 즉, Key: Value 자료형
레퍼런스 타입으로 참조값 전달
선언은 make 함수 및 축약 방식 사용 가능
순서 없음
축약 선언 예시:
map3 := make(map[string]int)
초기화 선언 예시:
map5 := map[string]int{
		"apple":  15,
		"banana": 200,
	}


## References
[Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ko.md)