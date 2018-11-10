package main

import (
	"github.com/gostudy03/xlog"
	"github.com/gostudy03/day15/tcp_chat/protocal"
	
	"fmt"
	"net"
	
)

func main() {
	fmt.Println("start server...")

	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
func process(conn net.Conn) {
	defer conn.Close()
	for {
		cmdNo, message, err := protocal.UnPack(conn)
		if err != nil {
			xlog.LogError("read package from client failed, close client,err:%v", err)
			return
		}

		err = processMessage(conn, cmdNo, message)
		if err != nil {
			xlog.LogError("process message failed, err:%v", err)
			continue
		}

		xlog.LogDebug("process message succ, cmdNo:%d package:%v", cmdNo, message)
	}
}

func processMessage(conn net.Conn, cmdNo uint16, message interface{}) (err error){
	switch(cmdNo) {
	case protocal.GetRoomListCmd:
		getRoomList, ok := message.(*protocal.GetRoomList)
		if !ok {
			xlog.LogError("convert to protocal.GetRoomList failed, message:%#v", message)
			return
		}
		return procGetRoomList(conn, cmdNo, getRoomList)

	case protocal.UserEnterRoomCmd:
		enterRoom, ok := message.(*protocal.UserEnterRoom)
		if !ok {
			xlog.LogError("convert to protocal.GetRoomList failed, message:%#v", message)
			return
		}
		return procEnterRoom(conn, cmdNo, enterRoom)
	}

	return
}

func procGetRoomList(conn net.Conn, cmdNo uint16, getRoomList *protocal.GetRoomList)(err error) {
	xlog.LogDebug("start process get room list, user_id:%d", getRoomList.UserId)

	allRoomList := &protocal.AllRoomList{}
	for _, room := range roomMgr.GetRoomList() {
		allRoomList.RoomList = append(allRoomList.RoomList, room)
	}

	data, err := protocal.Pack(protocal.AllRoomListCmd, allRoomList)
	if err != nil {
		xlog.LogError("pack message failed, data:%#v, err:%v", allRoomList, err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		xlog.LogError("send data to client failed, err:%v, data:%#v", err, allRoomList)
	}
	return
}


func procEnterRoom(conn net.Conn, cmdNo uint16, enterRoom *protocal.UserEnterRoom)(err error) {
	xlog.LogDebug("start process user enter room, user_id:%d, room_id:%d", 
		enterRoom.UserId, enterRoom.RoomId)

	user := &User{
		UserId: enterRoom.UserId,
		UserName: enterRoom.UserName,
		Conn:conn,
	}

	roomMgr.EnterRoom(user, enterRoom.RoomId)
	return
}