package structtest

import "fmt"

type shopBasket struct{ cnt, price int }

func (b shopBasket) purchase() int {
	return b.cnt * b.price
}

func (b *shopBasket) rePruchaseRef(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func (b shopBasket) rePruchaseVal(cnt, price int) {
	b.cnt += cnt
	b.price += price
	fmt.Println(b.cnt, b.price)
}

func UserType4() {
	// 리시버 전달(값, 참조 형식)
	// 함수는 기본적으로 값 호출 -> 변수의 값이 복사 된 후 내부 전달(원본수정x) -> 맵, 슬라이스 등은 참조전달
	// 리시버(구조체)도 마찬가지로 포인터를 활용해서 메소드 내에서 원본 수정 가능

	// basket 1
	basket1 := shopBasket{3, 5000}
	fmt.Println(basket1.purchase())
	// 참조전달
	basket1.rePruchaseRef(7, 1000)
	fmt.Println(basket1.purchase())

	// basket 2
	basket2 := shopBasket{3, 5000}
	fmt.Println(basket2.purchase())
	// 값 전달
	basket2.rePruchaseVal(7, 1000)
	fmt.Println(basket2.purchase())
}
