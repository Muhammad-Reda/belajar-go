package mapob

import "fmt"

func MapCache() {
	cache := make(map[string]int)

	calculatedCache := func(key string) int {

		if result, exists := cache[key]; exists {
			fmt.Printf("Key exists %d", result)
			return result
		}

		result := len(key) * 10
		cache[key] = result
		return result
	}

	calculatedCache("Hello")
	calculatedCache("world")
	calculatedCache("Hello")

}
