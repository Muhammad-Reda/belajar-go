package swt

import "fmt"

func SwitchType() {
	var day interface{} = 2.4

	switch v := day.(type) {
	case int:
		switch v {
		case 1:
			fmt.Println("int Monday")
		}
	case string:
		switch v {
		case "1":
			fmt.Println("string Monday")
		}
	default:
		fmt.Println("Neiter int or string")
	}
}
