package main

import "fmt"

func main() {
	// 声明并初始化切片
	var slice1 []string            // 声明一个空切片
	fmt.Println("slice1:", slice1) // 输出: slice1: []
	fmt.Println(slice1 == nil)     // 输出: true，切片未初始化
	//声明一个整形切片并初始化
	var b = []int{}
	fmt.Println("b:", b) // 输出: b: []
	fmt.Println(b == nil)
	var c = []bool{false, true, false} // 声明并初始化一个布尔切片
	fmt.Println("c:", c)               // 输出: c: [false true false]
	fmt.Println(c == nil)              // 输出: false，切片已初始化

	//切片的循环
	var d = []string{"apple", "banana", "cherry"}
	for i, v := range d {
		fmt.Printf("d[%d]: %s\n", i, v) // 输出: d[0]: apple, d[1]: banana, d[2]: cherry
	}

	//append函数

}
