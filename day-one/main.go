// 必须是 package main 才能独立运行
package main

import "fmt"

// Go 程序从这里开始执行
func main() {
	apple := 5
	banana := 3

	// 调用你自己定义的函数
	total := addFruits(apple, banana)

	fmt.Println(total) // 输出 8
}

// 定义一个普通的函数（函数名通常小写开头）
func addFruits(apple int, banana int) int {
	// 变量声明推荐使用 := 简化
	total := apple + banana
	return total
}
