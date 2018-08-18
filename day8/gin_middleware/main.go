package main

import (
	"net/http"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StatCost(c *gin.Context) {
		start := time.Now()
		fmt.Printf("start stat cost\n")
		c.Next()
		lattancy := time.Since(start)
		fmt.Printf("process request cost:%d ms\n", lattancy/1000/1000)
		
}

func handleUserInfo(c *gin.Context) {
	fmt.Printf("request start process\n")
	time.Sleep(3*time.Second)
	c.JSON(http.StatusOK, "38333k333")
}

func main() {
	r := gin.Default()
	r.Use(StatCost)

	r.GET("/user/info", handleUserInfo)
	r.Run(":8080")
}