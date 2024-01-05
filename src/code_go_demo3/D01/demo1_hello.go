package main // 程序包名 main入口

// 导多个包
import (
	"fmt"
	"time"
)

func main() { // main入口  要求左花括号在同一行！！！
	// go中的表达式建议不加分号
	fmt.Println("hello go")
	time.Sleep(3 * time.Second)
}
