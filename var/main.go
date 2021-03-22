package main

import "fmt"

//声明变量
// var name string
// var age int
// var isOk bool

//批量声明
var (
	name string //""
	age  int    //0
	isOk bool   //false
)

// func main() {
// 	name = "wang"
// 	age = 16
// 	isOk = true
// 	fmt.Println(isOk)
// 	fmt.Printf("name:%s", name) //%s:占位符，使name的值去替换占位符
// 	fmt.Println(age)

// 	s1 := "wang"
// 	fmt.Println(s1)

// 	//函数外的每个语句都必须以关键字开始（var，const，func等）
// 	//:= 不能使用在函数外
// 	//_ 多用于占位符，表示忽略值

// }

func to() (int, string)  {
	return 10, "name"
	
}

func main()  {
	x, _ := to()
	_, y := to()
	fmt.Println(x)
	fmt.Println(y)
	
}
