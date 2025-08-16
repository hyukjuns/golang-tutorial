package structtest

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

// receiver, Public Method
func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func StructTest1() {
	kim := Account{number: "123", balance: 1000, interest: 3.1}
	lee := Account{number: "456", balance: 2000}
	park := Account{number: "789", interest: 5.1}
	nam := Account{"789", 4000, 5.1}
	cho := Account{"245-904", 15000000, 0.03}

	fmt.Println(kim, "\n", lee, "\n", park, "\n", nam)
	fmt.Println(int(cho.Calculate()))
}
