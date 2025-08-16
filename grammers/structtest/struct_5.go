package structtest

import (
	"fmt"
	"reflect"
)

// label
type Carinfo struct {
	name    string "car name"
	color   string "carcolor"
	company string "carcompany"
}

func StructTest5() {
	tag := reflect.TypeOf(Carinfo{})
	for i := 0; i < tag.NumField(); i++ {
		fmt.Println(tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}
}
