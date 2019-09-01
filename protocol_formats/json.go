package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Json Marshalling and Unmarshalling")

	type Person struct {
		name string
		age  string
	}

	p1 := Person{"Rajdeep", "24"}

	/*	p1 := new(Person)
		p1.name = "Rajdep"
		p1.age = "24"
	*/
	fmt.Println(p1)

	barr, err := json.Marshal(p1)

	fmt.Println(barr, err)

	fmt.Println(" Now unmarshaling ")
	var p2 Person
	err2 := json.Unmarshal(barr, &p2)
	fmt.Println(p2, err2)
}
