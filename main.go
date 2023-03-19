package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/hello", hello)
	r.GET("hello", foo)
	r.Run(":9090")
}

func hello(c *gin.Context) {
	name := c.Query("name")
	user := struct {
		Name string
	}{name}

	c.JSON(http.StatusOK, gin.H{
		"hello": user,
	})
}

func foo(c *gin.Context) {
	var name string
	name = c.Query("name")
	c.IndentedJSON(http.StatusOK, "hello")
	c.IndentedJSON(http.StatusOK, name)
}
