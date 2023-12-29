package myfunc

import "github.com/gin-gonic/gin"

func Hello1(context *gin.Context) {
	name := "周坚深"
	context.HTML(200, "D01/hello1.html", name)
}

type Student struct {
	Name string
	Age  int
}

func Hello2(context *gin.Context) {
	// 创建结构体对应实例
	s1 := Student{Name: "沈以容", Age: 22}
	context.HTML(200, "D02/hello2.html", s1)
}
