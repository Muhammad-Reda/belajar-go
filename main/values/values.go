package main

import (
	"fmt"
	// "runtime"
	// "time"
)

// func checkOs() {
// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("OS X.")
// 	case "linux":
// 		fmt.Println("OS Linux.")
// 	default:
// 		fmt.Printf("%s \n", os)
// 	}
// }

func main() {
	// var firstName string = "Muhammad"
	// var lastName string = "redha"

	// number := 20
	// var dec float32 = 25.3

	// fmt.Println(firstName + lastName)
	// fmt.Println(float32(number) + dec)
	// fmt.Println(!false)
	// fmt.Println(add(1, 2))

	// x, y := add(2, 4)
	// fmt.Println(x)
	// fmt.Println(y)

	// switch val := 1; val {
	// case 2:
	// 	fmt.Println(true)
	// default:
	// 	fmt.Println(false)
	// }

	// checkOs()
	// fmt.Println(time.Now().Weekday() + 1)

	// switch t := time.Now().Hour(); {
	// case t < 12:
	// 	fmt.Println("Good morning")
	// case t < 17:
	// 	fmt.Println("Good afternoon")
	// default:
	// 	fmt.Println("Good evening")
	// }

	// fmt.Println(time.Now())

	// {
	// 	defer fmt.Println("world")
	// 	a := 2
	// 	n := 3

	// 	fmt.Println(add(a, n))
	// 	// fmt.Println("Hello")
	// }

	{
		defer fmt.Println("done")
		fmt.Println("counting")

		for i := 0; i < 10; i++ {
			defer fmt.Println(i)
		}

	}

}

// func add(x, y int) (string, int) {
// 	return "Result: ", (x + y)
// }
