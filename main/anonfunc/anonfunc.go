package anonfunc

import "fmt"

func AnonFunc() {
	// func as var
	name := func(name string) string {
		return "Hai" + name
	}

	fmt.Println(name("Redha"))

	// pass func as arguments
	add := passFunc(func(x, y int) int { return x + y }) //pass anon func
	fmt.Println(add)

	multiple := func(x, y int) int { return x * y }
	val := passFunc(multiple)
	fmt.Print(val)

	sayHello()("Redha")
}

func passFunc(f func(x, y int) int) int {
	return f(1, 2)
}

// Returning anon func
func sayHello() func(name string) {
	hello := func(name string) {
		fmt.Println("Hello, ", name)
	}
	return hello
}
