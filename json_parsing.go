package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func main() {
	var p Person
	json.NewDecoder(strings.NewReader(`{"first_name": "John", "last_name": "Doe"}`)).Decode(&p)
	fmt.Println(p.FirstName, p.LastName)
}
