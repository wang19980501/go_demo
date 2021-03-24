package main

import "fmt"

func main() {
	// a := 1
	// b := 2
	// if a == b {
	// 	fmt.Println(true)
	// } else if a == 1 && b == 2 {
	// 	fmt.Println(false)
	// }
	c := 80

	switch {
	case 80 <= c < 90:
		fmt.Println("优秀")
	}
}
