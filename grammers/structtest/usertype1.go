// Golang은 객체지향 타입을 구조체로 정의 (클래)
package structtest

import "fmt"

type Car struct {
	name  string
	color string
	price int64
	tax   int64
}

func (c Car) Price() int64 {
	return c.price + c.tax
}
func UserType1() {
	hyundai := Car{name: "ionic", price: 100, color: "black", tax: 10}
	toyota := Car{name: "japan", price: 100, color: "white", tax: 10}

	hyundaiPrice := hyundai.Price()
	toyotaPrice := toyota.Price()

	fmt.Println(&hyundai, hyundai, hyundaiPrice)
	fmt.Println(&toyota, toyota, toyotaPrice)

}
