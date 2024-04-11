package main

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"
)

func main() {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))

	r.GET("/api/v1/fn-1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Query("item"),
		})
	})

	r.GET("/api/v1/fn-2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Query("item"),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
