# golang-tutorial


Package
1. DIR/FILE.go

    `package DIR`
2. go mod init MODULE

    `import MODLUE/DIR`

Array vs Slice
- 길이고정 vs 길이가변
- 값 타입 vs 참조 타입
- 복사 전달  vs 참조 값 전달
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
    - COPY-값복사: copy(slice2, slice1) // slice2의 용량이 5이기때문에, 5까지만 복사함
    - COPY2-참조복사: SLICE2 = SLICE1[0:3] //  참조복사, 배열 혹은 슬라이스를 부분 추출해서 카피하면 참조 복사가 됨
    - COPY3-용량지정: SLICE2 = SLICE1[start:end:capacity] // 용량을 명시적으로 카피할 수 있음
