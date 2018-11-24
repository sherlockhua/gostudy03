package main

import (
	"fmt"
	"time"
	"github.com/gostudy03/web_chat/dal"
	"github.com/gostudy03/web_chat/common"

	"github.com/gostudy03/xlog"
	"sort"
)

var (
	roomMgr *RoomMgr
)

type RoomList []*common.Room


/*
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/
func (rl RoomList)Len() int{
	roomList := []*common.Room(rl)
	return len(roomList)
}

func (rl RoomList)Less(i, j int) bool{
	roomList := []*common.Room(rl)
	r1 := roomList[i]
	r2 := roomList[j]

	return r1.Online > r2.Online
}

func (rl RoomList)Swap(i, j int) {
	roomList := []*common.Room(rl)
	roomList[i], roomList[j] = roomList[j], roomList[i]
}


type RoomInfo struct {
	Room *common.Room
	UserMap map[int64]*UserInfo

}

func (r *RoomInfo) DeleteUser(user *UserInfo) {
	delete(r.UserMap, user.User.UserId)
}

func (r *RoomInfo) AddUser(user *UserInfo) (isAlreadyLogin bool) {
	_,  isAlreadyLogin = r.UserMap[user.User.UserId]
	if isAlreadyLogin {
		return 
	}

	r.UserMap[user.User.UserId] = user
	err := dal.UpdateRoomOnline(r.Room.RoomId)
	if err != nil {
		xlog.LogError("update room online failed, room_id:%d err:%v", r.Room.RoomId, err)
		return
	}
	return
}

type RoomMgr struct {
	RoomMap map[int64]*RoomInfo
}

func NewRoomMgr() *RoomMgr{
	return &RoomMgr{
		RoomMap:make(map[int64]*RoomInfo, 16),
	}
}

func (r *RoomMgr) GetRoom(roomId int64)(roomInfo *RoomInfo, err error) {
	roomInfo, ok := r.RoomMap[roomId]
	if !ok {
		xlog.LogError("room not exists, room_id:%d", roomId)
		err = fmt.Errorf("room not exists, id:%d", roomId)
		return
	}

	return
}

func (r *RoomMgr) Init(roomList []*common.Room) (err error) {
	for _, room := range roomList {
		roomInfo := &RoomInfo{
			Room: room,
			UserMap: make(map[int64]*UserInfo, 1024),
		}

		r.RoomMap[room.RoomId] = roomInfo
	}
	
	return
}

func (r *RoomMgr) GetRoomList() (roomList []*common.Room, err error) {
	for _, v := range r.RoomMap {
		roomList = append(roomList, v.Room)
	}
	
	var sortRoomList RoomList = RoomList(roomList)
	sort.Sort(sortRoomList)
	for _, r := range sortRoomList {
		xlog.LogDebug("sort result:%#v", r)
	}
	roomList = []*common.Room(sortRoomList)
	return
}

func (r *RoomMgr) SyncRoomList()  {
	for {
		time.Sleep(time.Second)
		roomList, err := dal.GetAllRoomList()
		if err != nil {
			xlog.LogError("get all room List from Db failed, err:%v", err)
			continue
		}
		for _, room := range roomList {
			_, ok := r.RoomMap[room.RoomId]
			if !ok {
				roomInfo := &RoomInfo{
					Room: room,
					UserMap: make(map[int64]*UserInfo, 1024),
				}
		
				r.RoomMap[room.RoomId] = roomInfo
				continue
			}

			r.RoomMap[room.RoomId].Room = room
		}
	}
}

func InitRoomMgr() (err error) {

	roomMgr = NewRoomMgr()
	roomList, err := dal.GetAllRoomList()
	if err != nil {
		return
	}

	err = roomMgr.Init(roomList)
	go roomMgr.SyncRoomList()
	return
}