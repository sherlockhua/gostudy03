package dal

import (
	"github.com/gostudy03/web_chat_new/common"
	"github.com/gostudy03/xlog"
)

func Register(userInfo *common.User) (err error) {
	sqlstr := "insert into user(user_id, username, nickname, sex, password)values(?,?,?,?, ?)"
	_, err = Db.Exec(sqlstr, userInfo.UserId, userInfo.Username, userInfo.Nickname,
		userInfo.Sex, userInfo.Passwd)
	if err != nil {
		xlog.LogError("insert failed, err:%v", err)
		return
	}
	return
}

func GetUserInfoByName(username string) (userInfo *common.User, err error) {

	userInfo = &common.User{}
	sqlstr := "select user_id, username, nickname, sex, password from user where username=?"
	err = Db.Get(userInfo, sqlstr, username)
	if err != nil {
		xlog.LogError("get user info by username:%s failed, err:%v", username, err)
		return
	}
	return
}


func GetUserInfoById(userId int64) (userInfo *common.User, err error) {
	
		userInfo = &common.User{}
		sqlstr := "select user_id, username, nickname, sex, password from user where user_id=?"
		err = Db.Get(userInfo, sqlstr, userId)
		if err != nil {
			xlog.LogError("get user info by user_id:%d failed, err:%v", userId, err)
			return
		}
		return
	}
