package main

import (
	"gin231229/myfunc"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/s", "static") // 相对路径 文件夹名称 (映射关系)
	//r.StaticFS("/s",http.Dir("static"))
	r.GET("/hello2", myfunc.Hello2)
	r.Run()
}
