package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/", func(c *gin.Context) {
		_, ok := c.GetPostForm("body")
		if !ok {
			c.String(http.StatusBadRequest, "badrequest")
		} else {
			c.Redirect(http.StatusOK, "/todo_list")
		}
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
