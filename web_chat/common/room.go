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