package structtest

import "fmt"

// pricate struct
type car struct {
	color string
	name  string
}

func StructTest3() {
	// 함수의 매개변수 전달시 항상 값복사 이므로,
	// 원본 참조가 필요한 경우 포인터 형으로 넘겨야 한다.
	c1 := car{"red", "test"}
	c2 := new(car)
	c2.color, c2.name = "white", "sonata"
	c3 := &car{}
	c4 := &car{"black", "test2"}

	fmt.Println(c1, c2, c3, c4)
	fmt.Printf("%#v,%#v,%#v,%#v", c1, c2, c3, c4)

}
