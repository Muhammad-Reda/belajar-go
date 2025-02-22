package variadic

import (
	"fmt"
	"slices"
	"strings"
)

func Variadic(nums ...string) string {
	fmt.Printf("%v\n", nums)
	slices.Reverse(nums)
	max := slices.Max(nums)
	fmt.Printf("Max: %s\n", max)
	return strings.Join(nums, "")
}
