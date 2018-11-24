package main

import (
	"github.com/gostudy03/web_chat/dal"
	"strconv"
	"github.com/gin-gonic/gin"
	
	"github.com/gorilla/websocket"
	"github.com/gostudy03/xlog"
	
	"net/http"
)


var upgrader = websocket.Upgrader{} // use default options

func wsHandle(ctx *gin.Context) {
	logined := IsLogin(ctx)
	if !logined {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		return
	}

	userId := GetUserId(ctx)
	if userId <= 0 {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		return
	}
	
	user, err := dal.GetUserInfoById(userId)
	if err != nil {
		ctx.Redirect(http.StatusMovedPermanently, "/user/login")
		return
	}

	roomIdStr, ok := ctx.GetQuery("room_id")
	if !ok {
		ctx.Redirect(http.StatusMovedPermanently, "/index")
		return
	}

	roomId, err := strconv.ParseInt(roomIdStr, 10, 64)
	if err != nil {
		xlog.LogError("convert room_id:%s failed, err:%v", roomIdStr, err)
		ctx.Redirect(http.StatusMovedPermanently, "/index")
		return
	}

	xlog.LogDebug("user id:%d enter room:%d", userId, roomId)
	roomInfo, err := roomMgr.GetRoom(roomId)
	if err != nil {
		ctx.Redirect(http.StatusMovedPermanently, "/index")
		return
	}

	xlog.LogDebug("room info:%#v", roomInfo)
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		xlog.LogError("upgrade to websocket failed, err:%v", err)
		return
	}

	userInfo := NewUserInfo()
	userInfo.User = user
	userInfo.Conn = c
	userInfo.RoomInfo = roomInfo
	isAlreadLogin := roomInfo.AddUser(userInfo)
	if isAlreadLogin {
		xlog.LogFatal("user is already login, user_id:%d", userInfo.User.UserId)
		ctx.Redirect(http.StatusMovedPermanently, "/index")
		return
	}

	go userInfo.ReadLoop()
	go userInfo.SendLoop()
/*
	defer c.Close()
	for {
		
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
		
		c.WriteMessage(websocket.TextMessage, []byte("this is server push"))
		time.Sleep(time.Second)
		
	}
	*/
}