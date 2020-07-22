package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, I'm Jinmyeong!")
	})
	r.Run()
}

func setupServer() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, I'm Jinmyeong!")
	})
	return r
}
