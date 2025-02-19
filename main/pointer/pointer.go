package main

import "fmt"

func main() {
	n := 20
	var p *int
	p = &n

	*p = *p * 100
	fmt.Println("address: ", p)
	fmt.Println("Value *p: ", *p)
	fmt.Println("value n: ", n)

	
}