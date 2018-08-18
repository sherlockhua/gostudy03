package main

import (
	"github.com/gin-gonic/gin"
	
)

type Result struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type UserInfo struct {
	Result
	UserName string `json:"username"`
	Passwd   string `json:"passwd"`
}

func handleUserInfo(c *gin.Context) {


	var userInfo = &UserInfo {
		UserName: "skkss",
		Passwd: "SSSS",
	}

	c.XML(200, userInfo)
}



func main() {
	r := gin.Default()
	r.GET("/user/info", handleUserInfo)
	
	r.Run(":9090")
}
