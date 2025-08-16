package fileio

import (
	"fmt"
	"os"
)

func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func FileTest() {

	file, err := os.Create("test.txt")
	errCheck1(err)
	defer file.Close()

	s1 := []byte{48, 49, 50, 51, 52}
	n1, err := file.Write([]byte(s1))
	errCheck2(err)
	fmt.Printf("Test 1: %d bytes\n", n1)

	file.Sync()

	s2 := "\ns2-Hello World\n"
	n2, err := file.WriteString(s2)
	errCheck2(err)
	fmt.Printf("Test 2: %d bytes\n", n2)

	file.Sync()

	s3 := "test\n"
	n3, err := file.WriteAt([]byte(s3), 50)
	errCheck1(err)

	fmt.Printf("Test 3: %d bytes\n", n3)

	file.Sync()

	n4, err := file.WriteString("s4-Hello World\n")
	errCheck1(err)

	fmt.Printf("Test 4: %d bytes\n", n4)
	file.Sync()

}
