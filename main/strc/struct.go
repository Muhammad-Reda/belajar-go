package strc

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	Address
}

type Address struct {
	country string
	city    string
	zip     int
}

func ShowStruct() {
	redha := Person{name: "Redha", age: 22, Address: Address{"Indonesia", "Rokan Hilir", 28983}}

	fmt.Printf("Name: %s\n", redha.name)
	fmt.Printf("Age: %d\n", redha.age)
	fmt.Printf("Country: %s\n", redha.country)
	fmt.Printf("City: %s\n", redha.city)
	fmt.Printf("ZIP: %d\n", redha.zip)
}
