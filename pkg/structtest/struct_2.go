package structtest

import "fmt"

type Account2 struct {
	number   string
	balance  float64
	interest float64
}

func (a Account2) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func StructTest2() {
	// pointer
	var kim *Account2 = new(Account2)
	kim.number = "123"
	kim.balance = 10000000
	kim.interest = 0.015

	// 축약
	hong := &Account2{number: "12345", balance: 15000000, interest: 0.04}

	// 축약
	lee := new(Account2)
	lee.number = "456"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println(kim, hong, lee)
	fmt.Printf("%#v, %#v, %#v", kim, hong, lee)

	fmt.Println()
	fmt.Println(int(kim.Calculate()))
}
