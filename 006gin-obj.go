package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   uint64
	Name string
}

func main() {
	//var users = []User{{ID: 1, Name: "张三"}, {ID: 2, Name: "李四"}}
	tt := map[string]int{"a": 1, "b": 2}
	var r = gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, tt)
	})
	r.Run(":8081")
}
