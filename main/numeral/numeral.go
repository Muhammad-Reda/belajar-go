package main

import "fmt"

func main() {
	hello := "Hello"

	for i := 0; i < len(hello); i++ {
		fmt.Println(hello[i])
		fmt.Printf("%b", hello[i])
	}

	n := 2548
	fmt.Printf("%q", n)
}
