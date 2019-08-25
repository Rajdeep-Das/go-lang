package main

import "fmt"

func main() {

	fmt.Println(" Creating Strcut ")

	type Person struct {
		name  string
		addr  string
		phone string
	}

	p1 := new(Person)
	fmt.Println(p1)

	p1.name = "Raj"
	p1.addr = "Test Addr"
	p1.phone = "123"

	fmt.Println(p1)
}
