package main

import "fmt"

func main() {
	// 声明变量的方式一：完整
	var name1 string = "zhou" // 字符串类型
	var age1 int = 33         //数字类型
	fmt.Println("name1 =", name1)
	fmt.Println("age1 =", age1)

	// 声明变量的方式二：不指定类型 自动匹配
	var name2 = "shen"
	var age2 = 22
	fmt.Println("name2 =", name2)
	fmt.Println("age2 =", age2)

	// 声明变量的方式三：不赋值 有默认值
	var name3 string
	var age3 int
	fmt.Println("name3 =", name3)
	fmt.Println("age3 =", age3)

	// 声明变量的方式四：不var
	name4 := "chen"
	age4 := 30
	fmt.Println("name4 =", name4)
	fmt.Println("age4 =", age4)
}
