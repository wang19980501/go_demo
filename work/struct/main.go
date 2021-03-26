package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	var cat Cat
	cat.Name = "小白"
	cat.Age = 12
	cat.Color = "红色"
	fmt.Println(cat)

}
