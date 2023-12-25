package main

import (
	"code_go_demo1/api"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	Add()        // 包名相同直接调用
	api.BaiDu()  // 先导包后使用
	api.Google() // 先导包后使用
	// 函数需要首字母大写 才能被外面引用
	// 小写则认为只有在自己包内才能访问
}
