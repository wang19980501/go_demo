package main

import "fmt"

func main() {
	a := make(map[string]string)
	a["2"] = "王家哈"
	a["1"] = "18"
	a["3"] = "aa"

	fmt.Println(a)

	//map查找

	// value, ok := a["名字"]
	// if ok {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Print("没有")
	// }

	//map的遍历

	//map的遍历用for-range

	// for k, v := range a {
	// 	fmt.Println(v)
	// 	fmt.Println(k)
	// }
}
