package mapob

import "fmt"

func BasicMap() {
	person := make(map[string]string)
	person["name"] = "Redha"
	// person["age"] = "23"

	fmt.Println("Person: ", person)

	mixed := map[string]interface{}{
		"name": "Redha",
		// "age":     24,
		"hobbies": []string{"Code", "Read", "Game"},
		"active":  true,
		"address": map[string]interface{}{
			"country": "Indonesia", "city": "Rohil", "zip": 28983, "mapLagi": map[int]string{1: "satu"}},
	}

	if address, ok := mixed["address"].(map[string]interface{}); ok {
		fmt.Println("address: ", address)

		if mapLagi, ok := address["mapLagi"]; ok {
			fmt.Println("Map lagi: ", mapLagi)
		}
	}

	if _, exists := mixed["age"]; !exists {
		fmt.Println("Age doesn't exists: ")
	}

	delete(person, "name")
	fmt.Println(person)
}
