package main

import (
	//"fmt"
	//"github.com/gostudy03/web_chat/id_gen"
	"github.com/gostudy03/web_chat/dal"
	//"github.com/gostudy03/web_chat/common"
	"github.com/gostudy03/xlog"
	//"crypto/md5"
	"github.com/gin-gonic/gin"
)

const (
	UserPasswdSalt = "222964304830267393"
)

func main() {
	dns := "root:123456@tcp(192.168.20.200:3306)/golang?parseTime=True"
	err := dal.InitDb(dns)
	if err != nil {
		xlog.LogError("init db failed, err:%v", err)
		return
	}
	r := gin.Default()
	r.Use(UserMiddleware)
	r.LoadHTMLGlob("./views/*")
	r.Static("/static/", "./static/")
	r.GET("/user/login", loginView)
	r.POST("/user/register", registerHandle)
	r.POST("user/login", loginHandle)
	r.GET("/index", indexView)
	r.Run(":8888")

	/*
		url := "http://localhost:9090/id/gen"
		id, err := id_gen.GetId(url)
		if err != nil {
			fmt.Printf("err:%v\n", err)
			return
		}

		fmt.Printf("call id server succ, id:%v\n", id)
		var userInfo common.User
		userInfo.Nickname = "五连败"
		userInfo.UserId = int64(id)
		userInfo.Username = "carry01"
		userInfo.Passwd = fmt.Sprintf("123456.%s", UserPasswdSalt)

		result := md5.Sum([]byte(userInfo.Passwd))
		userInfo.Passwd = fmt.Sprintf("%x", string(result[:]))
		err = dal.Register(&userInfo)
		if err != nil {
			xlog.LogError("register failed, err:%v", err)
			return
		}
	*/
}
