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

	username := c.Query("username")
	passwd := c.DefaultQuery("passwd", "dkdkdkdkdkdkdkd")

	var result UserInfo = UserInfo{
		UserName: username,
		Passwd:   passwd,
	}

	result.Code = 0
	result.Message = "success"

	c.JSON(http.StatusOK, result)
}

func handleUserParams(c *gin.Context) {

	username := c.Param("username")
	passwd := c.Param("passwd")

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
	r.GET("/user/info", handleUserInfo)
	r.GET("/user/info/:username/:passwd", handleUserParams)

	r.Run(":9090")
}
