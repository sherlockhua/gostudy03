package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type UserInfo struct {
	UserName string `form:"username" json:"username"`
	Passwd   string `form:"passwd" json:"passwd"`
	Age int `form:"age" json:"age"`
	Sex string `form:"sex" json:"sex"`
}

func handleUserInfo(c *gin.Context) {
	var userInfo UserInfo
	err := c.ShouldBind(&userInfo)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

func handleUserInfoJson(c *gin.Context) {
	var userInfo UserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, userInfo)
}


func handleUserInfoQuery(c *gin.Context) {
	var userInfo UserInfo
	err := c.ShouldBind(&userInfo)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

func main() {
	r := gin.Default()

	// /v1/user/login
	// /v1/user/login2gi	
	v1Group := r.Group("/v1")
	v1Group.POST("/user/info", handleUserInfo)
	v1Group.POST("/user/infojson", handleUserInfoJson)
	v1Group.GET("/user/info", handleUserInfoQuery)
	r.Run(":9090")
}
