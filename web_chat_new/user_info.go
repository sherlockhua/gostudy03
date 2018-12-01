package main

import (
	"github.com/gostudy03/web_chat_new/common"
	"github.com/gostudy03/web_chat_new/proto"
	"github.com/gostudy03/web_chat_new/dal"
	"github.com/gostudy03/web_chat_new/id_gen"
	"github.com/gorilla/websocket"
	"github.com/gostudy03/xlog"
	"encoding/json"
	"fmt"
)


type UserInfo struct {
	User *common.User
	Conn *websocket.Conn
	RoomInfo *RoomInfo

	WriteChan chan interface{}
}

func NewUserInfo() *UserInfo {
	return &UserInfo{
		WriteChan:make(chan interface{}, 5000),
	}
}

func (u *UserInfo) SendLoop() {

	for msg := range u.WriteChan {
		err := u.Conn.WriteJSON(msg)
		if err != nil {
			u.RoomInfo.DeleteUser(u)
			xlog.LogWarn("send message failed, err:%v data:%v", err, msg)
			continue
		}

		xlog.LogDebug("send message succ, username:%s, data:%s", u.User.Username, msg)
	}
}

func (u *UserInfo) AddMessage(msg interface{}) {
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
		/*
		dataStr := string(data)
		dataStr = fmt.Sprintf("%s：%s\n", u.User.Username, string(data))
		data = []byte(dataStr)
		*/
		xlog.LogDebug("recv message, user_name:%s data:%s", u.User.Username, string(data))
		//把用户发过来的消息广播出去了
		roomInfo := u.RoomInfo
		for _, user := range roomInfo.UserMap {
			var msg proto.Message
			msg.Type = proto.MessageTypeUserTalk
			

			var userTalk proto.UserTalk
			userTalk.UserName = u.User.Username
			userTalk.UserId = u.User.UserId
			userTalk.Content = string(data)
			xlog.LogDebug("recv message, content:%s", userTalk.Content)
			userTalk.ImageUrl = fmt.Sprintf("/static/image/%d.jpg", userTalk.UserId%15+1)

			data, _ := json.Marshal(&userTalk)
			msg.Data = data
			user.AddMessage(&msg)
		}

		var msg common.RoomMessage
		msg.RoomId = u.RoomInfo.Room.RoomId
		msg.Content = string(data)
		msg.UserId = u.User.UserId
		go func(message *common.RoomMessage){
			url := "http://localhost:9090/id/gen"
			id, err := id_gen.GetId(url)
			if err != nil {
				xlog.LogError("call id generator failed, err:%v", err)
				return
			}
			message.MsgId = id
			err = dal.InsertMessage(message)
			if err != nil {
				xlog.LogError("insert message failed, err:%v msg:%#v", err,message)
				return
			}
		}(&msg)
	}
}