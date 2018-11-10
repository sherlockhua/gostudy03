package main

import (
	"github.com/gostudy03/day15/tcp_chat/protocal"
)

type RoomInfo struct {
	room *protocal.Room
	//通过user_id获取具体用户的map
	userMap map[uint64]*User
}
