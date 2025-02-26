package structtest

import "fmt"

func StructTest4() {
	// 구조체 익명선언
	car1 := struct{ name, color string }{"toyota", "white"}
	fmt.Println(car1)
	fmt.Printf("%#v\n", car1)

	// 배열
	cars := []struct{ name, color string }{
		{"test", "red"},
		{"test2", "black"},
		{"test3", "white"},
	}

	for _, c := range cars {
		fmt.Printf("%s, %s, --- %#v\n", c.name, c.color, c)
	}
}
