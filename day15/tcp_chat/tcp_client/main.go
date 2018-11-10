package main

import (
	"os"
	"bufio"
	"time"
	"github.com/gostudy03/day15/tcp_chat/protocal"
	"fmt"
	"github.com/gostudy03/xlog"
	"net"
	"math/rand"
	"context"
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
		UserName: fmt.Sprintf("user%d", userId),
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
		fmt.Scanf("%d\n", &roomId)
		roomInfo, err := enterRoom(conn, roomId)
		if err != nil {
			xlog.LogError("enter room failed, err:%v", err)
			continue
		}

		processRoomMessage(conn, roomInfo)
	}
}

func processRoomMessage(conn net.Conn, roomInfo *protocal.Room) {
	fmt.Printf("enter room succ")
	ctx, cancel := context.WithCancel(context.Background())

	defer func() {
		cancel()
	}()

	go recvMessage(ctx, conn)
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			xlog.LogError("read string failed, err:%v", err)
			continue
		}

		err = sendText(conn, text, roomInfo)
		if err != nil {
			xlog.LogError("send text failed, err:%v", err)
			continue
		}
	}
}

func sendText(conn net.Conn, text string, roomInfo*protocal.Room) (err error) {

	userSendText := &protocal.UserSendText{
		RoomId: roomInfo.RoomId,
		UserId: user.UserId,
		UserName: user.UserName,
		Content:text,
	}

	data, err := protocal.Pack(protocal.UserSendTextCmd, userSendText)
	if err != nil {
		return
	}

	_, err = conn.Write(data)
	return
}

func recvMessage(ctx context.Context, conn net.Conn) {
	for {
		select {
		case <- ctx.Done():
			return
		default:
		}

		cmdNo, result, err := protocal.UnPack(conn)
		if err != nil {
			return
		}

		switch(cmdNo) {
		case protocal.BoardcastUserEnterRoomCmd:
			broadcastUserEnterRoom, ok := result.(*protocal.BoardcastUserEnterRoom)
			if !ok {
				break
			}

			fmt.Printf("username:%s enter room\n", broadcastUserEnterRoom.EnterUserName)

		case protocal.UserRecvTextCmd:
			recvText, ok := result.(*protocal.UserRecvText)
			if !ok {
				break
			}

			fmt.Printf("%s:\n", recvText.AuthorUserName)
			fmt.Printf("  %s\n", recvText.Content)
		}
	}
}