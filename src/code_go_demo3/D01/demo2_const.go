package main

import "fmt"

// const定义枚举
const ( // const枚举关键字iota 每行iota累加1 第一行iota默认值0
	BEIJING  = 10 * iota // iota=0
	SHANGHAI             // iota=1
	SHENZHEN             // iota=2
)

const (
	a, b = iota + 1, iota + 2 // iota=0  a=1  b=2  ——  iota + 1, iota + 2
	c, d                      //iota=1  c=2  d=3
	e, f                      // iota=2  e=3  f=4
	g, h = iota * 2, iota * 3 // iota=3  g=6  h=9  ——  iota * 2, iota * 3
	i, j                      // iota=4  i=8  j=12
)

func main() {
	// 常量 只读属性--------------------------------------------------
	const length int = 10
	// length = 100 // 常量不允许修改

	// 枚举--------------------------------------------------
	fmt.Println("BEIJING =", BEIJING)
	fmt.Println("SHANGHAI =", SHANGHAI)
	fmt.Println("SHENZHEN =", SHENZHEN)

	fmt.Println("a =", a, "b =", b)
	fmt.Println("c =", c, "d =", d)
	fmt.Println("e =", e, "f =", f)

	fmt.Println("g =", g, "h =", h)
	fmt.Println("i =", i, "j =", j)

	// iota只能配合const()一起使用  iota只有在const进行累加效果
	// var a int = iota
}
