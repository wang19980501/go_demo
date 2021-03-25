package main

import "fmt"
import "strings"

func main()  {
	a := "我"
	b := "是"
	c := "王-佳-豪"

	fmt.Println("字符串拼接：", a + b + c)
	fmt.Println("字符串截取：")
	fmt.Println("字符串长度计算：")
	fmt.Println("字符串分割：", strings.Split(c, "-"))

	
}