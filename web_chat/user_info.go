package main

import (
	"github.com/gostudy03/web_chat/common"
	"github.com/gorilla/websocket"
	"github.com/gostudy03/xlog"
	"fmt"
)


type UserInfo struct {
	User *common.User
	Conn *websocket.Conn
	RoomInfo *RoomInfo

	WriteChan chan *common.Message
}

func NewUserInfo() *UserInfo {
	return &UserInfo{
		WriteChan:make(chan *common.Message, 5000),
	}
}

func (u *UserInfo) SendLoop() {

	for msg := range u.WriteChan {
		err := u.Conn.WriteMessage(websocket.TextMessage, msg.Data)
		if err != nil {
			u.RoomInfo.DeleteUser(u)
			xlog.LogWarn("send message failed, err:%v data:%v", err, string(msg.Data))
			continue
		}

		xlog.LogDebug("send message succ, username:%s, data:%s", u.User.Username,string(msg.Data))
	}
}

func (u *UserInfo) AddMessage(msg *common.Message) {
	select {
	case u.WriteChan <- msg:
	default:
		xlog.LogError("user chan is full")
		return
	}
}

func (u *UserInfo)ReadLoop() {
	defer u.Conn.Close()
	for {
		msgType, data, err := u.Conn.ReadMessage()
		if err != nil {
			u.RoomInfo.DeleteUser(u)
			return
		}

		if msgType != websocket.TextMessage {
			xlog.LogWarn("recv message not text, data:%v", string(data))
			continue
		}

		dataStr := string(data)
		dataStr = fmt.Sprintf("%s：%s\n", u.User.Username, string(data))
		data = []byte(dataStr)
		xlog.LogDebug("recv message, user_name:%s data:%s", u.User.Username, string(data))
		//把用户发过来的消息广播出去了
		roomInfo := u.RoomInfo
		for _, user := range roomInfo.UserMap {
			var msg common.Message
			
			msg.Data = data
			user.AddMessage(&msg)
		}
	}
}