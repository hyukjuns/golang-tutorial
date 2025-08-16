package structtest

import "fmt"

type spec struct {
	length int
	height int
	width  int
}

type CarSpec struct {
	name    string
	color   string
	company string
	detail  spec
}

func StructTest6() {
	car1 := CarSpec{
		"mycar",
		"foo",
		"bar",
		spec{1, 2, 3},
	}
	fmt.Println(car1.name, car1.color, car1.company)
	fmt.Printf("%#v\n", car1.detail)
	fmt.Println()

	fmt.Println(car1.detail.height, car1.detail.length, car1.detail.width)
}
