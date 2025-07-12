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
	var e = []int{1, 2, 3}
	e = append(e, 4, 5)  // 向切片添加元素
	fmt.Println("e:", e) // 输出: e: [1 2 3 4 5]

	//追加单个元素
	s1 := []int{1, 2, 3}
	fmt.Printf("s1 (before append): %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))

	s1 = append(s1, 4) // 将 4 追加到 s1
	fmt.Printf("s1 (after append 4): %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))

	//追加多个元素
	// 2. 追加多个元素
	s2 := []string{"apple", "banana"}
	fmt.Printf("s2 (before append): %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s2 = append(s2, "orange", "grape") // 追加 "orange" 和 "grape"
	fmt.Printf("s2 (after append multiple): %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	// 3. 追加另一个切片的所有元素 (使用 ... 展开操作符)
	s3 := []float64{1.1, 2.2}
	toAdd := []float64{3.3, 4.4, 5.5}
	fmt.Printf("s3 (before append): %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))

	s3 = append(s3, toAdd...) // 注意这里的 ...，它会把 toAdd 切片展开成独立的元素
	fmt.Printf("s3 (after append another slice): %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))

	//切片合并
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	// 使用 append 函数合并两个切片
	merged := append(arr1, arr2...) // 注意这里的 ...，它会把 arr2 切片展开成独立的元素
	fmt.Println("Merged slice:", merged)

	//切片删除元素方法
	//s = append(s[:i], s[i+1:]...)

}
