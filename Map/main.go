package main

import (
	"fmt"
)

func main() {
	//map 是一个引用类型（Reference Type），和切片一样，它在使用前必须先进行初始化。
	var mymap map[string]int
	fmt.Println(mymap == nil) // 输出：true

	// 初始化 map
	mymap = make(map[string]int)     //创建一个空 map
	mymap = make(map[string]int, 10) //创建一个空 map，预分配了10个元素的空间

	fmt.Println(mymap == nil) // 输出：false

}
