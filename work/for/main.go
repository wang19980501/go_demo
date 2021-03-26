package main

import "fmt"

func main() {
	//第一种方法
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%v*%v=%v\t", j, i, i*j)
	// 	}
	// 	fmt.Println()
	// }

	//第二种方法,
	for i := 1; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > j {
				continue
			}
			fmt.Printf("%v*%v=%v\t", i, j, i*j)
		}
		fmt.Println()
	}
}
