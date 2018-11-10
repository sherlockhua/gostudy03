package main

import (
	"net"
	"github.com/gostudy03/day15/tcp_chat/protocal"
)

type User struct {
	UserId uint64
	UserName string
	Conn net.Conn
}

func (u *User)handleEnterRoomResult(code int, room *protocal.Room) (err error) {

	 enterRoomResp := &protocal.UserEnterRoomResp{
		Code: code,
		RoomInfo: room,
	 }

	data, err := protocal.Pack(protocal.UserEnterRoomRespCmd, enterRoomResp)
	if err != nil {
		return
	}

	_, err = u.Conn.Write(data)
	return
}