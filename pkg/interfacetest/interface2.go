package interfacetest

import (
	"fmt"
)

func printContents(v interface{}) {
	fmt.Printf("Type: %T, value: %v\n", v, v)
}

func checkType(arg interface{}) {
	switch arg.(type) {
	case bool:
		fmt.Println("Type is boolean", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("Type is int", arg)
	case float64:
		fmt.Println("Type is float", arg)
	case string:
		fmt.Println("Type is string", arg)
	case nil:
		fmt.Println("Type is nil", arg)
	default:
		fmt.Println("No Type Found", arg)
	}
}
func InterfaceTest2() {
	/*
		Type Assertion: 타입 변환
		interfaceValue.(type)
	*/
	var a interface{} = 15
	b := a
	c := a.(int)
	printContents(a)
	printContents(b)
	printContents(c)

	checkType(123)
	checkType(123.123)
	checkType("Hello")
	checkType(nil)
	checkType(true)

}
