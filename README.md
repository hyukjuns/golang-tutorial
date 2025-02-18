# golang-tutorial


```markdown
/cmd: main 함수 배치
/pkg: 외부 호출될 소스 배치 사용

Pacakge

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

Array vs Slice
- 길이고정 vs 길이가변
- 값 타입 vs 참조 타입
- 대입연산자: 복사 전달  vs 참조 값 전달
- 비교연산자 허용 vs 비교연산자 불가
- len vs len, cap

Array: 
    - arr1 := [5]int{1,2,3,4,5}

Slice: 
    - slice1 := []int{1,2,3,4,5}
    - slice1 = append(slice1, 6, 7)
    - 용량 초과시 추가 용량은 초기 용량의 X2
    - 정렬확인: sort.IntsAreSorted(SLICE) // bool
    - 정렬: sort.Ints(SLICE) | sort.Strings(SLICE)
    - COPY 전제조건: make로 만든 슬라이스만 카피 가능 (용량이 지정되어 있는 슬라이스만 카피 가능)
    - COPY 슬라이스 값복사 방법: copy(slice2, slice1) // slice2의 용량이 5이기때문에, 5까지만 복사함
    - COPY2 참조복사 방법 2: SLICE2 = SLICE1[0:3] //  참조복사, 배열 혹은 슬라이스를 부분 추출해서 카피하면 참조 복사가 됨
    - COPY3 용량지정: SLICE2 = SLICE1[start:end:capacity] // 용량을 명시적으로 카피할 수 있음

Map:
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
```

## References
[Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ko.md)