package main

import (
	"fmt"

	"github.com/wang19980501/go_demo/work/slice/utils"
)

/**
	//slice 切片
	定义：name表示变量名，T表示切片中的元素类型
	var name []T
**/
func main() {
	// var a = []int{}
	// fmt.Println(len(a))

	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
		fmt.Println(a)
	}
	fmt.Println(a)
	utils.Cal

}
