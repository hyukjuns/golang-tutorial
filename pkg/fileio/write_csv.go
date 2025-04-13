package fileio

import (
	_ "bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func FileTest2() {
	file, err := os.Create("testfile2.csv")
	errCheck1(err)

	defer file.Close()

	wr := csv.NewWriter(file)

	wr.Write([]string{"kim", "10"})
	wr.Write([]string{"kim2", "20"})
	wr.Write([]string{"kim3", "30"})

	wr.Flush()

	fi, err := file.Stat()

	fmt.Printf("file size: %d bytes\n", fi.Size())
	fmt.Println("file name:", fi.Name())
	fmt.Println("file mode:", fi.Mode())
}
