package main

import (
	
	"github.com/gin-gonic/gin"
	"github.com/gostudy03/xlog"
	
	"net/http"

)

func roomEnterView(ctx *gin.Context) {
	xlog.LogDebug("room enter view00000000000000000000000000000000")
	logined := IsLogin(ctx)
	if !logined {
		userId := GetUserId(ctx)
		xlog.LogDebug("room enter view, not login, user_id:%d", userId)
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		return
	}
	
	userId := GetUserId(ctx)
	xlog.LogDebug("room enter view, user_id:%d", userId)
	roomId, ok := ctx.GetQuery("room_id")
	if !ok {
		ctx.Redirect(http.StatusMovedPermanently, "/index")
		return
	}

	ctx.HTML(http.StatusOK, "./views/room.html", roomId)
}
