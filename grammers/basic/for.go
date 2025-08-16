package basic

import "fmt"

func For01() {
	for i := 0; i < 5; i++ {
		fmt.Println("ex1: ", i)
	}

	// error case 1
	/*
		for i := 0; i<5; i++
		{
			fmt.Println("ex1: ", i)
		}
	*/

	// error case 2
	/*
		for i := 0; i < 5; i ++
			fmt.Println("ex1: ", i)
	*/

	// infinite
	/*
		for {
			fmt.Println("hello,")
			fmt.Println("world!")
		}
	*/

	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println(index, name)
	}
	for _, name := range loc {
		fmt.Println(name)
	}

}

func For02() {

	// case 1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i

	}
	fmt.Println("ex1: ", sum1)

	// case 2
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
	}
	fmt.Println("ex2: ", sum2)

	// case 3
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3: ", sum3)

	// case 4
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4: ", i, j)
	}

	// error case
	/*
		for i, j := 0, 0; i <= 10; i, j = i++, j+10 {
			fmt.Println("ex4: ", i, j)
		}
	*/

}

func For03() {

here:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break here
			}
			fmt.Println("ex1: ", i, j)
		}
	}

loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue loop2
			}
			fmt.Println("ex2: ", i, j)
		}
	}

here2:
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue here2
		}
		fmt.Println("ex3: ", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex3: ", i)
	}
}
