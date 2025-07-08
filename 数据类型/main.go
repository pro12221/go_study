package main

import (
	"fmt"
	"reflect"
)

func main() {
	//数字定义
	var a int8 = 4
	var b int32 = 4
	var c int64 = 4
	d := 4
	fmt.Printf("a: %T %v \n", a, a)
	fmt.Printf("b: %T %v \n", b, b)
	fmt.Printf("c: %T %v \n", c, c)
	fmt.Printf("d: %T %v \n", d, d)
	//查看数据类型
	fmt.Println("a的类型是:", reflect.TypeOf(a))

	// 布尔类型
	var b1 bool = true
	var b2 bool
	fmt.Printf("b1: %T \n", b1)
	fmt.Printf("b2: %T %v \n", b2, b2)

	// 字符串类型
	var str1 string = "Hello, World!"
	s2 := `
	这是一个多行字符串
	可以包含换行符和制表符
	`
	fmt.Printf("str1: %T %s \n", str1, str1)
	fmt.Printf("s2: %T %s \n", s2, s2)
	// byte和rune
	var b3 byte = 'A' // byte是uint8的别名
	var r rune = '中'  // rune是int32的别名
	fmt.Printf("b3: %T %c \n", b3, b3)
	fmt.Printf("r: %T %c \n", r, r)

}
