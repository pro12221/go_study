package main

import (
	"fmt"
	"reflect"
	"strings"
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
	var b3 byte = 'A'                  // byte是uint8的别名
	var r rune = '中'                   // rune是int32的别名
	fmt.Printf("b3: %T %c \n", b3, b3) // %c格式化输出字符
	fmt.Printf("r: %T %d \n", r, r)    // %d格式化输出整数值
	s := "Hello, 世界"
	ss_rune_slice := []rune(s) // 将字符串转换为rune切片
	fmt.Printf("字符串的rune切片: %c \n", ss_rune_slice)

	// 字符串方法

	var my111 = "Hello, 世界"
	fmt.Printf("字符串长度: %d \n", len(my111))

	var my222 = "Hello"
	var my333 = "时间"
	fmt.Printf("字符串拼接: %s,%s \n", my222, my333)
	fmt.Println("字符串拼接:", my222+","+my333)

	var my444 = "123-456-789"
	var arr = strings.Split(my444, "-") // 使用Split函数将字符串分割成切片
	fmt.Printf("分割后的字符串切片: %v \n", arr)

	var my555 = "555-666-777"
	var arr1 = strings.Join(strings.Split(my555, "-"), "#") // 先分割成切片再用Join连接
	fmt.Printf("连接后的字符串: %s \n", arr1)

	// 字符串遍历
	fmt.Println("--- 按字节遍历 (ASCII 字符串) ---")
	var my666 = "ascsac"
	for i := 0; i < len(my666); i++ {
		fmt.Printf("字节 %d: %c \n", i, my666[i])
	}
	fmt.Println("--- 按字符 (rune) 遍历 ---")
	var my777 = "Hello, 世界"
	for i, r := range my777 {
		fmt.Printf("字符 %d: %c \n", i, r)
	}
}
