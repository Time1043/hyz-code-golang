package main

import "fmt"

func main() {
	// 整型
	fmt.Println(233 + 11)
	// 字符串
	fmt.Println("周坚深" + "沈以容")
	fmt.Println("周坚深", "沈以容")
	// 布尔
	fmt.Println(1 > 2)
	if 1 > 2 {
		fmt.Println("太阳打西边出来")
	} else {
		fmt.Println("东升西落")
	}
}
