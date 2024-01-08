package myfunc

import "github.com/gin-gonic/gin"

func Hello1(context *gin.Context) {
	name := "周坚深"
	context.HTML(200, "D01/hello1.html", name)
}

// 结构体
type Student struct {
	Name string
	Age  int
}

func Hello2(context *gin.Context) {
	// 创建结构体对应实例
	s1 := Student{Name: "沈以容", Age: 22}
	context.HTML(200, "D02/hello2.html", s1)
}

// 数组
func Hello3(context *gin.Context) {
	var arr [3]Student
	arr[0] = Student{Name: "chen", Age: 44}
	arr[1] = Student{Name: "liu", Age: 45}
	arr[2] = Student{Name: "lin", Age: 23}
	context.HTML(200, "D03/hello3.html", arr)
}
