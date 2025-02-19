package main

import (
	"fmt"

	"example.com/depedency/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println(morestrings.ReverseRunes("Hello world"))
	fmt.Println(cmp.Diff("Hello world", "Hello go"))
}
