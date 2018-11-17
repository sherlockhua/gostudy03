package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gostudy03/xlog"
)

const (
	UserIDKey   = "user_id"
	UserNameKey = "username"
)

func UserMiddleware(ctx *gin.Context) {

	sess, err := globalSessions.SessionStart(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.Set(UserIDKey, int64(0))
		return
	}

	var userId int64
	var username string
	var ok bool

	defer func()  {
		xlog.LogDebug("use_id:%d username:%s", userId, username)
	}()
	userId, ok = sess.Get(UserIDKey).(int64)
	if !ok {
		ctx.Set(UserIDKey, int64(0))
		return
	}

	ctx.Set(UserIDKey, int64(userId))

	username, ok = sess.Get(UserNameKey).(string)
	if !ok {
		return
	}

	
	ctx.Set(username, username)
	defer sess.SessionRelease(ctx.Writer)
	ctx.Next()
}

func IsLogin(ctx *gin.Context) bool {

	userId := GetUserId(ctx)
	if userId <= 0 {
		return false
	}

	return true
}

func GetUserId(ctx *gin.Context) (userId int64) {

	userIdTmp, exists := ctx.Get(UserIDKey)
	if !exists {
		return
	}

	var ok bool
	userId, ok = userIdTmp.(int64)
	if !ok {
		return
	}
	return
}

func GetUserName(ctx *gin.Context) (username string) {

	usernameTmp, exists := ctx.Get(UserNameKey)
	if !exists {
		return
	}

	var ok bool
	username, ok = usernameTmp.(string)
	if !ok {
		return
	}
	return
}
