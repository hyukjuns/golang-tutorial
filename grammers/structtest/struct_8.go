package structtest

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

func (e Excutives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Excutives struct {
	Employee     // is a
	specialBonus float64
}

func StructTest8() {
	ep1 := Employee{"kim", 10000000, 1000000}
	ep2 := Employee{"lee", 10000000, 1000000}
	ex := Excutives{
		Employee{"park", 20000000, 2000000},
		1000000,
	}

	fmt.Println(int(ep1.Calculate()))
	fmt.Println(int(ep2.Calculate()))
	fmt.Println(int(ex.Calculate()))

}
