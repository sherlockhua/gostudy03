package main

import (
	"github.com/gostudy03/day15/tcp_chat/protocal"
	"github.com/gostudy03/xlog"
	"net"
)

type Message struct {
	CmdNo uint16
	Body  interface{}
}

type User struct {
	UserId   uint64
	UserName string
	Conn     net.Conn
	OutBox   chan *Message
}

func NewUser(userId uint64, UserName string, conn net.Conn) (user *User) {
	user = &User{
		UserId:   userId,
		UserName: UserName,
		Conn:     conn,
		OutBox:   make(chan *Message, 1024),
	}

	go user.sendMessage()
	return
}

func (u *User) handleEnterRoomResult(code int, room *protocal.Room) (err error) {

	enterRoomResp := &protocal.UserEnterRoomResp{
		Code:     code,
		RoomInfo: room,
	}

	data, err := protocal.Pack(protocal.UserEnterRoomRespCmd, enterRoomResp)
	if err != nil {
		return
	}

	_, err = u.Conn.Write(data)
	return
}


func (u *User) handleSendTextResult(code int, room *protocal.Room) (err error) {
	
		recvText := &protocal.UserRecvText{
			Code:     code,
		}
	
		data, err := protocal.Pack(protocal.UserRecvTextCmd, recvText)
		if err != nil {
			return
		}
	
		_, err = u.Conn.Write(data)
		return
	}

func (u *User) AppendMessage(msg *Message) {
	select {
	case u.OutBox <- msg:
	default:
		xlog.LogError("user message chan full")
		return
	}
}

func (u *User) sendMessage() {
	for msg := range u.OutBox {
		data, err := protocal.Pack(msg.CmdNo, msg.Body)
		if err != nil {
			xlog.LogError("pack message:%#v failed, err:%v", msg, err)
		}

		_, err = u.Conn.Write(data)
		if err != nil {
			xlog.LogError("send to client failed, err:%v", err)
			continue
		}
	}
}
