package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"github.com/gostudy03/web_chat/dal"
)

func indexView(ctx *gin.Context) {
	logined := IsLogin(ctx)
	if !logined {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		return
	}

	//roomList, err := dal.GetAllRoomList()
	roomList, err := roomMgr.GetRoomList()
	if err != nil {
		ctx.Redirect(http.StatusMovedPermanently, "/index")
		return
	}

	ctx.HTML(http.StatusOK, "./views/index.html", roomList)
}
