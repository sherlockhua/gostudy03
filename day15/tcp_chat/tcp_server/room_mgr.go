package main

import(
	"github.com/gostudy03/day15/tcp_chat/protocal"
	"sync"
)

var (
	roomMgr *RoomMgr
)

func init() {
	loveRoom := &protocal.Room{
		RoomId: 1,
		Name: "谈情说爱",
		RoomCap: 500,
		Desc: "",
		Online:0,
	}

	var allRoomList []*protocal.Room
	allRoomList = append(allRoomList, loveRoom)
	GoRoom := &protocal.Room{
		RoomId: 2,
		Name: "Go开发论坛",
		RoomCap: 500,
		Desc: "",
		Online:0,
	}

	allRoomList = append(allRoomList, GoRoom)
	JavaRoom := &protocal.Room{
		RoomId: 3,
		Name: "Java开发论坛",
		RoomCap: 500,
		Desc: "",
		Online:0,
	}

	allRoomList = append(allRoomList, JavaRoom)
	roomMgr = NewRoomMgr(allRoomList)
}

func NewRoomMgr(roomList []*protocal.Room) (roomMgr *RoomMgr) {
	roomMgr = &RoomMgr{}
	for _, r := range roomList {
		roomInfo := &RoomInfo{
			room:r,
			userMap: make(map[uint64]*User, 1024),
		}

		roomMgr.AllRoomList = append(roomMgr.AllRoomList, roomInfo)
	}

	return 
}

type RoomMgr struct {
	AllRoomList []*RoomInfo
	lock sync.Mutex
}

func (r *RoomMgr) GetRoomList() (roomList []*protocal.Room) {
	for _,  roomInfo := range r.AllRoomList {
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
	return
}