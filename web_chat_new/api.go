package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gostudy03/xlog"
	"github.com/gostudy03/web_chat_new/dal"
	"net/http"
	"fmt"
)

const (
	ErrSuccess      = 0
	ErrUserNotLogin = 1001
	ErrServerBusy   = 1002
	ErrParameterInvalid = 1003
)

type ApiData struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ApiUserInfo struct {
	UserId int64 `json:"user_id"`
	Username   string    `json:"username"`
	Nickname   string    `json:"nickname"`
	Sex        int       `json:"sex"`
}

func GetMessage(code int) string {
	switch code {
	case ErrServerBusy:
		return "服务繁忙"
	case ErrUserNotLogin:
		return "用户未登录"
	case ErrSuccess:
		return "成功"
	case ErrParameterInvalid:
		return "参数非法"
	}

	xlog.LogError("未知错误，code：%d", code)
	return "未知错误"
}

func responseError(code int, ctx *gin.Context) {
	var apiData ApiData
	apiData.Code = code
	apiData.Message = GetMessage(code)

	xlog.LogError("response error, code：%d", code)
	ctx.JSON(http.StatusOK, &apiData)
}


func responseSuccess(data interface{}, ctx *gin.Context) {
	var apiData ApiData
	apiData.Code = ErrSuccess
	apiData.Message = GetMessage(ErrSuccess)
	apiData.Data = data

	ctx.JSON(http.StatusOK, &apiData)
}

func userInfoHandle(ctx *gin.Context) {
	
	logined := IsLogin(ctx)
	if !logined {
	    responseError(ErrUserNotLogin, ctx)
		return
	}

	userId := GetUserId(ctx)
	if userId <= 0 {
		responseError(ErrUserNotLogin, ctx)
		return
	}

	userInfo, err := dal.GetUserInfoById(userId)
	if err != nil {
		responseError(ErrServerBusy, ctx)
		return
	}

	var apiUserInfo ApiUserInfo
	apiUserInfo.Nickname = userInfo.Nickname
	apiUserInfo.Sex = userInfo.Sex
	apiUserInfo.UserId = userInfo.UserId
	apiUserInfo.Username = userInfo.Username

	responseSuccess(&apiUserInfo, ctx)
}

func roomMessageHandle(ctx *gin.Context) {
	
	logined := IsLogin(ctx)
	if !logined {
	    responseError(ErrUserNotLogin, ctx)
		return
	}

	roomId, ok := ctx.GetQuery("room_id")
	if !ok {
		responseError(ErrParameterInvalid, ctx)
		return
	}

	rId, err := strconv.ParseInt(roomId, 10, 64)
	if err != nil {
		responseError(ErrParameterInvalid, ctx)
		xlog.LogError("invalid paramter, roomId:%v", roomId)
		return
	}

	messageList, err := dal.GetRecentMsg(rId, 20)
	if err != nil {
		xlog.LogError("get recent message failed, err:%v, roomid:%d", err, rId)
		responseError(ErrServerBusy, ctx)
		return
	}

	for _, message := range messageList {
		message.ImageUrl = fmt.Sprintf("/static/image/%d.jpg", message.UserId%15+1)
	}

	length := len(messageList)
	for i := 0; i < length/2;i++ {
		messageList[i], messageList[length-i-1] = messageList[length-i-1], messageList[i] 
	}
	responseSuccess(messageList, ctx)
}
