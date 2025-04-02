# golang-tutorial

## Directory Structure
/cmd: main 함수 배치
/pkg: 외부 호출될 소스 배치 사용

## Go 병행처리
### Sync, 원자성
고루틴 흐름을 직접 관리, Do-Once, Add-Done-Wait 활용하여 흐름을 직접 관리

원자성: 기능적으로 분할 불가능한 완정 보증된 일련의 조작, 모두 성공 or 모두 실패, pkg: sync/atomic 에서 원자적 연산자 제공, 주로 공용변수에 관한 계산에서 사용

*A WaitGroup waits for a collection of goroutines to finish. 
https://pkg.go.dev/sync#WaitGroup

### Gorutine, Mutex
순차적인 main 함수 흐름에서 벗어난 병렬 실행 흐름을 Gorutine으로 만들수 있음, 고루틴 사용시 공유 데이터 동기화에 문제가 발생할 수 있는데, 이때 Mutex를 사용하면 개별 실행 흐름에서 공유 데이터 동기화를 위한 데이터 제어가 가능해짐

### Mutex
고루틴 사용시 데이터 동기화(타이밍, 싱크)를 위해 사용
RWMutex(읽기/쓰기 뮤텍스 잠금), RMutex(읽기 뮤텍스 잠금)가 존재
```
# Mutex
mutex := new(sync.Mutex)
# RWMutex
mutex := new(sync.RWMutex)
# RMutex
mutex := new(sync.RMutex)

# 동기화 상태(조건) 메소드
wait, sigal, broadcast
```

### Goroutine
타 언어의 쓰레드 개념, 생성방법이 매우 간단(go 예약어)하며 리소스를 매우 적게 사용함(수많은 고루틴 동시 생성 및 실행 가능)
비동기적 함수 루틴 실행에 사용함 (함수간에는 채널을 통해 통신함)
공유 메모리 사용 시 정확한 동기화 코딩이 필요하고, 싱글루틴에 비해 항상 빠른 처리 결과는 아니다.

채널(Channel)
채널은 동기적으로 실행됨, 레퍼런스타입 값 전달(참조전달)
고루틴간에 상호 정보(데이터) 교환 및 실행 흐름 동기화를 위해 사용
실행 흐름 제어 가능(동기, 비동기), 일반 변수로 선언 후 사용 가능
데이터 전달 자료형 선언 후 사용(지정된 타입만 주고 받을 수 있음)
interface{} 전달을 통해 자료형 상관없이 전송 및 수신 가능, 포인터(슬리이스, 맵)을 전달하는 경우 주의
멀티프로세싱 처리에서 교착상태(경합) 발생 주의

채널간 통신시 버퍼를 사용하여 비동기 처리가능(큐 처럼 사용)
(수신과 송신에 버퍼(큐) 사용)
`ch := make(chan bool, 4)`

close로 채널을 닫기 전까지, range로 채널의 값을 읽을 수 있음
```go
ch := make(chan bool)
go func() {
	for i := 0; i < 5; i++ {
		ch <- true
	}
	close(ch) // 5회 채널에 값 전송 후 채널 닫기
}()

// 채널이 닫힐때까지 값을 수신
for i := range ch {
	fmt.Println(i)
}
```

```bash
송신: 채널 <- 데이터
수신:     <- 채널
```

```go
go FunctionName()
```
## Go 객체지향
### Interface

![](/concepts/interface.png)

인터페이스에는 메소드만 구현하고, 해당 인터페이스의 메소드들을 리시버 함수로 갖는 구조체들은 모두 해당 인터페이스로 호출할 수 있으며 이를 덕타이핑 이라 부름 (오리처럼 걷고, 짖고, 헤엄치는 것은 오리라고 본다는 뜻)

*덕타이핑: 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식

인터페이스의 메소드를 리시버 함수로 구현한 구조체는 해당 인터페이스로서 사용될 수 있게됨

즉, 인터페이스는 객체의 동작과 골격을 표현하고 단순히 동작에 대한 방법만 묘사하므로 추상화를 제공한다고 볼 수 있고, 이로써 프로그래밍의 유연성이 좋아지게됨

만약 빈 인터페이스를 사용할 경우 Type Assertion(타입 변환)을 통해 타입 변환 후 사용해야 함

`interfaceValue.(type)`

```go
type INTERFACENAME interface {
    method1() RETURNTYPE
    method2()
}
```

### Struct

Golang의 객체지향 타입은 구조체로 정의할 수 있으며 타 언어와 비교하면 클래스, 상속 개념은 없지만 상태와 메소드를 분리해서 정의하고 사용할 수 있음 (결합성 없음)
즉, 구조체와 메서드간 연견을 통해서 타 언어의 클래스 형식처럼 사용 가능(객체지향적)

사용자 정의 타입은 사용자가 임의로 선언한 타입 생성 방식이며 Golang이 제공하는 구조체, 인터페이스, 기본타입, 함수 등 다양한 타입을 임의의 타입으로 재구성하여 사용할 수 있도록 함

Receiver는 구조체와 연결된 함수를 의미하며 연결된 구조체의 메소드처럼 사용 가능함, 리시버도 일반 함수와 마찬가지로 구조체의 포인터를 전달해서 실제 원본의 값을 수정할 수 있음

또한, 구조체 임베디드 패턴을 사용하여 메소드 오버라이딩을 구성할 수 있음

### Function (함수)

Closure: 외부함수의 변수를 참조하는 익명함수를 리턴하는 함수 호출시 클로져가 생성됨, 클로저는 외부함수의 지역변수를 캡쳐하여 객체로서 메모리에 상주시키고, 객체화된 외부 지역변수를 계속해서 사용할 수 있음 (비동기작업, 누적함수에 사용됨, 메모리 누수 주의)

Defer: 함수 호출시 실행을 지연시키는 기능, Defer로 호출한 함수는 메인 쓰레드가 종료되기 직전에 호출됨(finally와 유사하며, Stack으로 생각하면 이해하기 쉬움)

익명함수: 함수 선언과 동시에 실행 되는 함수로서 클로저의 밑바탕이 됨

기타 함수의 매개변수는 함수, 일반 변수 등 다양하게 사용 가능하고, 리턴값이 여러개일 수 있음

참고로 함수의 첫글자가 대문자이면 퍼블릭, 소문자이면 프라이빗 함수로 선언됨

### Pointer
원본 변수의 메모리 주소를 담고 있는 자료형

"*" 표시로 포인터임을 명시하며, 참조전달을 위해 자주 사용됨, Golang에서는 포인터를 사용한 메모리 주소 변경은 불가능함

사용 사례 로는 만약 크기가 큰 배열을 함수로 전달해야하는 경우 값 전달은 변수가 복사되므로 메모리가 낭비되지만, 포인터로 넘기면 참조형식이므로 메모리 부하를 상대적으로 줄일 수 있음

### Pacakge 구조
메인 함수에서 외부 패키지 함수를 호출할 경우 아래 스텝으로 진행

1. 현재 작업 디렉토리에서 Module 초기화, go.mod 생성됨

```bash
go mod init MODULE
```

2. 패키지 함수 작성

Path: pkg/DIR/FILE.go
    
```go
package DIR // 디렉터리 이름
func Funcname() // 함수 첫 단어는 대문자로 선언하여 외부에서 참조되도록 구현 (퍼블릭 함수)
```

3. 메인함수에서 호출하여 사용

cmd/main.go

```go
import MODLUE/pkg/DIR // 모듈명/pkg/디렉터리
DIR.FUNCNAME()
```

### Array vs Slice
Array: 길이 고정, 대입 연산시 값 복사됨, len 개념 존재

```go
arr1 := [5]int{1,2,3,4,5} // 선언
```

Slice: 길이 가변, 대입 연산시 참조(주소값)이 전달됨, len, cap 개념 존재

len, cap: len은 슬라이스의 길이, cap은 슬라이스의 용량, cap이 있음으로서 메모리 주소공간의 낭비를 막을 수 있고, 만약 슬라이스가 확장되면서 cap 용량이  초과되면 추가 용량은 초기 용량의 x2로 늘어남

```go
// 선언
slice1 := []int{1,2,3,4,5}
slcie2 := make([]int, 50, 100) // make(슬라이스 자료형, len, cap)
// 슬라이스 확장
slice3 = append(slice1, 6, 7)
slice3 = append(slice1, slice2...)
// 참조복사
slice2 = slice1[start:end:capacity]
// 슬라이스의 값 복사 방법
copy(slice2, slice1)
```

### Map
Golang의 Key: Value 자료형으로서 순서가 없는 자료형 이며 Hashtable, Dictionary(Python) 형식으로 사용되고 대입 연산시 참조값이 전달됨

```go
// map[KEY]VALUE
map1 := make(map[string]int)
map2 := map[string]int {
		"apple":  15,
		"banana": 200,
	}
```

## References
[Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ko.md)