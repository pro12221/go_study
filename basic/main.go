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
	//fmt包
	fmt.Print("Hello, ")           // 不换行
	fmt.Println("World!")          // 自动换行
	fmt.Printf("Number: %d\n", 42) // 格式化输出
	s := fmt.Sprintf("Value: %v", 3.14)
	fmt.Println(s) // 输出格式化字符串\
	name = "Alice"
	age = 30
	info := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(info) // 输出格式化字符串
	fmt.Printf("info的类型是: %T\n", info)

}

// %v	默认格式	fmt.Printf("%v", data)
// %+v	显示结构体字段名	fmt.Printf("%+v", struct)
// %#v	Go语法表示	fmt.Printf("%#v", data)
// %T	类型信息	fmt.Printf("%T", data)
// %d	十进制整数	fmt.Printf("%d", 42)
// %f	浮点数	fmt.Printf("%.2f", 3.1415)
// %s	字符串	fmt.Printf("%s", "text")
// %p	指针地址	fmt.Printf("%p", &x)
