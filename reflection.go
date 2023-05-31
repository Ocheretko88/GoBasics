package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	FirstName string `json:"first_name"`
	Age       int    `json:"age"`
}

func inspectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t.Name())
	fmt.Println(t.Kind())
}

func main() {
	p := Human{FirstName: "John", Age: 30}
	inspectType(p)
}
