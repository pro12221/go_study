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

	// 常量声明
	const PI float64 = 3.14159
	const Greeting = "Hello" // 类型推断

	fmt.Println("PI:", PI, "Greeting:", Greeting)

	// iota：用于创建连续的枚举值
	const (
		_ = iota // _ 用于跳过第一个值 (通常是 0)
		A
		B
		C
	)
	fmt.Println("A:", A, "B:", B, "C:", C) // 输出 A: 1, B: 2, C: 3
}
