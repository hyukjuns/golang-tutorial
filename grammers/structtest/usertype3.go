package structtest

import "fmt"

type totCost func(int, int) int

func describe(cnt int, price int, fn totCost) {
	fmt.Printf("cnt: %d, price: %d, orderPrice: %d", cnt, price, fn(cnt, price))
}
func UserType3() {

	var orderPrice totCost = func(cnt int, price int) int {
		return cnt*price + 1000
	}
	describe(3, 100, orderPrice)

}
