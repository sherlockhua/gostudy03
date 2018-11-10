package main

import (
	"time"
	"github.com/gostudy03/day15/tcp_chat/protocal"
	"fmt"
	"github.com/gostudy03/xlog"
	"net"
	"math/rand"
)

var (
	roomList *protocal.AllRoomList
	user  *UserInfo
)

func init() {

	rand.Seed(time.Now().UnixNano())
	userId := rand.Int63()
	user = &UserInfo{
		UserId: uint64(userId),
		UserName: fmt.Sprint("user%d", userId),
	}
	xlog.LogDebug("generate default user_info, user_id:%d username:%s",  user.UserId, user.UserName)
}



func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		xlog.LogError("Error dialing", err.Error())
		return
	}
	defer conn.Close()
	
	//1. 拉取房间列表
	xlog.LogDebug("start get room list\n")
	roomList, err = getRoomList(conn)
	if err != nil {
		xlog.LogError("get room list failed, err:%v", err)
		return
	}

	for {
		showRoomList(roomList)
		var roomId uint64
		fmt.Scanf("%d", &roomId)
		err = enterRoom(conn, roomId)
		if err != nil {
			xlog.LogError("enter room failed, err:%v", err)
			continue
		}

		processRoomMessage(conn)
	}
}

func processRoomMessage(conn net.Conn) {

}
