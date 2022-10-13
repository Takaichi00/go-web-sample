package main

import (
	"net/http"
	"strconv"
	fizzbuzz "todo/src/fizz_buzz"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/fizzbuzz/:num", func(c *gin.Context) {
		num, _ := strconv.Atoi(c.Param("num"))
		result := fizzbuzz.Convert(num)
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
