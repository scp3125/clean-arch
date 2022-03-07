package web

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
