package main

import "fmt"

/**
	//数组的定义
	//定义一个长度为3元素类型为int的数组
	var a [3]int

	//定义一个有初始值的数组
	var a [3]int{1,2,3}
	var b [3]string{"北京", "上海", "广州"}

	//定义一个长度根据初始值来确定的数组
	var c[...]int{1,2,3,4}
**/

func main() {
	//遍历数组
	//第一种
	// a := [...]string{"北京", "上海", "广州"}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// //第二种
	// for index, vaule := range a {
	// 	fmt.Println(index, vaule)
	// }

	//联系题 1、求数组[1,3,5,7,8]所有元素的和

	// a := [5]int{1, 3, 5, 7, 8}
	// sum := 0
	// for index, value := range a {
	// 	sum = value + sum
	// 	fmt.Printf("index=%v,len(a)=%v,sum=%v,value=%v", index, len(a)-1, sum, value)
	// 	fmt.Println()
	// }
	// fmt.Println("sum=", sum)

	//2、找出数组中和为指定的两个元素下标，比如从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)

	b := []int{1, 2, 4, 5, 6}
	for index, value := range b {
		for index1, value1 := range b {
			c := value + value1
			if c == 10 {
				fmt.Printf("index1=%v,index2=%v\n", index, index1)
			}
		}
	}
}
