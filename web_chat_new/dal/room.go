package dal

import (
	"github.com/gostudy03/web_chat_new/common"
	"github.com/gostudy03/xlog"
)

func GetAllRoomList() (roomList []*common.Room, err error) {
	sqlstr := "select id, room_id, room_name, `desc`, cap, create_time, online from room where status=1"
	 err = Db.Select(&roomList, sqlstr)
	if err != nil {
		xlog.LogError("insert failed, err:%v", err)
		return
	}
	return
}


func UpdateRoomOnline(roomId int64) (err error) {
	sqlstr := "update  room set online = online+1 where room_id=?"
	_, err = Db.Exec(sqlstr, roomId)
	return
}