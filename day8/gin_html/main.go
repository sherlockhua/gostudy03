package main

import (
	"github.com/gin-gonic/gin"

)


func handleHtml(c *gin.Context) {
	
	c.HTML(200, "post", "ksdkfskfkskfsdkfs")
}
	

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	r.GET("/user/info", handleHtml)
	
	r.Run(":9090")
}
