package main

import (
	"github.com/gostudy03/xlog"
	"github.com/gostudy03/day15/tcp_chat/protocal"
	"fmt"
	"net"
)

func getRoomList(conn net.Conn) (roomList *protocal.AllRoomList, err error) {

	var getRoomList = &protocal.GetRoomList {
		UserId: user.UserId,
	}

	data, err := protocal.Pack(protocal.GetRoomListCmd, getRoomList)
	if err != nil {
		fmt.Printf("pack failed, err:%v\n", err)
		return
	}
	fmt.Printf("get room list cmd, data:%d, err:%v\n", len(data), err)

	_, err = conn.Write(data)
	if err != nil {
		fmt.Printf("write data failed, err:%v\n", err)
		return
	}

	cmdNo, result, err := protocal.UnPack(conn)
	if err != nil {
		xlog.LogError("read message from server failed, err:%v", err)
		return
	}

	if cmdNo != protocal.AllRoomListCmd {
		err = fmt.Errorf("unexpected package, cmd:%d, data:%v", cmdNo, data)
		xlog.LogError("unexpected package, cmd:%d, data:%v", cmdNo, data)
		return
	}
	
	roomList, ok := result.(*protocal.AllRoomList)
	if !ok {
		err = fmt.Errorf("convert to *protocal.AllRoomList failed")
		xlog.LogError("err:%v, data:%#v", err, result)
		return
	}

	return
}

func showRoomList(roomList *protocal.AllRoomList) {
	fmt.Printf("==========================激情聊天室================\n")
	fmt.Printf("房间列表\n")
	for _, room := range roomList.RoomList{
		showRoom(room)
	}

	fmt.Printf("please select room id to enter\n")
}

func showRoom(room *protocal.Room) {
	fmt.Printf("房间编号:%d\n", room.RoomId)
	fmt.Printf("房间名称:%s\n", room.Name)
	fmt.Printf("房间描述:%s\n", room.Desc)
	fmt.Printf("在线人数:%d\n", room.Online)
	fmt.Println()
}

func enterRoom(conn net.Conn, roomId uint64) (roomInfo *protocal.Room, err error) {
	//1. 校验roomid是否合法
	var validRoomId bool
	for _, room := range roomList.RoomList {
		if room.RoomId == roomId {
			validRoomId = true
			break
		}
	}
	_ = validRoomId
/*
	if validRoomId == false {
		err = fmt.Errorf("invalid room id")
		return
	}
*/
	var enterRoom = &protocal.UserEnterRoom{
		RoomId: roomId,
		UserId: user.UserId,
		UserName: user.UserName,
	}

	//打包成网络字节流
	data, err := protocal.Pack(protocal.UserEnterRoomCmd, enterRoom)
	if err != nil {
		xlog.LogError("pack failed, err:%v\n", err)
		return
	}
	
	_, err = conn.Write(data)
	if err != nil {
		fmt.Printf("write data failed, err:%v\n", err)
		return
	}

	cmdNo, result, err := protocal.UnPack(conn)
	if err != nil {
		xlog.LogError("read message from server failed, err:%v", err)
		return
	}

	if cmdNo != protocal.UserEnterRoomRespCmd {
		err = fmt.Errorf("unexpected package, cmd:%d, data:%v", cmdNo, data)
		xlog.LogError("unexpected package, cmd:%d, data:%v", cmdNo, data)
		return
	}
	
	roomResp, ok := result.(*protocal.UserEnterRoomResp)
	if !ok {
		err = fmt.Errorf("convert to *protocal.AllRoomList failed")
		xlog.LogError("err:%v, data:%#v", err, result)
		return
	}

	if roomResp.Code != protocal.ErrCodeSuccess {
		err = fmt.Errorf("enter room failed, code:%d", roomResp.Code)
		xlog.LogError("enter room failed, code:%d", roomResp.Code)
		return
	}

	fmt.Printf("enter room success\n")
	showRoom(roomResp.RoomInfo)
	roomInfo = roomResp.RoomInfo
	return
}