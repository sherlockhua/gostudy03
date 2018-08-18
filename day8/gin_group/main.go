package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	username := c.PostForm("username")
	passwd := c.PostForm("passwd")

	var result UserInfo = UserInfo{
		UserName: username,
		Passwd:   passwd,
	}

	result.Code = 0
	result.Message = "success"

	c.JSON(http.StatusOK, result)
}


func main() {
	r := gin.Default()

	// /v1/user/login
	// /v1/user/login2
	v1Group := r.Group("/v1")
	v1Group.GET("/user/login", handleUserInfo)
	v1Group.GET("/user/login2", handleUserInfo)

	r.Run(":9090")
}
