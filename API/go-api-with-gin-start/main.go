package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	fmt.Println("got req for home");
	c.JSON(200, gin.H{
		"message": "this is home",
	})

}

func main() {
	router := gin.Default()
	router.GET("/", home)
	router.Run("localhost:8080")
}
