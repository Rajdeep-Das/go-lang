package main

import "fmt"

func main() {

	fmt.Println("Creating Map")
	idMap := make(map[string]int)

	idMap["Joe"] = 5
	idMap["Raj"] = 1

	id, p := idMap["Joe"]
	fmt.Println("Is 'Joe' Key is Present: ", p)
	fmt.Println("Length  of Map : ", len(idMap), " ", "Value:", id)

	fmt.Println("Iterating Throug Map")
	for key, val := range idMap {
		fmt.Println(key, val)
	}
}
