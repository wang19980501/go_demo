package main

import "fmt"

func main() {
	// fmt.Println("go" + "lang")

	// fmt.Println("1 + 1 = ", 1+1)

	// fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!true)

	// var stockCode = 123
	// var endDate = "2020-12-31"
	// var url = "code=%d&endTime=%s"
	// var target_url = fmt.Sprintf(url, stockCode, endDate)
	// fmt.Println(target_url)

	// stockCode := 1235
	// endTime := "12-12"
	// url := "asd"

	// fmt.Print(stockCode, endTime, url)

	// // %d 表示整型数字，%s 表示字符串
	// var stockcode = 123
	// var enddate = "2020-12-31"
	// var url = "Code=%d&endDate=%s"
	// var target_url = fmt.Sprintf(url, stockCode, endDate)
	// fmt.Println(target_url)

	// << 左移符号，左移后变为二进制数
	// a := 1
	// fmt.Print(a << 10)

	//运算符
	var a int = 10
	b := 20
	// c := 30

	fmt.Println("加号 a + b = ", a+b)
	fmt.Println("减法 a - b = ", a-b)
	fmt.Println("乘法 a * b = ", a*b)
	fmt.Println("除法 a / b = ", a/b)
	fmt.Println("求余 a %/ b = ", a%b)
	a++
	fmt.Println("自增 a++ = ", a)
	a--
	fmt.Println("自减 a-- = ", a)
	fmt.Println("变量的实例地址", &a)
	fmt.Println("变量的实例地址", a^b)
}
