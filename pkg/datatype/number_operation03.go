package main

import (
	"math"
)

func main() {
	// overflows
	var n1 uint8 = math.MaxUint8 + 1
	var n2 uint8 = math.MaxUint16 + 1
	var n3 uint8 = math.MaxUint32 + 1
	var n4 uint8 = math.MaxUint64 + 1
	var n5 uint8 = -1
}
