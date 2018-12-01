package main

import (
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/gostudy03/xlog"
	"github.com/gostudy03/web_chat_new/dal"
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
	roomIdStr, ok := ctx.GetQuery("room_id")
	if !ok {
		ctx.Redirect(http.StatusMovedPermanently, "/index")
		return
	}

	roomId, err := strconv.ParseInt(roomIdStr, 10, 64)
	if err != nil {
		ctx.Redirect(http.StatusMovedPermanently, "/index")
		return
	}
	roomInfo, err := dal.GetRoomInfoById(roomId)
	if err != nil {
		ctx.Redirect(http.StatusMovedPermanently, "/index")
		return
	}

	xlog.LogDebug("room_info:%#v", roomInfo)
	ctx.HTML(http.StatusOK, "./views/room.html", gin.H{
		"RoomId":roomInfo.RoomId,
		"RoomName":roomInfo.RoomName,
	})
}
