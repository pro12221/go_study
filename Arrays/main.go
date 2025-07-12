package main

import "fmt"

func main() {
	// 数组声明和初始化
	var a [5]int
	var b [3]int
	b[0] = 1
	b[1] = 2
	b[2] = 3
	fmt.Println("a:", a) // 输出: a: [0 0 0 0 0]
	fmt.Println("b:", b) // 输出: b: [1	 2 3]

	// 数组遍历
	var my1212 = [...]string{"Hello", "World", "Go", "Language"}
	for i := 0; i < len(my1212); i++ {
		fmt.Printf("my1212[%d]: %s\n", i, my1212[i])
	}
	// k,v 遍历数组
	var my1213 = [...]int{1, 2, 3, 4, 5}
	for k, v := range my1213 {
		fmt.Println(k, v)
	}

}
