package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("program started")

	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"data":"Hello world",
		})
	})

	r.Run()

	fmt.Println("finished program")
}
