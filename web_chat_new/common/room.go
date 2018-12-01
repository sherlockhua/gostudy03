package common

import (
	"time"
)

type Room struct {
	Id int64 `db:"id"`
	RoomId int64 `db:"room_id"`
	RoomName string `db:"room_name"`
	Desc string `db:"desc"`
	Status int `db:"status"`
	Cap int `db:"cap"`
	Online int `db:"online"`
	CreateTime time.Time `db:"create_time"`
}

type RoomMessage struct {
	Id		int64 `db:"id" json:"id"`
	MsgId uint64 `db:"msg_id" json:"msg_id"`
	RoomId int64 `db:"room_id" json:"room_id"`
	Content string `db:"content" json:"content"`
	UserId int64 `db:"user_id" json:"user_id"`
	ImageUrl string `json:"image_url"`

}