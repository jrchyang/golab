// 语法/main.go
package main

import "fmt"

// 演示基本语法结构
func main() {
	// 变量声明和初始化
	var a int = 10
	b := "hello" // 类型推断
	fmt.Printf("变量: a=%d, b=%s\n", a, b)

	// 条件语句
	if a > 5 {
		fmt.Println("a > 5")
	}

	// 循环（类似 while）
	i := 0
	for i < 3 {
		fmt.Println("循环次数：", i)
		i++
	}

	// 调用函数
	result := add(3, 4)
	fmt.Println("3 + 4 =", result)

	// 切片操作
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println("切片：", slice)
}

// 函数定义
func add(x, y int) int {
	return x + y
}
