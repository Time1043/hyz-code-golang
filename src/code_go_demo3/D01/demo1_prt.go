package main

import "fmt"

func main() {
	// 内置函数
	print("山桃红花满上头， \n") // 没有间隔
	println("蜀江春水拍山流。")  // 尾部有换行

	// fmt包
	fmt.Print("花红易衰似郎意，\n")
	fmt.Println("水流无限似侬愁。")

	// 可传多个字符串  打印输出中间自带空格
	fmt.Println("离九霄而膺天命", "情何以堪", "御四海而哀苍生", "心为之伤")

	// fmt格式化  占位符 %s(字符串) %d(整数) %f(小数) -> 源码doc.go
	fmt.Printf("有些事%s没%d两重，上了秤%.2f斤打不住。", "不上秤", 4, 1000.119)
	// 不要当成占位符%
	fmt.Printf("%s，得到的最终数据是：97%%s", "周坚深")
}
