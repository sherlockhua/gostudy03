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


func GetRoomInfoById(roomId int64) (roomInfo *common.Room, err error) {
	roomInfo = &common.Room{}
	sqlstr := "select id, room_id, room_name, `desc`, cap, create_time, online from room where status=1 and room_id=?"
	 err = Db.Get(roomInfo, sqlstr, roomId)
	if err != nil {
		xlog.LogError("insert failed, err:%v", err)
		return
	}
	return
}


func UpdateRoomOnline(roomId int64, count int) (err error) {
	sqlstr := "update  room set online = online+? where room_id=?"
	_, err = Db.Exec(sqlstr, count, roomId)
	return
}


func InsertMessage(msg *common.RoomMessage) (err error) {
	/*
	MsgId int64 `db:"msg_id"`
	RoomId int64 `db:"room_id"`
	Content string `db:"content"`
	UserId int64 `db:"user_id"`*/
	sqlstr := "insert into message(msg_id, room_id, content, user_id)values(?,?,?,?)"
	_, err = Db.Exec(sqlstr, msg.MsgId, msg.RoomId, msg.Content,
		msg.UserId)
	if err != nil {
		xlog.LogError("insert failed, err:%v", err)
		return
	}
	return
}

func GetRecentMsg(roomId int64, limit int) (messageList []*common.RoomMessage, err error) {
	/*
	MsgId int64 `db:"msg_id"`
	RoomId int64 `db:"room_id"`
	Content string `db:"content"`
	UserId int64 `db:"user_id"`*/
	sqlstr := "select id, msg_id, room_id, content, user_id from message where room_id=? order by id desc limit ?"
	err = Db.Select(&messageList, sqlstr, roomId, limit)
	return
}