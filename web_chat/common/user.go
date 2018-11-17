package common

import (
	"time"
)

type User struct {
	Id         int64     `db:"id"`
	UserId     int64     `db:"user_id"`
	Username   string    `db:"username"`
	Nickname   string    `db:"nickname"`
	Sex        int       `db:"sex"`
	Passwd     string    `db:"password"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}
