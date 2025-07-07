package main

import "fmt"

func main() {
	// 声明并初始化变量
	var name string = "Alice"
	var age int = 30
	var isStudent bool = true

	fmt.Println("Name:", name, "Age:", age, "Is Student:", isStudent)

	// 类型推断 (短声明语法)： := 只能用于声明新变量
	city := "New York"
	gpa := 3.8
	fmt.Println("City:", city, "GPA:", gpa)

	// 批量声明
	var (
		a int = 1
		b int = 2
	)
	fmt.Println("a:", a, "b:", b)

	// 重新赋值 (不需要 :=)
	name = "Bob"
	fmt.Println("New Name:", name)
}
