package interfacetest

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// Dog receiver method
func (d Dog) bite() {
	fmt.Println(d.name, d.weight, "bite!!!")
}

func (d Dog) sounds() {
	fmt.Println(d.name, d.weight, "barks!!!")
}

func (d Dog) run() {
	fmt.Println(d.name, d.weight, "running!!!")
}
func (d Dog) cry() {
	fmt.Println(d.name, d.weight, "crying!!!")
}

// Cat receiver method
func (c Cat) bite() {
	fmt.Println(c.name, c.weight, "bite!!!")
}

func (c Cat) sounds() {
	fmt.Println(c.name, c.weight, "barks!!!")
}

func (c Cat) run() {
	fmt.Println(c.name, c.weight, "running!!!")
}
func (c Cat) cry() {
	fmt.Println(c.name, c.weight, "crying!!!")
}

func act(animal Behavior) {
	animal.bite()
	animal.sounds()
	animal.run()
}

// 익명 인터페이스, 타입을 정의하지 않음
func actAdvanced(animal interface{ cry() }) {
	animal.cry()
}

// interface 매개변수는 모든 타입이 허용됨
func printValue(s interface{}) {
	fmt.Println(s)
}

// interface
type Behavior interface {
	bite()
	sounds()
	run()
}

func InterfaceTest1() {

	/*
		인터페이스는 객체의 동작과 골격을 표현
		단순히 동작에 대한 방법만 묘사하므로 추상화를 제공
		인터페이스의 메소드를 구현한 타입(구조체)는 인터페이스로 사용가능
		이를 덕타이핑 이라하며, Golang의 유연성을 대표하는 기능
		덕타이핑: 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식

		type INTERFACENAME interface {
			method1() RETURNTYPE
			method2()
		}
	*/

	/*
		// 선언 1
		dog1 := Dog{"poll", 10}
		var inter1 Behavior
		inter1 = dog1
		inter1.bite()

		// 선언 2
		dog2 := Dog{"cherry", 20}
		inter2 := Behavior(dog2)
		inter2.bite()

		// 선언 3
		slice_inter := []Behavior{dog1, dog2}

		for idx, _ := range slice_inter {
			slice_inter[idx].bite()
		}
		for _, val := range slice_inter {
			val.bite()
		}
	*/

	dog := Dog{"foo", 10}
	cat := Cat{"bar", 20}
	act(dog)
	act(cat)

	actAdvanced(dog)
	actAdvanced(cat)

	printValue(dog)
	printValue(cat)
	printValue(true)
	printValue(15)
	printValue("hello world")

}
