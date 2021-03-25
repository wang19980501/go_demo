package main

import "fmt"

//常量
//定义了常量之后不能修改 
const pi = 3.1415926


const (
	a1 = iota
	a2
	a3
)

func main()  {
	// pi = 123
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)


} 
