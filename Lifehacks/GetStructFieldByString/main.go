package main

import (
	"fmt"
	"reflect"
)

type Steps struct {
	Step string
	Name string
	Date string
}

func main() {
	e := Steps{}
	field := "Step"
	newStep := "Name"
	mutable := reflect.ValueOf(&e).Elem()
	mutable.FieldByName(field).SetString(newStep)
	fmt.Println(e.Step)
}
