package deferfunc

import "fmt"

func add(a, b int) {
	fmt.Println(a + b)
}

func showAdd(f func(a, b int)) func(a, b int) {
	return f
}

func DeferFunc() {
	defer fmt.Println("End")
	fmt.Println("Start")

	val := showAdd(func(a, b int) { fmt.Println(a + b) })
	val(1, 2)

	add(2, 3)
}
