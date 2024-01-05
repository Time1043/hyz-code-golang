package main

import "fmt"

// 声明全局变量 (方式123皆可)
var name string = "zhou"
var sex = "男"
var age int

func main() {
	fmt.Printf("name = %s, sex = %s, age = %d \n", name, sex, age)

	// 定义多个变量：完整
	var x, y int = 5, 6
	fmt.Printf("x = %d, type of x = %T\n", x, x)
	fmt.Printf("y = %d, type of y = %T\n", y, y)

	// 定义多个变量：不指定类型 自动匹配
	var z, who = 10, "doctor"
	fmt.Printf("z = %d, type of z = %T\n", z, z)
	fmt.Printf("who = %s, type of who = %T\n", who, who)

	// 定义多个变量：不赋值 有默认值
	var a, b int
	fmt.Printf("a = %d, type of a = %T\n", a, a)
	fmt.Printf("b = %d, type of b = %T\n", b, b)

	// 定义多个变量：多行括号
	var (
		alpha int  = 100
		flag  bool = true
	)
	fmt.Printf("alpha = %d, type of alpha = %T\n", alpha, alpha)
	fmt.Printf("flag = %t, type of flag = %T\n", flag, flag)
}
