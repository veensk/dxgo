package main

import "fmt"

func main() {
	// 1. 打印多个值
	fmt.Println("Hello", "Go", "Language!")

	// 2. 打印数字
	fmt.Println(100)

	// 3. 打印变量
	message := "Learning Go is fun!" // := 是变量声明的简写
	fmt.Println(message)

	// 4. 格式化打印
	name := "张三"
	age := 25
	fmt.Printf("名字：%s, 年龄：%d \n", name, age)
}
