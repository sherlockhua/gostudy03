package main

import (
	"github.com/gostudy03/day15/tcp_chat/protocal"
	"sync"
	"sync/atomic"
)

var (
	roomMgr *RoomMgr
)

func init() {
	loveRoom := &protocal.Room{
		RoomId:  1,
		Name:    "谈情说爱",
		RoomCap: 500,
		Desc:    "",
		Online:  0,
	}

	var allRoomList []*protocal.Room
	allRoomList = append(allRoomList, loveRoom)
	GoRoom := &protocal.Room{
		RoomId:  2,
		Name:    "Go开发论坛",
		RoomCap: 500,
		Desc:    "",
		Online:  0,
	}

	allRoomList = append(allRoomList, GoRoom)
	JavaRoom := &protocal.Room{
		RoomId:  3,
		Name:    "Java开发论坛",
		RoomCap: 500,
		Desc:    "",
		Online:  0,
	}

	allRoomList = append(allRoomList, JavaRoom)
	roomMgr = NewRoomMgr(allRoomList)
}

func NewRoomMgr(roomList []*protocal.Room) (roomMgr *RoomMgr) {
	roomMgr = &RoomMgr{}
	for _, r := range roomList {
		roomInfo := &RoomInfo{
			room:    r,
			userMap: make(map[uint64]*User, 1024),
		}

		roomMgr.AllRoomList = append(roomMgr.AllRoomList, roomInfo)
	}

	return
}

type RoomMgr struct {
	AllRoomList []*RoomInfo
	lock        sync.Mutex
}

func (r *RoomMgr) GetRoomList() (roomList []*protocal.Room) {
	for _, roomInfo := range r.AllRoomList {
		roomList = append(roomList, roomInfo.room)
	}

	return
}

func (r *RoomMgr) EnterRoom(user *User, roomId uint64) (err error) {

	var curRoomInfo *RoomInfo
	r.lock.Lock()
	for _, roomInfo := range r.AllRoomList {
		if roomInfo.room.RoomId == roomId {
			curRoomInfo = roomInfo
			break
		}
	}
	r.lock.Unlock()

	if curRoomInfo == nil {
		return user.handleEnterRoomResult(protocal.ErrCodeInvalidRoomId, nil)
	}

	atomic.AddUint32(&curRoomInfo.room.Online, 1)
	err = user.handleEnterRoomResult(protocal.ErrCodeSuccess, curRoomInfo.room)

	//2. 通知当前房间里的其他用户，有新用户加入房间
	broadcastEnterRoom := &protocal.BoardcastUserEnterRoom{
		EnterUserId:   user.UserId,
		EnterUserName: user.UserName,
		RoomInfo:      curRoomInfo.room,
	}

	msg := &Message{
		CmdNo: protocal.BoardcastUserEnterRoomCmd,
		Body:  broadcastEnterRoom,
	}

	for _, otherUser := range curRoomInfo.userMap {
		otherUser.AppendMessage(msg)
	}

	//3. 把当前用户保存到userMap里面
	curRoomInfo.userMap[user.UserId] = user
	return
}


func (r *RoomMgr) SendText(sendText *protocal.UserSendText) (err error) {
	
		var curRoomInfo *RoomInfo
		r.lock.Lock()
		for _, roomInfo := range r.AllRoomList {
			if roomInfo.room.RoomId == sendText.RoomId {
				curRoomInfo = roomInfo
				break
			}
		}
		r.lock.Unlock()
	
		if curRoomInfo == nil {
			//TODO: 需要保存一个当前在线的用户列表
			//return user.handleSendText(protocal.ErrCodeInvalidRoomId, nil)
			return
		}
	
		//2. 从房间里找到这个用户
		/*
		user, ok := curRoomInfo.userMap[sendText.UserId]
		if !ok {
			return
		}*/
	
		//3. 通知当前房间里的其他用户，有用户发言
		broadcastRecvText := &protocal.UserRecvText{
			Code: protocal.ErrCodeSuccess,
			RoomId: sendText.RoomId,
			AuthorUserId: sendText.UserId,
			AuthorUserName: sendText.UserName,
			Content: sendText.Content,
		}
	
		msg := &Message{
			CmdNo: protocal.UserRecvTextCmd,
			Body:  broadcastRecvText,
		}
	
		for _, otherUser := range curRoomInfo.userMap {
			otherUser.AppendMessage(msg)
		}
	
		return
	}
	