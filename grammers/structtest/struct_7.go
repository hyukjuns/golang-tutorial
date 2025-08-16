package structtest

import "fmt"

type Account3 struct {
	number   string
	balance  float64
	interest float64
}

// 포인터형으로 값 복사
func NewAccount(number string, balance float64, interest float64) *Account3 {
	return &Account3{number, balance, interest}
}

// 값 복사
func (a Account3) CalculateVal(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

// 주소 전달
func (a *Account3) CalculateRef(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func StructTest7() {
	kim := Account3{number: "245-901", balance: 10000000, interest: 0.01}
	lee := Account3{number: "123-901", balance: 10000000, interest: 0.01}

	// var lee *Account3 = new(Account3)
	// lee.number = "123-456"
	// lee.balance = 13000000
	// lee.interest = 0.025

	fmt.Println(kim, lee)

	park := NewAccount("123-123", 10000000, 0.01)
	fmt.Println(park)

	lee.CalculateVal(1000000)
	kim.CalculateRef(1000000)
	park.CalculateRef(1000000)

	fmt.Println(int(lee.balance), int(kim.balance), int(park.balance))
}
