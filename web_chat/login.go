package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gostudy03/web_chat/common"
	"github.com/gostudy03/web_chat/dal"
	"github.com/gostudy03/web_chat/id_gen"
	"github.com/gostudy03/xlog"
	"net/http"
	"strings"
)

func loginView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "./views/login.html", nil)
}

func makeNewPassword(password string) (newPassword string) {
	passwd := fmt.Sprintf("%s.%s", password, UserPasswdSalt)
	result := md5.Sum([]byte(passwd))
	newPassword = fmt.Sprintf("%x", string(result[:]))
	return
}

func registerHandle(ctx *gin.Context) {

	username := ctx.PostForm("username")
	nickname := ctx.PostForm("nickname")
	password := ctx.PostForm("password")
	confirmPassword := ctx.PostForm("confirm_password")

	password = strings.TrimSpace(password)
	confirmPassword = strings.TrimSpace(confirmPassword)
	nickname = strings.TrimSpace(nickname)
	username = strings.TrimSpace(username)

	if password != confirmPassword {
		xlog.LogError("两次不一样， password:%s confirm password:%s", password, confirmPassword)
		jsContent := `<html><script>alert("两次密码不一样！");</script></html>`
		ctx.Writer.Write([]byte(jsContent))
		ctx.Writer.Flush()
		return
	}
	url := "http://localhost:9090/id/gen"
	id, err := id_gen.GetId(url)
	if err != nil {
		xlog.LogError("call id generator failed, err:%v", err)
		return
	}

	var userInfo common.User
	userInfo.Nickname = nickname
	userInfo.UserId = int64(id)
	userInfo.Username = username
	userInfo.Passwd = makeNewPassword(password)

	err = dal.Register(&userInfo)
	if err != nil {
		xlog.LogError("register failed, err:%v", err)
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, "/user/login")
}

func loginHandle(ctx *gin.Context) {

	sess, err := globalSessions.SessionStart(ctx.Writer, ctx.Request)
	if err != nil {
		xlog.LogError("session start failed, err:%v", err)
		return
	}

	defer sess.SessionRelease(ctx.Writer)

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	userInfo, err := dal.GetUserInfoByName(username)
	if err != nil {
		xlog.LogError("get user info failed, err:%v, username:%v", err, username)
		return
	}

	newPassword := makeNewPassword(password)
	if newPassword != userInfo.Passwd {
		xlog.LogError("username or password not right, password:%s newpassword:%s db pass:%s",
			password, newPassword, userInfo.Passwd)
		return
	}

	xlog.LogDebug("login succ, user_info:%#v", userInfo)
	sess.Set(UserIDKey, userInfo.UserId)
	sess.Set(UserNameKey, userInfo.Username)
	//千万不能调用flush，会把session设置的值给清空了。并不是保存到redis中。
	//sess.Flush()

	ctx.Redirect(http.StatusMovedPermanently, "/index")
}
